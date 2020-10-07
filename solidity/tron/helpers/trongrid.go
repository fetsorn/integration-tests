package helpers

import (
	"bytes"
	"net/http"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"log"
	"time"
)

type keyval map[string]interface{}

type BroadcastResult struct {
	Result bool `json:"result"`
	TxId string `json:"txid"`
}

type TxResultDataRet struct {
	ContractRet string `json:"contractRet"`
}

type TxResultRawData struct {
	RefBlockBytes string `json:"ref_block_bytes"`
}

type TxResultData struct {
	Ret []TxResultDataRet
	TxId string `json:"tx_id"`
	RawData TxResultRawData `json:"raw_data"`
	ContractAddress string `json:"contract_address"`
}

type TxResult struct {
	Success bool `json:"success"`
	Data []TxResultData `json:"data"`
}

func HexFromPk(pk string) (string) {
	privateKey, err := crypto.HexToECDSA(pk)
    if err != nil {
        log.Fatal(err)
    }

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	data := elliptic.Marshal(crypto.S256(), publicKeyECDSA.X, publicKeyECDSA.Y)

	return "41" + hex.EncodeToString(crypto.Keccak256(data[1:]))[24:]
}

func HexToTronAddress(hexAddress string) (string) {
	address := common.HexToAddress(hexAddress)
	return base58.CheckEncode(address.Bytes(), 0x41)
}

func ApiRequest(endpoint string, data keyval) (string, error) {
	buf, _ := json.Marshal(data)
	rbuf := bytes.NewReader(buf)
	txContent, _ := http.Post(endpoint, "application/json", rbuf)
	rBodyBuf := new(bytes.Buffer)
	rBodyBuf.ReadFrom(txContent.Body)
	return rBodyBuf.String(), nil
}

func SignAndBroadcast(endpoint string, tx string, key string) (string, error) {
	body, err := ApiRequest(endpoint + "/wallet/gettransactionsign", keyval{"transaction": tx, "privateKey": key})
	if err != nil {
		return body, err
	}
	rbuf := bytes.NewReader([]byte(body))
	txContent, _ := http.Post(endpoint + "/wallet/broadcasttransaction", "application/json", rbuf)
	rBodyBuf := new(bytes.Buffer)
	rBodyBuf.ReadFrom(txContent.Body)

	return rBodyBuf.String(), err
}
func WaitForEvent(txid string, event string, endpoint string) (map[string]string, error) {
	for i := 0; i < 60; i++ {
		ret, err := GetEventByName(txid, event, endpoint)
		if err == nil {
			return ret, nil
		}
		time.Sleep(time.Second)
	}
	return nil, errors.New("event not appeared")
}

func WaitForTx(txid string, endpoint string) (*TxResult) {
	var deployResult TxResult
	for i := 0; i < 60; i++ {
		txContent, err := http.Get(endpoint + "/v1/transactions/" + txid)
		rBodyBuf := new(bytes.Buffer)
		rBodyBuf.ReadFrom(txContent.Body)
		if err != nil {
			log.Fatal(err)
		}

		rBodyBufString := rBodyBuf.String()

		json.Unmarshal([]byte(rBodyBufString), &deployResult)

		if deployResult.Success && (len(deployResult.Data) > 0) && (len(deployResult.Data[0].Ret) > 0) && (deployResult.Data[0].Ret[0].ContractRet != "") {
			if deployResult.Data[0].Ret[0].ContractRet != "SUCCESS" {
				return nil
			}
			return &deployResult
		}
		time.Sleep(time.Second)
	}
	return nil
}

func DeployContract(endpoint string, contractAbi string, _bytecode string, owner string, key string, name string, params ...interface{}) (*string, error) {
	var broadcastResult BroadcastResult

	paramsEncoder, err := abi.JSON(bytes.NewReader([]byte(contractAbi)))

	if err != nil {
		return nil, err
	}

	paramsEncoded, err := paramsEncoder.Pack("", params...)
	if err != nil {
		return nil, err
	}

	var bytecode string
	if _bytecode[0:2] == "0x" {
		bytecode = _bytecode[2:]
	} else {
		bytecode = _bytecode
	}

	tx, err := ApiRequest(endpoint + "/wallet/deploycontract", keyval{
		"abi":contractAbi, "bytecode": bytecode, "name": name, "owner_address": owner,
		"fee_limit": "1000000000",
		"parameter": hex.EncodeToString(paramsEncoded),
		"origin_energy_limit": "1000000000"})

	if err != nil {
		return nil, err
	}

	br, err := SignAndBroadcast(endpoint, tx, key)

	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(br), &broadcastResult)

	if !broadcastResult.Result {
		return nil, errors.New("broadcast failed")
	}

	deployResult := WaitForTx(broadcastResult.TxId, endpoint)
	if deployResult == nil {
		return nil, errors.New("deploy failed")
	}
	return &deployResult.Data[0].ContractAddress, nil
}

type Transactor struct {
	endpoint string
	abiEncoder abi.ABI
	address string
}

func NewTransactor(endpoint string, abiString string, address string) (*Transactor, error) {
	abiEncoder, err := abi.JSON(bytes.NewReader([]byte(abiString)))
	if err != nil {
		return nil, err
	}
	return &Transactor{endpoint: endpoint, abiEncoder: abiEncoder, address: address}, nil
}

type ReadContractResultValue struct {
	Value bool `json:"result"`
}

type ReadContractResult struct {
	Result ReadContractResultValue `json:"result"`
	ConstantResult []string `json:"constant_result"`
}

func (transactor *Transactor) ReadContract(key string, f string, result interface{}, args ...interface{}) (error) {
	var contractResult ReadContractResult

	parameter, err := transactor.abiEncoder.Pack(f, args...)

	if err != nil {
		return err
	}

	method := transactor.abiEncoder.Methods[f]

	tx, err := ApiRequest(transactor.endpoint + "/wallet/triggerconstantcontract", keyval{
		"owner_address": HexFromPk(key),
		"contract_address": transactor.address,
		"function_selector": method.Sig,
		"parameter": hex.EncodeToString(parameter)[8:],
		"fee_limit": "1000000000",
		"call_value": 0})

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(tx), &contractResult)

	if !contractResult.Result.Value {
		return errors.New("view call failed")
	}

	s, err := hex.DecodeString(contractResult.ConstantResult[0])
	if err != nil {
		return err
	}

	return transactor.abiEncoder.Unpack(result, f, []byte(s))
}

func (transactor *Transactor) Transact(key string, f string, args ...interface{}) (*string, error) {
	var kv keyval
	var broadcastResult BroadcastResult

	parameter, err := transactor.abiEncoder.Pack(f, args...)

	if err != nil {
		return nil, err
	}

	method := transactor.abiEncoder.Methods[f]

	tx, err := ApiRequest(transactor.endpoint + "/wallet/triggersmartcontract", keyval{
		"owner_address": HexFromPk(key),
		"contract_address": transactor.address,
		"function_selector": method.Sig,
		"parameter": hex.EncodeToString(parameter)[8:],
		"fee_limit": "1000000000",
		"call_value": 0})

	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(tx), &kv)
	transaction, _ := json.Marshal(kv["transaction"])
	transactionStr := string(transaction)

	br, err := SignAndBroadcast(transactor.endpoint, transactionStr, key)
	json.Unmarshal([]byte(br), &broadcastResult)

	if !broadcastResult.Result {
		return nil, errors.New("broadcast failed")
	}

	return &broadcastResult.TxId, nil
}

type EventResultItem struct {
	EventName string `json:"event_name"`
	Result map[string]string `json:"result"`
}

func GetEventByName(txid string, event string, endpoint string) (map[string]string, error) {
	var eventResult []EventResultItem

	txContent, err := http.Get(endpoint + "/event/transaction/" + txid)
	rBodyBuf := new(bytes.Buffer)
	rBodyBuf.ReadFrom(txContent.Body)
	if err != nil {
		log.Fatal(err)
	}

	rBodyBufString := rBodyBuf.String()

	json.Unmarshal([]byte(rBodyBufString), &eventResult)

	for i := 0; i < len(eventResult); i++ {
		if eventResult[i].EventName == event {
			return eventResult[i].Result, nil
		}
	}

	return nil, errors.New("no required event " + event + " at " + txid)
}

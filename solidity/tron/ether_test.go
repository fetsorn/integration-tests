package main

import (
	"testing"
	"encoding/hex"
    "os"
    "log"
    "math/big"
	"math/rand"
	"errors"
	"rh_tests/helpers"
	"rh_tests/api/nebula"
	"rh_tests/api/ibport"
	"rh_tests/api/luport"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

var config helpers.Config
var addresses helpers.DeployedAddresses

func ReadConfig() {
	var err error
	config, err = helpers.LoadConfiguration()
    if err != nil {
        log.Fatal(err)
    }
	addresses, err = helpers.LoadAddresses()
    if err != nil {
        log.Fatal(err)
    }
}

func signData(dataHash [32]byte, validSignsCount int, isReverse bool) (*big.Int, error) {
    var r [5][32]byte
    var s [5][32]byte
	var v [5]uint8
	
    for position, validatorKey := range config.OraclePK {
		validatorEthKey, _ := crypto.HexToECDSA(validatorKey)
		
		seckey := math.PaddedBigBytes(validatorEthKey.D, validatorEthKey.Params().BitSize/8)
		sign, _ := secp256k1.Sign(dataHash[:], seckey)

        copy(r[position][:], sign[:32])
        copy(s[position][:], sign[32:64])

        if (position < validSignsCount) {
            v[position] = sign[64] + 27
        } else {
            v[position] = 0 // generate invalid signature
		}
	}
	
	var transactor *helpers.Transactor
	if !isReverse {
		transactor, _ = helpers.NewTransactor(config.Endpoint, nebula.NebulaABI, addresses.Nebula)
	} else {
		transactor, _ = helpers.NewTransactor(config.Endpoint, nebula.NebulaABI, addresses.NebulaReverse)
	}

	tx, err := transactor.Transact(config.TronPK, "sendHashValue", dataHash, v[:], r[:], s[:])

	if err != nil {
		return nil, err
	}

	receipt := helpers.WaitForTx(*tx, config.Endpoint)
	if receipt == nil {
		return nil, errors.New("tx returned REVERT " + *tx)
	}
	pulseResult, err := helpers.WaitForEvent(*tx, "NewPulse", config.Endpoint)
    if err != nil {
        return nil, err
	}

	i := new(big.Int)
	i.SetString(pulseResult["pulseId"], 10)
	return i, nil
}

func TestMain(m *testing.M) {
    ReadConfig()
	os.Exit(m.Run())
}

func Random32Byte() [32]byte {
	var out [32]byte
	rand.Read(out[:])
	return out
}

func filluint(data []byte, pos uint, val uint64) {
	var i int
	for i = 31; i >= 0; i-- {
		data[i + int(pos)] = byte(val % 256)
		val = val / 256
	}
}

func filladdress(data []byte, pos uint, addressStr string) {
	address := common.HexToAddress(addressStr)
	copy(data[pos:], address[:])
}

func bytes32fromhex(s string) ([32]byte) {
	var ret [32]byte
	decoded, err := hex.DecodeString(s[:])
	if err != nil {
		return ret
	}
	copy(ret[:], decoded[:])
	return ret
}

func sendData(key string, value []byte, blockNumber *big.Int, subscriptionId [32]byte, isReverse bool) (bool) {
	var transactor *helpers.Transactor

	if !isReverse {
		transactor, _ = helpers.NewTransactor(config.Endpoint, nebula.NebulaABI, addresses.Nebula)
	} else {
		transactor, _ = helpers.NewTransactor(config.Endpoint, nebula.NebulaABI, addresses.NebulaReverse)
	}

	tx, err := transactor.Transact(config.TronPK, "sendValueToSubByte", value, blockNumber, subscriptionId)
	if err != nil {
		return false
	}

	txResult := helpers.WaitForTx(*tx, config.Endpoint)
	if txResult == nil {
		return false
	}

	return true
}

type PulseData struct {
	DataHash [32]byte
	Height *big.Int
}

func GetPulseData(pulseId big.Int, transactor *helpers.Transactor) (*PulseData, error) {
	pulseData := new(PulseData)
	out := pulseData

	err := transactor.ReadContract(config.TronPK, "pulses", out, &pulseId)
	if err != nil {
		return nil, err
	}

	return pulseData, nil
}

func TestPulseSaved(t *testing.T) {
	d := Random32Byte()
	pulseId, err := signData(d, 5, false)
	if err != nil {
		t.Error("can't send signed data", err)
	} else {
		transactor, _ := helpers.NewTransactor(config.Endpoint, nebula.NebulaABI, addresses.Nebula)
		pulseData, err := GetPulseData(*pulseId, transactor)

		if err != nil {
			t.Error("can't get pulse hash", err)
		}

		if d != pulseData.DataHash {
			t.Error("data mismatch")
		}
	}
}

func TestPulseCorrect3(t *testing.T) {
	d := Random32Byte()
	_, err := signData(d, 3, false)
	if err != nil {
		t.Error("can't send signed data", err)
	}
}

func TestPulseInCorrect2(t *testing.T) {
	d := Random32Byte()
	_, err := signData(d, 2, false)
	if err == nil {
		t.Error("transaction 2/3 valid sigs should be rejected")
	}
}

func TestInvalidHash(t *testing.T) {
	var attachedData [2]byte
	attachedData[0] = 1
	attachedData[1] = 2

	// generate invalid proof
	proof := crypto.Keccak256Hash(attachedData[:])

	pulseId, err := signData(proof, 5, false)
	if err != nil {
		t.Error("can't submit proof", err)
	} else if sendData(config.TronPK, attachedData[:], pulseId, bytes32fromhex(addresses.SubscriptionId), false) {
		t.Error("this tx should fail because of invalid hash")
	}
}

func TestMint(t *testing.T) {
	var attachedData [1+32+32+20]byte

	attachedData[0] = 'm' // mint
	filluint(attachedData[:], 1, 123456789) // req id
	filluint(attachedData[:], 1+32, 10) // amount = 10
	filladdress(attachedData[:], 1+32+32, "9561C133DD8580860B6b7E504bC5Aa500f0f0103") // address

	proof := crypto.Keccak256Hash(attachedData[:])
	pulseId, err := signData(proof, 5, false)
	if err != nil {
		t.Error("can't submit proof", err)
	} else if !sendData(config.TronPK, attachedData[:], pulseId, bytes32fromhex(addresses.SubscriptionId), false) {
		t.Error("can't submit data")
	}
}

func TestChangeStatusFail(t *testing.T) {
	var attachedData [1+32+1]byte

	attachedData[0] = 'c' // change status
	filluint(attachedData[:], 1, 111111111) // req id = 111111111
	attachedData[1+32] = 1

	proof := crypto.Keccak256Hash(attachedData[:])
	pulseId, err := signData(proof, 5, false)
	if err != nil {
		t.Error("can't submit proof", err)
	}

	if sendData(config.TronPK, attachedData[:], pulseId, bytes32fromhex(addresses.SubscriptionId), false) {
		t.Error("request should fail")
	}
}

func TestChangeStatusOk(t *testing.T) {
	var dummyAddress [32]byte

	ibportTransactor, _ := helpers.NewTransactor(config.Endpoint, ibport.IBPortABI, addresses.IBPort)

	tx, err := ibportTransactor.Transact(config.TronPK, "createTransferUnwrapRequest", big.NewInt(10000000), dummyAddress)
    if err != nil {
        log.Fatal(err)
	}

	helpers.WaitForTx(*tx, config.Endpoint)
	requestCreatedEvent, err := helpers.WaitForEvent(*tx, "RequestCreated", config.Endpoint)
	if err != nil {
		t.Error("request failed", err)
	} else {
		reqId := new(big.Int)
		reqId.SetString(requestCreatedEvent["0"], 10)
		var attachedData [1+32+1]byte
		attachedData[0] = 'c' // change status
		filluint(attachedData[:], 1, reqId.Uint64())

		attachedData[1+32] = 2 // next status

		proof := crypto.Keccak256Hash(attachedData[:])
		pulseId, err := signData(proof, 5, false)
		if err != nil {
			t.Error("can't submit proof", err)
		}

		if !sendData(config.TronPK, attachedData[:], pulseId, bytes32fromhex(addresses.SubscriptionId), false) {
			t.Error("request failed", pulseId)
		}
	}
}

type Request struct {
	Id *big.Int
	HomeAddress common.Address
	ForeignAddress [32]byte
	Amount *big.Int
	Status uint8
}

func TestLock(t *testing.T) {
	var dummyAddress [32]byte

	amount := big.NewInt(12345)

	luportTransactor, _ := helpers.NewTransactor(config.Endpoint, luport.LUPortABI, addresses.LUPort)

	tx, err := luportTransactor.Transact(config.TronPK, "createTransferUnwrapRequest", amount, dummyAddress)
	if err != nil {
		t.Error(err)
	} else {
		receipt := helpers.WaitForTx(*tx, config.Endpoint)
		if receipt == nil {
			t.Error("can't create request " + *tx)
		} else {
			newRequestEvent, err := helpers.WaitForEvent(*tx, "NewRequest", config.Endpoint)
			if err != nil {
				t.Error(err)
			} else {
				swapId := new(big.Int)
				swapId.SetString(newRequestEvent["swapId"], 10)
				requests, _ := getRequestsQueue()
				requestStruct := findRequestById(requests, swapId)
	
				if requestStruct.Amount.Cmp(amount) != 0 {
					t.Error("failed")
				} else if requestStruct.Status != 1 {
					t.Error("unexpected status", requestStruct.Status)
				}
			}
		}
	}
}

func TestUnlock(t *testing.T) {
	var attachedData [1+32+32+20]byte

	attachedData[0] = 'u' // unlock
	filluint(attachedData[:], 1, 123456789) // req id
	filluint(attachedData[:], 1+32, 10) // amount = 10
	filladdress(attachedData[:], 1+32+32, "9561C133DD8580860B6b7E504bC5Aa500f0f0103") // address

	proof := crypto.Keccak256Hash(attachedData[:])
	pulseId , err:= signData(proof, 5, true)
	if err != nil {
		t.Error("can't submit proof", err)
	} else if !sendData(config.TronPK, attachedData[:], pulseId, bytes32fromhex(addresses.ReverseSubscriptionId), true) {
		t.Error("can't submit data")
	}
}

func getRequestsQueue() ([]Request, error) {
	var (
		id = new([]*big.Int)
		homeAddress = new([]common.Address)
		foreignAddress = new([][32]byte)
		amount = new([]*big.Int)
		status = new([]uint8)
	)
	out := &[]interface{}{
		id,
		homeAddress,
		foreignAddress,
		amount,
		status,
	}
	luportTransactor, _ := helpers.NewTransactor(config.Endpoint, luport.LUPortABI, addresses.LUPort)
	err := luportTransactor.ReadContract(config.TronPK, "getRequests", out)

	if err != nil {
		return nil, err
	}

	length := len(*id)

	if (length != len(*homeAddress)) || (length != len(*foreignAddress)) || (length != len(*amount)) || (length != len(*status)) {
		log.Fatal("invalid response")
	}

	ret := make([]Request, length)

	for i := 0; i < length; i++ {
		ret[i] = Request{Id: (*id)[i], HomeAddress: (*homeAddress)[i], ForeignAddress: (*foreignAddress)[i], Amount: (*amount)[i], Status: (*status)[i]}
	}

	return ret, nil
}

func findRequestById(requests []Request, id *big.Int) (*Request) {
	for i := 0; i < len(requests); i++ {
		r := requests[i]
		if r.Id.Cmp(id) == 0 {
			return &r
		}
	}
	return nil
}

func TestApprove(t *testing.T) {
	var attachedData [1+32]byte
	var dummyAddress [32]byte

	amount := big.NewInt(12345)

	luportTransactor, _ := helpers.NewTransactor(config.Endpoint, luport.LUPortABI, addresses.LUPort)

	tx, err := luportTransactor.Transact(config.TronPK, "createTransferUnwrapRequest", amount, dummyAddress)
	if err != nil {
		t.Error(err)
	} else {
		receipt := helpers.WaitForTx(*tx, config.Endpoint)
		if receipt == nil {
			t.Error("tx failed")
		}

		newRequestEvent, err := helpers.WaitForEvent(*tx, "NewRequest", config.Endpoint)
		if err != nil {
			t.Error(err)
		} else {
			swapId := new(big.Int)
			swapId.SetString(newRequestEvent["swapId"], 10)
			requests, err := getRequestsQueue()
			if err != nil {
				t.Error(err)
			} else {
				r := findRequestById(requests, swapId)
				if r == nil {
					t.Error("no request in queue")
				}
				attachedData[0] = 'a' // approve
				filluint(attachedData[:], 1, swapId.Uint64()) // req id
				proof := crypto.Keccak256Hash(attachedData[:])
				pulseId, err := signData(proof, 5, true)
				if err != nil {
					t.Error("can't submit proof", err)
				} else if !sendData(config.OraclePK[0], attachedData[:], pulseId, bytes32fromhex(addresses.ReverseSubscriptionId), true) {
					t.Error("can't submit data")
				}
				requests, _ = getRequestsQueue()
				r = findRequestById(requests, swapId)
				if r != nil {
					t.Error("request should not be in the queue")
				}
	
			}
		}
	}
}

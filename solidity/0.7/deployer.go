package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"rh_tests/api/ibport"
	"rh_tests/api/luport"
	"github.com/Gravity-Tech/gravity-core/common/contracts"

	"rh_tests/helpers"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func pubFromPK(pk string) common.Address {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func oraclesFromPK(oraclePK [5]string) [5]common.Address {
	var oracles [5]common.Address
	for i := 0; i < 5; i++ {
		oracles[i] = pubFromPK(oraclePK[i])
	}
	return oracles
}

func deployIBPort(addresses *helpers.DeployedAddresses, fromAddress common.Address, ethConnection *ethclient.Client, transactor *bind.TransactOpts, config *helpers.Config) {
	erc20MintableAddr, tx, tokenMintable, err := ibport.DeployToken(transactor, ethConnection, "TSTM", "TST mintable")

	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)

	addresses.ERC20Mintable = common.Bytes2Hex(erc20MintableAddr.Bytes())

	oracles := oraclesFromPK(config.OraclePK)
	nebulaAddr, tx, nebula, err := contracts.DeployNebula(transactor, ethConnection, 0, oracles[0], oracles[:], big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)
	addresses.Nebula = common.Bytes2Hex(nebulaAddr.Bytes())

	ibportAddress, tx, _, err := ibport.DeployIBPort(transactor, ethConnection, nebulaAddr, erc20MintableAddr)
	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)
	addresses.IBPort = common.Bytes2Hex(ibportAddress.Bytes())

	tx, err = tokenMintable.AddMinter(transactor, ibportAddress)
	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)

	tx, err = tokenMintable.Mint(transactor, fromAddress, big.NewInt(100000000000))
	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)

	tx, err = tokenMintable.Approve(transactor, ibportAddress, big.NewInt(100000000000))
	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), ethConnection, tx)

	tx, err = nebula.Subscribe(transactor, ibportAddress, 1, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := bind.WaitMined(context.Background(), ethConnection, tx)
	if err != nil {
		log.Fatal(err)
	}

	newSubEvent, err := nebula.NebulaFilterer.ParseNewSubscriber(*receipt.Logs[0])
	if err != nil {
		log.Fatal(err)
	}

	addresses.SubscriptionId = common.Bytes2Hex(newSubEvent.Id[:])
}

func deployLUPort(addresses *helpers.DeployedAddresses, fromAddress common.Address, ethConnection *ethclient.Client, transactor *bind.TransactOpts, config *helpers.Config) {
	erc20Addr, tx, token, err := luport.DeployToken(transactor, ethConnection, "TST", "TST")

	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), ethConnection, tx)

	addresses.ERC20 = common.Bytes2Hex(erc20Addr.Bytes())

	oracles := oraclesFromPK(config.OraclePK)
	nebulaReverseAddr, tx, nebula, err := contracts.DeployNebula(transactor, ethConnection, 0, oracles[0], oracles[:], big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), ethConnection, tx)
	addresses.NebulaReverse = common.Bytes2Hex(nebulaReverseAddr.Bytes())

	luportAddress, tx, _, err := luport.DeployLUPort(transactor, ethConnection, nebulaReverseAddr, erc20Addr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), ethConnection, tx)
	addresses.LUPort = common.Bytes2Hex(luportAddress.Bytes())

	tx, err = token.Mint(transactor, fromAddress, big.NewInt(100000000000))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), ethConnection, tx)

	tx, err = token.Mint(transactor, luportAddress, big.NewInt(100000000000))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), ethConnection, tx)

	tx, err = token.Approve(transactor, luportAddress, big.NewInt(100000000000))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), ethConnection, tx)

	tx, err = nebula.Subscribe(transactor, luportAddress, 1, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := bind.WaitMined(context.Background(), ethConnection, tx)
	if err != nil {
		log.Fatal(err)
	}

	newSubEvent, err := nebula.NebulaFilterer.ParseNewSubscriber(*receipt.Logs[0])
	if err != nil {
		log.Fatal(err)
	}

	addresses.ReverseSubscriptionId = common.Bytes2Hex(newSubEvent.Id[:])
}

func deployGravity(addresses *helpers.DeployedAddresses, fromAddress common.Address, ethConnection *ethclient.Client, transactor *bind.TransactOpts, config *helpers.Config) {
	oracles := oraclesFromPK(config.OraclePK)

	gravityAddress, tx, _, err := contracts.DeployGravity(transactor, ethConnection, oracles[:], big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	_, err = bind.WaitMined(context.Background(), ethConnection, tx)
	if err != nil {
		log.Fatal(err)
	}

	addresses.Gravity = common.Bytes2Hex(gravityAddress.Bytes())
}

func main() {
	var addresses helpers.DeployedAddresses

	config, err := helpers.LoadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("using endpoint", config.Endpoint)

	ethConnection, err := ethclient.DialContext(context.Background(), config.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(config.OraclePK[0])
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)


	gasPrice, err := ethConnection.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transactor := bind.NewKeyedTransactor(privateKey)
	transactor.GasLimit = uint64(8000000) // in units
	transactor.GasPrice = gasPrice

	deployGravity(&addresses, fromAddress, ethConnection, transactor, &config)
	deployIBPort(&addresses, fromAddress, ethConnection, transactor, &config)
	deployLUPort(&addresses, fromAddress, ethConnection, transactor, &config)

	log.Println(helpers.SaveAddresses(addresses))
}

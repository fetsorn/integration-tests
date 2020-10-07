package main

import (
	"math/big"
	"rh_tests/api/token"
	"rh_tests/api/nebula"
	"rh_tests/api/ibport"
	"rh_tests/api/luport"
	"rh_tests/helpers"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func oraclesAddressFromPK(oraclePK [5]string) ([5]common.Address) {
	var oracles [5]common.Address
	for i := 0; i < 5; i++ {
		oracles[i] = common.HexToAddress(helpers.HexFromPk(oraclePK[i]))
	}
	return oracles
}

func main() {
	var addresses helpers.DeployedAddresses
	config, err := helpers.LoadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("using endpoint", config.Endpoint)

	myAddress := helpers.HexFromPk(config.TronPK)
	oracles := oraclesAddressFromPK(config.OraclePK)

	tokenAddress, err := helpers.DeployContract(config.Endpoint, token.TokenABI, token.TokenBin, myAddress, config.TronPK, "Token")
	if err != nil {
		log.Fatal(err)
	}
	addresses.ERC20 = *tokenAddress

	nebulaAddress, err := helpers.DeployContract(config.Endpoint, nebula.NebulaABI, nebula.NebulaBin, myAddress, config.TronPK, "Nebula",
		uint8(0), common.HexToAddress(myAddress), oracles, big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}
	
	addresses.Nebula = *nebulaAddress

	nebulaReverseAddress, err := helpers.DeployContract(config.Endpoint, nebula.NebulaABI, nebula.NebulaBin, myAddress, config.TronPK, "Nebula",
		uint8(0), common.HexToAddress(myAddress), oracles, big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}
	
	addresses.NebulaReverse = *nebulaReverseAddress

	ibPortAddress, err := helpers.DeployContract(config.Endpoint, ibport.IBPortABI, ibport.IBPortBin, myAddress, config.TronPK, "IBPort",
		common.HexToAddress(*nebulaAddress), common.HexToAddress(*tokenAddress))
	if err != nil {
		log.Fatal(err)
	}
	
	addresses.IBPort = *ibPortAddress

	luPortAddress, err := helpers.DeployContract(config.Endpoint, luport.LUPortABI, luport.LUPortBin, myAddress, config.TronPK, "LUPort",
		common.HexToAddress(*nebulaAddress), common.HexToAddress(*tokenAddress))
	if err != nil {
		log.Fatal(err)
	}
	
	addresses.LUPort = *luPortAddress

	tokenTransactor, err := helpers.NewTransactor(config.Endpoint, token.TokenABI, addresses.ERC20)
	if err != nil {
		log.Fatal(err)
	}

	txMint, err := tokenTransactor.Transact(config.TronPK, "mint", 
		common.HexToAddress(helpers.HexFromPk(config.TronPK)), big.NewInt(1000000000000000000))
	if err != nil {
		log.Fatal(err)
	}

	txMint2, err := tokenTransactor.Transact(config.TronPK, "mint", 
		common.HexToAddress(addresses.LUPort), big.NewInt(1000000000000000000))
	if err != nil {
		log.Fatal(err)
	}

	txAddMinter, err := tokenTransactor.Transact(config.TronPK, "addMinter", 
		common.HexToAddress(addresses.IBPort))
	if err != nil {
		log.Fatal(err)
	}
	
	txApprove, err := tokenTransactor.Transact(config.TronPK, "approve", common.HexToAddress(addresses.IBPort), big.NewInt(1000000000000000000))
	if err != nil {
		log.Fatal(err)
	}

	txApprove2, err := tokenTransactor.Transact(config.TronPK, "approve", common.HexToAddress(addresses.LUPort), big.NewInt(1000000000000000000))
	if err != nil {
		log.Fatal(err)
	}


	nebulaTransactor, err := helpers.NewTransactor(config.Endpoint, nebula.NebulaABI, addresses.Nebula)
	if err != nil {
		log.Fatal(err)
	}

	nebulaReverseTransactor, err := helpers.NewTransactor(config.Endpoint, nebula.NebulaABI, addresses.NebulaReverse)
	if err != nil {
		log.Fatal(err)
	}

	txSubsribe, err := nebulaTransactor.Transact(config.TronPK, "subscribe", common.HexToAddress(addresses.IBPort), uint8(1), big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	txSubsribeReverse, err := nebulaReverseTransactor.Transact(config.TronPK, "subscribe", common.HexToAddress(addresses.LUPort), uint8(1), big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	if helpers.WaitForTx(*txMint, config.Endpoint) == nil {
		log.Fatal("error on mint()")
	}
	if helpers.WaitForTx(*txMint2, config.Endpoint) == nil {
		log.Fatal("error on mint() 2")
	}
	if helpers.WaitForTx(*txAddMinter, config.Endpoint) == nil {
		log.Fatal("error on addMinter()")
	}
	if helpers.WaitForTx(*txApprove, config.Endpoint) == nil {
		log.Fatal("error on approve()")
	}
	if helpers.WaitForTx(*txApprove2, config.Endpoint) == nil {
		log.Fatal("error on approve() 2")
	}
	if helpers.WaitForTx(*txSubsribe, config.Endpoint) == nil {
		log.Fatal("error on subscribe()")
	}

	if helpers.WaitForTx(*txSubsribeReverse, config.Endpoint) == nil {
		log.Fatal("error on subscribe()")
	}

	subscribeResult, err := helpers.GetEventByName(*txSubsribe, "NewSubscriber", config.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	addresses.SubscriptionId = subscribeResult["id"]

	subscribeReverseResult, err := helpers.GetEventByName(*txSubsribeReverse, "NewSubscriber", config.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	addresses.ReverseSubscriptionId = subscribeReverseResult["id"]

	log.Println(helpers.SaveAddresses(addresses))
}
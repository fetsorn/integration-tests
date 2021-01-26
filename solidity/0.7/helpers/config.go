package helpers

import (
	"encoding/json"
	"os"
	"io/ioutil"
)

type Config struct {
	Endpoint string    `json:"endpoint"`
	OraclePK [5]string `json:"oraclepk"`
}

type DeployedAddresses struct {
	Gravity, Nebula, NebulaReverse, ERC20, ERC20Mintable, IBPort, LUPort, SubscriptionId, ReverseSubscriptionId string
}

func SaveAddresses(addresses DeployedAddresses) (string) {
	file, _ := json.MarshalIndent(addresses, "", " ")
	ioutil.WriteFile("./addresses.json", file, 0644)
	return string(file)
}

func LoadAddresses() (DeployedAddresses, error) {
	var config DeployedAddresses

	configFile, err := os.Open("./addresses.json")
	defer configFile.Close()

	if err != nil {
		return DeployedAddresses{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return DeployedAddresses{}, err
	}

	return config, nil
}

func LoadConfiguration() (Config, error) {
	var config Config

	configFile, err := os.Open("./config.json")
	defer configFile.Close()

	if err != nil {
		return Config{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

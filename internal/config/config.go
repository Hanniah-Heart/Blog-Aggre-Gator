package config
import (
	"os"
	"fmt"
	"encoding/json"
	"path/filepath"
)


const configFileName = ".gatorconfig.json"
type cfg struct {
	Db_url	string
}

func Read () (cfg, error) {
	var Config cfg
	fmt.Printf("Config variable is %v\n", Config)
	home, err := os.UserHomeDir(); if err != nil {return Config, fmt.Errorf("Error grabbing home directory: %v", err)}
	fmt.Printf("home variable is: %v\n", home)
	path := filepath.Join(home, configFileName)
	fmt.Printf("path variable is: %v\n", path)
	data, err := os.ReadFile(path); if err != nil {return Config, fmt.Errorf("Error reading file: %v\n", err)}
	fmt.Printf("data variable is: %v\n", data)

	/*
	file, err := os.Open(path); if err != nil {return Config, fmt.Errorf("Error opening file: %v", err)}
	defer file.close()
	fmt.Printf("file variable is: %v\n", file)
	decoder := json.NewDecoder(file)
	fmt.Printf("decoder variable is: %v\n", decoder)
	//if err := decoder.Decode(&Config); err != nil {return Config, fmt.Errorf("Error decoding file: %v", err)}
	//fmt.Printf("Config variable is: %v\n", Config)
	*/

	if err := json.Unmarshal(data, &Config); err != nil {return Config, fmt.Errorf("Error Unmarshalling File: %v", err)}
	fmt.Printf("Config variable is: %v\n", Config)

	return Config, nil
}


// if err != nil {return Config, fmt.Errorf("Error : %v", err)}

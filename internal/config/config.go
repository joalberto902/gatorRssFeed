package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

//configFileName is the name of the file that contains the configuration
const configFileName string = ".gatorconfig.json"

//Config struct represents who is currently logged in and
//the connection credentials for the PostgreSQL database
type Config struct {
	DbURL 			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

//getConfigFilePath takes no argument and return the filepath to the config file
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(homeDir, configFileName)

	return configFilePath, nil
}

//Read function returns the config written on the file
func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err 
	}

	payload, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config 
	if err = json.Unmarshal(payload, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

//SetUser mathod changes the CurrentUserName of the config and writes it 
// the config file.
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	err := write(*cfg)
	return err
}

//write function writes a the Config struct to a file 
func write(cfg Config) error {	
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	f, err := os.Create(configFilePath)
	if err != nil {
		return err
	}

	defer f.Close()

	payload, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if _, err = f.Write(payload); err != nil {
		return err
	}
	
	return nil
}



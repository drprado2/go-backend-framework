package configs

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type Configuration struct {
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	HttpServerPort   string
	MigrationsPath   string
}

var loadedConfiguration *Configuration

func findConfigPath(currentPath string)(string, error){
	files, err := filepath.Glob(currentPath + "*")
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if mat, err := regexp.Match(`config$`, []byte(file)); mat == true && err == nil {
			return filepath.Abs(file + "/config.json")
		}
	}
	path, err := findConfigPath("../" + currentPath)
	if path == "" || err != nil {
		return "", fmt.Errorf("Dont find config path\nError: %s", err)
	}
	return path, nil
}

func GetConfig() (*Configuration, error) {
	if loadedConfiguration != nil {
		return loadedConfiguration, nil
	}

	configuration := Configuration{}

	path, err := findConfigPath("")
	if err != nil {
		return nil, err
	}

	if err := gonfig.GetConf(path, &configuration); err != nil {
		return nil, err
	}
	overrideWithEnvVars(&configuration)
	loadedConfiguration = &configuration
	return loadedConfiguration, nil
}

func overrideWithEnvVars(config *Configuration) {
	if value := os.Getenv("DatabaseHost"); len(value) > 0 {
		config.DatabaseHost = value
	}
	if value := os.Getenv("DatabasePort"); len(value) > 0 {
		iValue, _ := strconv.Atoi(value)
		config.DatabasePort = iValue
	}
	if value := os.Getenv("DatabaseUser"); len(value) > 0 {
		config.DatabaseUser = value
	}
	if value := os.Getenv("DatabasePassword"); len(value) > 0 {
		config.DatabasePassword = value
	}
	if value := os.Getenv("DatabaseName"); len(value) > 0 {
		config.DatabaseName = value
	}
	if value := os.Getenv("HttpServerPort"); len(value) > 0 {
		config.HttpServerPort = value
	}
	if value := os.Getenv("MigrationsPath"); len(value) > 0 {
		config.MigrationsPath = value
	}
}

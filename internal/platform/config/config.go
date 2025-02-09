package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Rocksus/fundtract/internal/model/constant"
	"github.com/Rocksus/fundtract/internal/platform/log"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int `yaml:"port"`
}

func InitConfig(confDir string) (*Config, error) {
	fileName := constant.AppName + ".setting.yaml"

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	confPath, err := searchConfig(wd, confDir, fileName)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var cfg Config

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	if err := cfg.loadSecrets(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func searchConfig(wd, confDir, fileName string) (string, error) {
	var confPath string
	for {
		confPath = filepath.Join(wd, confDir, fileName)

		stat, err := os.Stat(confPath)
		if os.IsExist(err) && stat.IsDir() {
			return "", fmt.Errorf("path %s is a directory", confPath)
		}

		if err == nil {
			return confPath, nil
		}

		prt := filepath.Dir(wd)
		if wd == prt {
			return "", fmt.Errorf("config file %s not found", fileName)
		}

		wd = prt
	}
}

func (c *Config) loadSecrets() error {
	err := godotenv.Load()
	if err != nil {
		log.Info(".env file not found, will use os env instead")
	}
	return nil
}

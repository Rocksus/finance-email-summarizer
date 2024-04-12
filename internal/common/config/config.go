package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Rocksus/fundtract/internal/model/constant"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func Init(confDir string, env string) (*Config, error) {
	fileName := constant.AppName + "." + env + ".yaml"

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

func (c *Config) loadSecrets() error {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg(".env file not found, will use os env instead")
	}

	for k, v := range c.Database {
		searchKey := strings.ToUpper(k)
		credentials := os.Getenv(searchKey + "_DB_CREDS")
		if credentials == "" {
			return fmt.Errorf("invalid secret value for %s_DB_CREDS", searchKey)
		}
		host := os.Getenv(searchKey + "_DB_HOST")
		if host == "" {
			// default to connecting via localhost
			host = "localhost"
		}
		// default port
		var port = 5432
		rawPort := os.Getenv(searchKey + "_DB_PORT")
		if rawPort != "" {
			port, err = strconv.Atoi(rawPort)
			if err != nil {
				return err
			}
		}

		v.Credentials = credentials
		v.Host = host
		v.Port = port
	}

	c.Log.LogLevel = os.Getenv("LOG_LEVEL")

	return nil
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

const (
	envKey = "ENV"
)

func GetEnv() string {
	env := strings.ToLower(os.Getenv(envKey))

	switch env {
	case
		constant.Development,
		constant.Staging,
		constant.Production:
		break

	default:
		if env == "" {
			env = constant.Development
		} else {
			log.Fatal().
				Str("env", env).
				Msg("[config] environment does not exist")
		}
	}
	return env
}

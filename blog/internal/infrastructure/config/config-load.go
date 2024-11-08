package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

const BasicConfigPath = "./configs/app/dev.toml"

func LoadConfig(val interface{}, absolutePath string, relativePath string) {
	relativeEnv := getEnv("CONFIG_PATH", "")
	if relativeEnv == "" {
		relativeEnv = BasicConfigPath
	}
	if relativePath != "" && getEnv("CONFIG_PATH", "") == "" {
		relativeEnv = relativePath
	}

	var pathConf string

	if absolutePath != "" {
		pathConf = filepath.Join(absolutePath, relativeEnv)
	} else {
		pathConf = relativeEnv
	}

	_, err := toml.DecodeFile(pathConf, val)
	if err != nil {
		panic(err)
	}

}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

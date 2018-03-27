package clialioss

import (
	"path/filepath"
	"os"
	"strings"
	sCli "github.com/sinlov/golang_utils/cli"
	sFiles "github.com/sinlov/golang_utils/files"
	"fmt"
	"errors"
	"io/ioutil"
)

const (
	defaultConf = "alioss.json"
	gitRepo     = "cli-ali-oss-v1-golang"
	gitUser     = "sinlov"
	gitHost     = "github.com"
)

// if not find config Path just try to use GOPATH code github.com/sinlov/cli-ali-oss-v1-golang/alioss.json
// and current path in dev is $project_path/build
// if code defaultConf and run root path not found, return ""
func AliOssConfigPath(custom string) (string, string, error) {
	configPath := defaultConf
	if custom != "" {
		configPath = custom
	}
	configFilePath := filepath.Join(sCli.CommandPath(), configPath)
	currentPath := sCli.CurrentDirectory()
	if sFiles.IsFileExist(configFilePath) {
		return configFilePath, currentPath, nil
	}
	sCli.FmtYellow("\nWarning!\nCan not find config file at path: %s\n", sCli.CommandPath())
	goPathEnv := os.Getenv("GOPATH")
	goPathEnvS := strings.Split(goPathEnv, ":")
	isFindDevConf := false
	for _, path := range goPathEnvS {
		codePath := filepath.Join(path, "src", gitHost, gitUser, gitRepo)
		futurePath := filepath.Join(codePath, configPath)
		currentPath = filepath.Join(codePath, "build")
		if sFiles.IsFileExist(futurePath) {
			configFilePath = futurePath
			isFindDevConf = true
			break
		}
	}
	if isFindDevConf {
		sCli.FmtCyan("just use dev config at path: %s\n", configFilePath)
		return configFilePath, currentPath, nil
	} else {
		errInfo := fmt.Sprintf("can not load config at path: %s\nExit 1\n", configFilePath)
		configFilePath = ""
		return configFilePath, currentPath, errors.New(errInfo)
	}
}

func ReadConfigFileContent(filePath string) ([]byte, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return nil, err
		} else {
			return nil, err
		}
	}

	fileJson, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(fileJson)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

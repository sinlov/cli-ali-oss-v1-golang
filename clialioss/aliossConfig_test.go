package clialioss

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	"path/filepath"
	"os"
	"strings"
	"fmt"
)

func TestAliOssConfigPath(t *testing.T) {
	convey.Convey("mock TestAliOssConfigPath", t, func() {
		// mock
		goPathEnv := os.Getenv("GOPATH")
		goPathEnvS := strings.Split(goPathEnv, ":")
		goFirstPath := goPathEnvS[0]
		customPath := defaultConf
		data := struct {
			custom          string
			wantConfigFile  string
			wantCurrentPath string
		}{
			custom:          customPath,
			wantConfigFile:  filepath.Join(goFirstPath, "src", gitHost, gitUser, gitRepo, customPath),
			wantCurrentPath: filepath.Join(goFirstPath, "src", gitHost, gitUser, gitRepo, "build"),
		}
		convey.Convey("do TestAliOssConfigPath", func() {
			// do
			configPathDefault, currentPathDefault, errDefault := AliOssConfigPath("")
			configPathCustom, currentPathCustom, errCustom := AliOssConfigPath(data.custom)
			convey.Convey("verify TestAliOssConfigPath", func() {
				// verify
				if errDefault != nil {
					t.Errorf("read default config error, %s", errDefault)
				} else {
					convey.So(configPathDefault, convey.ShouldEqual, data.wantConfigFile)
					convey.So(currentPathDefault, convey.ShouldEqual, data.wantCurrentPath)
				}
				if errCustom != nil {
					t.Errorf("read custom error, %s", errCustom)
				} else {
					convey.So(configPathCustom, convey.ShouldEqual, data.wantConfigFile)
					convey.So(currentPathCustom, convey.ShouldEqual, data.wantCurrentPath)
				}
			})
		})
	})
}

func TestReadConfigFileContent(t *testing.T) {
	convey.Convey("mock TestReadConfigFileContent", t, func() {
		// mock

		convey.Convey("do TestReadConfigFileContent", func() {
			// do
			convey.Convey("verify TestReadConfigFileContent", func() {
				// verify
				convey.So("", convey.ShouldEqual, "")
			})
		})
	})
}

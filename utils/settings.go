package utils

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
)

type Settings struct {
	Ip              string `json:"ip"`
	Port            string `json:"port"`
	DbHost          string `json:"dbHost"`
	DbPort          string `json:"dpPort"`
	DbName          string `json:"dbName"`
	DbUser          string `json:"dbUser"`
	DbPass          string `json:"dbPass"`
	WsReadBuffSize  int    `json:"wsReadBuffSize"`
	WsWriteBuffSize int    `json:"wsWriteBuffSize"`
}

const (
	SETTINGS_PATH   = "settings.json"
	SETTINGS_PREFIX = "."
)

var filePrefix string

func init() {
	flag.StringVar(&filePrefix, "configs", "", "use custom prefix for setting files")
}

func (s *Settings) LoadArgs(customPath ...string) error {
	settingsExt := Settings{}

	settingsRaw, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", filePrefix, SETTINGS_PATH))

	if err != nil {
		return err
	}

	err = json.Unmarshal(settingsRaw, s)

	if err != nil {
		return err
	}

	if len(customPath) > 0 {
		for _, path := range customPath {
			settingsExtRaw, err := ioutil.ReadFile(path)

			if err != nil {
				return err
			}

			err = json.Unmarshal(settingsExtRaw, &settingsExt)

			if err != nil {
				return err
			}

			s = &settingsExt
		}
	}

	return nil
}

func LoadSetting(arg interface{}, customPath ...string) error {
	if filePrefix == "" {
		filePrefix = SETTINGS_PREFIX
	}

	if len(customPath) > 0 {
		for _, path := range customPath {
			glog.Infof("Loading: %s", path)
			settingsExtRaw, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", filePrefix, path))

			if err != nil {
				return err
			}

			err = json.Unmarshal(settingsExtRaw, arg)

			if err != nil {
				return err
			}
			glog.Infof("Loaded: %s", path)
		}
	} else {
		settingsRaw, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", filePrefix, SETTINGS_PATH))

		if err != nil {
			return err
		}

		err = json.Unmarshal(settingsRaw, arg)

		if err != nil {
			return err
		}
	}

	return nil

}

func (s *Settings) SaveArgs(customPath ...string) error {
	return errors.New("NEI")
}

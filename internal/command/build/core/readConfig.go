package core

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ReadConfig 讀取設定
func ReadConfig(filename string) (result *Config, err error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	} // if

	result = &Config{}
	err = yaml.Unmarshal(bytes, result)

	if err != nil {
		return nil, err
	} // if

	err = result.Check()

	if err != nil {
		return nil, err
	} // if

	return result, nil
}

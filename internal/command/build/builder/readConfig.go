package builder

import (
	"io/ioutil"

	"Sheeter/internal/command/build/config"

	"gopkg.in/yaml.v3"
)

// ReadConfig 讀取編譯設定
func ReadConfig(filename string) (result *config.Config, err error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	} // if

	result = &config.Config{}
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

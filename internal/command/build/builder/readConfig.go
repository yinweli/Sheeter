package builder

import (
	"io/ioutil"

	"Sheeter/internal/command/build/config"

	"gopkg.in/yaml.v3"
)

// ReadConfig 讀取編譯設定
func ReadConfig(filename string, buildConfig *config.Config) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	} // if

	err = yaml.Unmarshal(file, buildConfig)

	if err != nil {
		return err
	} // if

	return nil
}

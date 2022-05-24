package builder

import (
	"io/ioutil"

	"Sheeter/internal/command/build/config"

	"gopkg.in/yaml.v3"
)

// ReadConfig 讀取編譯設定
func ReadConfig(filename string) (result *config.Config, errs []error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		errs := append(errs, err)
		return nil, errs
	} // if

	result = &config.Config{}
	err = yaml.Unmarshal(file, result)

	if err != nil {
		errs := append(errs, err)
		return nil, errs
	} // if

	ok, errs := result.Check()

	if ok == false {
		return nil, errs
	} // if

	return result, nil
}

package initier

import (
	os "os"
	json "encoding/json"
)

type Device struct {
	Speaker string
	Headset string
}

type Config struct {
	Device `json:"device"`
}

const (
	NIRCMDC_PATH = "./nircmdc.exe"
	CONFIG_PATH = "./config.json"
	CUR_PATH = "./cur"
)



/*func initNirCMD() {
	panic(&NoNirCMDFileError{})
}*/

func CheckInitiated() error {
	_, err := os.Stat(NIRCMDC_PATH)
	if err != nil {
		return &NoNirCMDFileError{}
	}

	_, err = os.Stat(CONFIG_PATH)
	if err != nil {
		return &NoConfigFileError{}
	}

	_, err = os.Stat(CUR_PATH)
	if err != nil {
		return &NoCURFileError{}
	}

	return nil
}

func InitConfig() {
	config := &Config{
		Device{
			Speaker: "Speaker",
			Headset: "Headset",
		},
	}

	file, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(CONFIG_PATH, file, 0644)
	if err != nil {
		panic(err)
	}
}

func InitCUR() {
	err := os.WriteFile(CUR_PATH, []byte("0"), 0644)
	if err != nil {
		panic(err)
	}
}

func InitAll() {
	err := CheckInitiated()
	if err != nil {
		if err.Error() == (&NoNirCMDFileError{}).Error() {
			panic(err)
		}

		//InitNirCMD()
		InitConfig()
		InitCUR()
	}
}
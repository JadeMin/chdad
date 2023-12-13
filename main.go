package main

import (
	os "os"
	json "encoding/json"
)

import (
	initier "chdad/initier"
	deviceManager "chdad/deviceManager"
)

var (
	config *initier.Config
)



func init() {
	initier.InitAll()

	file, _ := os.ReadFile(initier.CONFIG_PATH)
	err := json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
}

func main() {
	CUR := deviceManager.GetCUR()

	switch CUR {
		// if speaker (0)
		case "0": {
			deviceManager.Switch(config.Device.Headset)
			os.WriteFile(initier.CUR_PATH, []byte("1"), 0644)
		}

		// if headset (1)
		case "1": {
			deviceManager.Switch(config.Device.Speaker)
			os.WriteFile(initier.CUR_PATH, []byte("0"), 0644)
		}

		// if not valid
		default: {
			initier.InitCUR()
			panic(&initier.CURFileParseError{})
		}
	}
}
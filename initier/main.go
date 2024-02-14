package initier

import (
	fmt "fmt"
	os "os"
	filepath "path/filepath"
	json "encoding/json"
)

type Device struct {
	Speaker string `json:"speaker"`
	Headset string `json:"headset"`
}
type Config struct {
	Device *Device `json:"device"`
}

var cwd string
var (
	NIRCMD_PATH string
	CONFIG_PATH string
	CUR_PATH string
)

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	cwd = filepath.Dir(ex)


	CONFIG_PATH = filepath.Join(cwd, "config.json")
	CUR_PATH = filepath.Join(cwd, "cur")
}

func hasNirCMD(directory string, findC bool) (string, error) {
	checker := filepath.Join(directory, "nircmd.exe")
	_, err := os.Stat(checker)
	if err == nil {
		return checker, nil
	}
	
	if findC == true {
		checker = filepath.Join(directory, "nircmdc.exe")
		_, err = os.Stat(checker)
		if err == nil {
			return checker, nil
		}
	}

	return "", err
}

func getNirCMDPath() (string, error) {
	nircmd, err := hasNirCMD(os.Getenv("WINDIR"), false)
	if err == nil {
		return nircmd, nil
	}

	nircmd, err = hasNirCMD(cwd, true)
	if err == nil {
		return nircmd, nil
	}

	return "", err
}



func InitNirCMD() {
	var err error

	NIRCMD_PATH, err = getNirCMDPath()
	if err != nil {
		panic(&NoNirCMDFileError{})
	}
}

func InitConfig() {
	_, err := os.Stat(CONFIG_PATH)
	if err != nil {
		fmt.Printf("%v\n", &NoConfigFileWarning{})

		config := &Config{
			&Device{
				Speaker: "스피커",
				Headset: "헤드폰",
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
}

func InitCUR() {
	_, err := os.Stat(CUR_PATH)
	if err != nil {
		fmt.Printf("%v\n", &NoCURFileWarning{})

		err = os.WriteFile(CUR_PATH, []byte("0"), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func InitAll() {
	InitNirCMD()
	InitConfig()
	InitCUR()
}
package deviceManager

import (
	os "os"
	exec "os/exec"
)

import (
	initier "chdad/initier"
)



func getSwitchCommander(deviceName string, role string) *exec.Cmd {
	return exec.Command(initier.NIRCMD_PATH, "setdefaultsounddevice", deviceName, role)
}

func GetCUR() string {
	file, err := os.ReadFile(initier.CUR_PATH)
	if err != nil {
		panic(err)
	}

	return string(file)
}

func Switch(deviceName string) {
	getSwitchCommander(deviceName, "0").Run()
	getSwitchCommander(deviceName, "2").Run()
}
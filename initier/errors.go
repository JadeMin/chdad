package initier



type NoNirCMDFileError struct {}
func (e *NoNirCMDFileError) Error() string {
	return "nircmdc.exe file is not found"
}

type NoConfigFileError struct {}
func (e *NoConfigFileError) Error() string {
	return "config.json file is not found"
}

type NoCURFileError struct {}
func (e *NoCURFileError) Error() string {
	return "CUR file is not found"
}

type CURFileParseError struct {}
func (e *CURFileParseError) Error() string {
	return "CUR state file is corrupted or not valid"
}
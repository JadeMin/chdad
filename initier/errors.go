package initier



type NoNirCMDFileError struct {}
func (e *NoNirCMDFileError) Error() string {
	return "NirCMD is not found.\nMaybe it's not installed, or your daddy went to buy milk."
}

type NoConfigFileWarning struct {}
func (e *NoConfigFileWarning) Error() string {
	return "config.json file is not found. Creating new one..."
}

type NoCURFileWarning struct {}
func (e *NoCURFileWarning) Error() string {
	return "CUR state file is not found. Creating new one..."
}

type CURFileParseWarning struct {}
func (e *CURFileParseWarning) Error() string {
	return "CUR state file is corrupted. Creating new one..."
}
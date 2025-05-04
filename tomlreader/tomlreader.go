package tomlreader

import (
	"email_url_checker/helpers"

	"github.com/BurntSushi/toml"
)

// Method for to read a toml file details
func ReadTomlFile(pFilename string) interface{} {
	var lFileDetails interface{}

	_, lErr := toml.DecodeFile(pFilename, &lFileDetails)
	if lErr != nil {
		helpers.LogError(lErr, "TRRT01")
	}

	return lFileDetails
}

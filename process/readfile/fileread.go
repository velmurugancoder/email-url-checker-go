package readfile

import (
	"bufio"
	"email_url_checker/global"
	"email_url_checker/helpers"
	urlchecker "email_url_checker/process/Urlchecker"
	"os"
	"strings"
)

func Readurl(pflag string) {
	lFile, lErr := os.Open("urls.txt")
	if lErr != nil {
		helpers.LogError(lErr, "RFRU01")
	}
	defer lFile.Close()

	scanner := bufio.NewScanner(lFile)
	for scanner.Scan() {

		//remove the white space from the Url
		lUrl := strings.TrimSpace(scanner.Text())
		if lUrl == "" {
			continue
		}

		if pflag == "C" {
			global.G_UrlChan <- lUrl
		} else {
			urlchecker.CheckURL(lUrl)

		}

	}

}

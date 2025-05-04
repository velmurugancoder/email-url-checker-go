package urlchecker

import (
	"email_url_checker/helpers"
	"log"
	"net/http"
	"strings"
	"time"
)

func CheckURL(pUrl string) {

	if !strings.HasPrefix(pUrl, "http") {
		pUrl = "https://" + pUrl
	}

	lClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	lResponse, lErr := lClient.Get(pUrl)
	if lErr != nil {
		helpers.LogError(lErr, "UCU01")
		return
	}
	defer lResponse.Body.Close()

	if lResponse.StatusCode == 200 {
		log.Printf("✅ %s is UP (%s)\n", pUrl, lResponse.Status)
	} else if lResponse.StatusCode >= 300 && lResponse.StatusCode < 400 {
		log.Printf("⚠️  %s is REDIRECTED (%s)\n", pUrl, lResponse.Status)
	} else {
		log.Printf("❌ %s returned status %s\n", pUrl, lResponse.Status)
	}

}

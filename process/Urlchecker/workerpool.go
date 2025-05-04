package urlchecker

import (
	"email_url_checker/global"
	"sync"
)

func Workers(pWaitGroup *sync.WaitGroup) {
	defer pWaitGroup.Done()

	for lUrl := range global.G_UrlChan {
		CheckURL(lUrl)
	}

}

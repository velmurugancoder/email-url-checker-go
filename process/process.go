package process

import (
	"email_url_checker/global"
	urlchecker "email_url_checker/process/Urlchecker"
	"email_url_checker/process/readfile"
	"email_url_checker/tomlreader"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

func Process_withoutchannel() {

	global.G_ErrorArr = nil
	lTimeStart := time.Now()
	readfile.Readurl("")
	global.WithoutChannel = time.Since(lTimeStart)

}

func Process_Bychannel() {

	lTimeStart := time.Now()

	var lWaitGroup sync.WaitGroup
	global.G_UrlChan = make(chan string)

	lWorker_limit := getwokerlimit()
	for lWorker := 1; lWorker <= lWorker_limit; lWorker++ {
		lWaitGroup.Add(1)
		go urlchecker.Workers(&lWaitGroup)
	}

	readfile.Readurl("C")
	close(global.G_UrlChan)
	lWaitGroup.Wait()

	global.WithChannel = time.Since(lTimeStart)

}

func getwokerlimit() int {

	lWorkerConfig := tomlreader.ReadTomlFile("./toml/Go_rotuine.toml")
	lWorkerCount := fmt.Sprintf("%v", lWorkerConfig.(map[string]interface{})["Worker_Pool"])

	lWorker_limit, lErr := strconv.Atoi(lWorkerCount)
	if lErr != nil {
		log.Println("Error : ", lErr)
		lWorker_limit = 3
	}

	return lWorker_limit
}

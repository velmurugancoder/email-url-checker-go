package main

import (
	"email_url_checker/global"
	"email_url_checker/process"
	"log"
	"os"
	"time"
)

func main() {

	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file: %v", lErr)
	}
	defer lFile.Close()

	log.SetOutput(lFile)

	process.Process_Bychannel()
	log.Println("  ")
	process.Process_withoutchannel()

	log.Println("Duration with Gorutine And channel    -> ", global.WithChannel)
	log.Println("Duration without Gorutine And channel -> ", global.WithoutChannel)

	log.Println("❌ Error Details ❌ ")
	for _, lRec := range global.G_ErrorArr {
		log.Println("Indicator : ", lRec.Indicator, " | Function : ", lRec.Function, " | Line : ", lRec.Line, " | Error : ", lRec.Errorstr)
	}
	log.Println("❌ Error Details ❌ ")

}

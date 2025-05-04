package helpers

import (
	"email_url_checker/global"
	"fmt"
	"runtime"
)

func LogError(pErr error, pIndicator string) {
	if pErr == nil {
		return
	}

	lPc, _, lLine, lOk := runtime.Caller(1)
	lDetails := runtime.FuncForPC(lPc)
	lFuncName := "unknown"
	if lOk && lDetails != nil {
		lFuncName = lDetails.Name()
	}

	global.G_ErrorArr = append(global.G_ErrorArr, global.Errorstruct{
		Indicator: pIndicator,
		Function:  lFuncName,
		Line:      fmt.Sprint(lLine),
		Errorstr:  pErr.Error(),
	})

}

package utils

// var logMsgChan = make(chan string)
var logMsg string

func SendlogMsg(v string) {
	logMsg += v + "\n"
	// logMsgChan <- v
	// logMsg += <-logMsgChan + "\n"
	// logMsgChan.Close()
}

func GetLogMsg() string {
	return logMsg
}

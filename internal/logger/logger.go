package logger

import (
	"io"
	"log"
	"os"
)

// Initialize 初始化日誌
func Initialize(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	} // if

	multiWriter := io.MultiWriter(file, os.Stdout)
	flag := log.LstdFlags | log.Lshortfile
	logInfo = log.New(multiWriter, "Info", flag)
	logError = log.New(multiWriter, "Error", flag)
}

// Finalize 結束日誌
func Finalize() {
	file, ok := logInfo.Writer().(*os.File)

	if ok == false {
		return
	}

	err := file.Close()

	if err != nil {
		panic(err)
	} // if
}

// Info 輸出訊息日誌
func Info(message string) (result bool) {
	if logInfo != nil {
		logInfo.Println(message)
	}

	return true
}

// Error 輸出錯誤日誌
func Error(message string) (result bool) {
	if logError != nil {
		logError.Println(message)
	}

	return false
}

var logInfo *log.Logger  // 訊息日誌物件
var logError *log.Logger // 錯誤日誌物件

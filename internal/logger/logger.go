package logger

import (
	"fmt"
	"log"
	"os"

	"Sheeter/internal"
)

var logFile *os.File // 日誌檔案物件

// Initialize 初始化日誌
func Initialize() {
	file, err := os.OpenFile(internal.LoggerFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	logFile = file

	log.SetOutput(logFile)
	log.Println("----------------------------------------")
}

// Finalize 結束日誌
func Finalize() {
	err := logFile.Close()

	if err != nil {
		panic(err)
	}
}

// Info 輸出訊息日誌
func Info(message string) (result bool) {
	output("Info", message)
	return true
}

// Error 輸出錯誤日誌
func Error(message string) (result bool) {
	output("Error", message)
	return false
}

// output 輸出日誌
func output(level string, message string) {
	result := fmt.Sprintf("[%s] %s\n", level, message)

	log.Printf(result)
	fmt.Printf(result)
}

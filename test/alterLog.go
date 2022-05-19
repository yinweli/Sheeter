package test

import (
	"bytes"
	"log"
	"os"
)

// AlterLog 替代日誌, 用於測試時檢查日誌輸出用
type AlterLog struct {
	buffer bytes.Buffer
}

// Initialize 初始化處理
func (this *AlterLog) Initialize() {
	log.SetOutput(&this.buffer)
}

// Finalize 結束處理
func (this *AlterLog) Finalize() {
	log.SetOutput(os.Stderr)
}

// String 取得日誌字串, 同時會清空日誌內容
func (this *AlterLog) String() string {
	defer this.buffer.Reset()
	return this.buffer.String()
}

// NewAlterLog 建立替代日誌
func NewAlterLog() *AlterLog {
	alterLog := &AlterLog{buffer: bytes.Buffer{}}
	alterLog.Initialize()
	return alterLog
}

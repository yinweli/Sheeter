package reader

import "Sheeter/internal/command/build/config"

// Reader 讀取介面
type Reader interface {
	Read(global *config.Global, element *config.Element) error
}

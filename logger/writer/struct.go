package writer

import "github.com/joaosoft/clean-infrastructure/domain"

type Fallback struct {
	reader domain.FallbackReader
	writer domain.FallbackWriter
}

type WriterHandler func(message []byte) error

package log

import (
	"io"
)

type Writer struct {
	Writer io.Writer
}

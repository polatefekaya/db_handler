package log

import (
	"DatabaseHandler/internals/log"
)

type ICustomLog interface {
	NewSlog() *log.CustomLog
}

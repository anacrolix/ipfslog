package ipfslog

import (
	ipfs_go_log "github.com/ipfs/go-log"
	why_go_logging "github.com/whyrusleeping/go-logging"
)

type Level = why_go_logging.Level

const (
	Debug   = why_go_logging.DEBUG
	Info    = why_go_logging.INFO
	Warning = why_go_logging.WARNING
)

func SetAllLoggerLevels(l Level) {
	ipfs_go_log.SetAllLoggers(l)
}

func SetModuleLevel(module string, l Level) {
	why_go_logging.SetLevel(l, module)
}

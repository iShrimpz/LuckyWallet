package logger

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	log.SetOutput(io.MultiWriter(
		os.Stdout,
		&lumberjack.Logger{
			Filename:   "./logs/lucky_wallet.log",
			MaxSize:    500, // megabytes
			MaxBackups: 28,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
			LocalTime:  true,
		},
	))
}

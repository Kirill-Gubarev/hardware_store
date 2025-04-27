package utils

import (
	"log"
	"os"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LogInit() {
	_, err := os.Stat("logs")
	if err != nil {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

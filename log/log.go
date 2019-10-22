package log

import "log"

func Info(v ...interface{}) {
	log.Println(v...)
}

func Error(v ...interface{}) {
	log.Println(v...)
}

func Fatal(v ...interface{}) {
	log.Fatalln(v...)
}

func ErrorIf(err error) {
	if err != nil {
		Error(err)
	}
}

func FatalIf(err error) {
	if err != nil {
		Fatal(err)
	}
}

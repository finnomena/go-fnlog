package fnlog

import (
	"fmt"
	"io"
	"os"
)

func Info(args ...interface{}) {
	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}
	f := getFields(args[0])
	log(Infolvl, f, args...)
}

func Debug(args ...interface{}) {
	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}
	f := getFields(args[0])
	log(Debuglvl, f, args...)
}

func Error(args ...interface{}) {
	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}
	f := getFields(args[0])
	log(Errorlvl, f, args...)
}

func Panic(args ...interface{}) {
	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}
	f := getFields(args[0])
	log(Paniclvl, f, args...)
	panic(nil)
}

func Trace(args ...interface{}) {
	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}
	f := getFields(args[0])
	log(Tracelvl, f, args...)
}

func Warn(args ...interface{}) {
	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}
	f := getFields(args[0])
	log(Warnlvl, f, args...)
}

func Fatal(args ...interface{}) {
	if len(args) == 0 {
		io.WriteString(os.Stdout, fmt.Sprint(args...)+"\n")
	}
	f := getFields(args[0])
	log(Fatallvl, f, args...)
	os.Exit(1)
}

package fnlog

type logLevel int

const (
	Debuglvl logLevel = iota
	Tracelvl
	Infolvl
	Warnlvl
	Errorlvl
	Fatallvl
	Paniclvl
)

var levelTxt map[logLevel]string = map[logLevel]string{
	Debuglvl: "debug",
	Tracelvl: "trace",
	Infolvl:  "info",
	Warnlvl:  "warn",
	Errorlvl: "error",
	Fatallvl: "fatal",
	Paniclvl: "panic",
}

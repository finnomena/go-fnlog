package fnlog

type logLevel int

const (
	Tracelvl logLevel = iota
	Debuglvl
	Infolvl
	Warnlvl
	Errorlvl
	Fatallvl
	Paniclvl
)

var levelTxt map[logLevel]string = map[logLevel]string{
	Tracelvl: "trace",
	Debuglvl: "debug",
	Infolvl:  "info",
	Warnlvl:  "warn",
	Errorlvl: "error",
	Fatallvl: "fatal",
	Paniclvl: "panic",
}

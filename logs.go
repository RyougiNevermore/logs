package logs

type Element interface {
	Level() Level
	SetLevel(Level) Element
	Bytes() []byte
	String() string
	JSON() string
	Trace() Element
	TraceFile() Element
	TraceFileWithoutGoPath() Element
	AddFields(...F) Element
	Val() E
}

type F struct {
	K string
	V interface{}
}

type Log interface {
	Log(Element)
}

func Info(v ...interface{}) Element {
	e := newElement(InfoLevel, "", v...)
	return e
}

func Infof(format string, v ...interface{}) Element {
	e := newElement(InfoLevel, format, v...)
	return e
}

func Debug( v ...interface{}) Element {
	e := newElement(DebugLevel, "", v...)
	return e
}

func Debugf(format string, v ...interface{}) Element {
	e := newElement(DebugLevel, format, v...)
	return e
}

func Warn(v ...interface{}) Element {
	e := newElement(WarnLevel, "", v...)
	return e
}

func Warnf(format string, v ...interface{}) Element {
	e := newElement(WarnLevel, format, v...)
	return e
}

func Error(v ...interface{}) Element {
	e := newElement(ErrorLevel, "", v...)
	return e
}

func Errorf(format string, v ...interface{}) Element {
	e := newElement(ErrorLevel, format, v...)
	return e
}

package log

var std *Logger

func Init(options ...ConfigOption) {
	SetDefaultConf(options...)
}

// not safe for concurrent use
func ResetDefault(l *Logger) {
	std = l
	Sugar = std.Sugar

	Debug = std.Debug
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal

	Debugf = std.Debugf
	Infof = std.Infof
	Warnf = std.Warnf
	Errorf = std.Errorf
	DPanicf = std.DPanicf
	Panicf = std.Panicf
	Fatalf = std.Fatalf

	Debugw = std.Debugw
	Infow = std.Infow
	Warnw = std.Warnw
	Errorw = std.Errorw
	DPanicw = std.DPanicw
	Panicw = std.Panicw
	Fatalw = std.Fatalw

	DebugWithTrace = std.DebugWithTrace
	InfoWithTrace = std.InfoWithTrace
	WarnWithTrace = std.WarnWithTrace
	ErrorWithTrace = std.ErrorWithTrace
	DPanicWithTrace = std.DPanicWithTrace
	PanicWithTrace = std.PanicWithTrace
	FatalWithTrace = std.FatalWithTrace
	DebugfWithTrace = std.DebugfWithTrace
	InfofWithTrace = std.InfofWithTrace
	WarnfWithTrace = std.WarnfWithTrace
	ErrorfWithTrace = std.ErrorfWithTrace
	DPanicfWithTrace = std.DPanicfWithTrace
	PanicfWithTrace = std.PanicfWithTrace
	FatalfWithTrace = std.FatalfWithTrace
	DebugwWithTrace = std.DebugwWithTrace
	InfowWithTrace = std.InfowWithTrace
	WarnwWithTrace = std.WarnwWithTrace
	ErrorwWithTrace = std.ErrorwWithTrace
	DPanicwWithTrace = std.DPanicwWithTrace
	PanicwWithTrace = std.PanicwWithTrace
	FatalwWithTrace = std.FatalwWithTrace
}

func Default() *Logger {
	return std
}

var (
	Sugar = std.Sugar

	Debug  = std.Debug
	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal

	Debugf  = std.Debugf
	Infof   = std.Infof
	Warnf   = std.Warnf
	Errorf  = std.Errorf
	DPanicf = std.DPanicf
	Panicf  = std.Panicf
	Fatalf  = std.Fatalf

	Debugw  = std.Debugw
	Infow   = std.Infow
	Warnw   = std.Warnw
	Errorw  = std.Errorw
	DPanicw = std.DPanicw
	Panicw  = std.Panicw
	Fatalw  = std.Fatalw

	DebugWithTrace   = std.DebugWithTrace
	InfoWithTrace    = std.InfoWithTrace
	WarnWithTrace    = std.WarnWithTrace
	ErrorWithTrace   = std.ErrorWithTrace
	DPanicWithTrace  = std.DPanicWithTrace
	PanicWithTrace   = std.PanicWithTrace
	FatalWithTrace   = std.FatalWithTrace
	DebugfWithTrace  = std.DebugfWithTrace
	InfofWithTrace   = std.InfofWithTrace
	WarnfWithTrace   = std.WarnfWithTrace
	ErrorfWithTrace  = std.ErrorfWithTrace
	DPanicfWithTrace = std.DPanicfWithTrace
	PanicfWithTrace  = std.PanicfWithTrace
	FatalfWithTrace  = std.FatalfWithTrace
	DebugwWithTrace  = std.DebugwWithTrace
	InfowWithTrace   = std.InfowWithTrace
	WarnwWithTrace   = std.WarnwWithTrace
	ErrorwWithTrace  = std.ErrorwWithTrace
	DPanicwWithTrace = std.DPanicwWithTrace
	PanicwWithTrace  = std.PanicwWithTrace
	FatalwWithTrace  = std.FatalwWithTrace
)

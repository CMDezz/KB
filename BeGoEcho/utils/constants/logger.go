package constants

const (
	LogBackEndPrefix          string = "[BACKEND]"
	LogBackEndModulePrefix    string = "[KB_V1]"
	LogBackEndMainInfoPrefix  string = "[Main_Info]"
	LogBackEndMainErrorPrefix string = "[Main_Error]"
	LogBackEndMainWarnPrefix  string = "[Main_Warn]"
	LogBackEndMainFatalPrefix string = "[Main_Fatal]"
	LogBackEndMainPanicPrefix string = "[Main_Panic]"
	LogBackEndMainDebugPrefix string = "[Main_Debug]"
)

const (
	LogInfoPrefix  = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainInfoPrefix
	LogErrorPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainErrorPrefix
	LogWarnPrefix  = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainWarnPrefix
	LogFatalPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainFatalPrefix
	LogPanicPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainPanicPrefix
	LogDebugPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainDebugPrefix
)

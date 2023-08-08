package logx

import "time"

const (
	KeyLevel                  = "logger.level"
	KeyWithCaller             = "logger.with.caller"
	KeyFormatterType          = "logger.formatter.type"
	KeyTimestampEnable        = "logger.formatter.timestamp.enable"
	KeyTimestampFormat        = "logger.formatter.timestamp.format"
	KeyPadLevelText           = "logger.formatter.pad.level.text"
	KeyStdoutEnable           = "logger.output.stdout.enable"
	KeyFileEnable             = "logger.output.file.enable"
	KeyFileDir                = "logger.output.file.dir"
	KeyFileName               = "logger.output.file.name"
	KeyFileRotationEnable     = "logger.output.file.rotation.enable"
	KeyFileRotationMaxSize    = "logger.output.file.rotation.max.size"
	KeyFileRotationMaxBackups = "logger.output.file.rotation.max.backups"
	KeyFileRotationMaxAge     = "logger.output.file.rotation.max.age"
	KeyFileRotationCompress   = "logger.output.file.rotation.compress"

	DatetimeFormat = time.DateTime + ".000"
)

var (
	defCfg = &Config{
		Level:                        "info",
		WithCaller:                   true,
		FormatterType:                "text",
		FormatterTimestampEnable:     true,
		FormatterTimestampFormat:     DatetimeFormat,
		FormatterPadLevelText:        true,
		OutputStdoutEnable:           true,
		OutputFileEnable:             false,
		OutputFileDir:                "/var/log/sdp",
		OutputFileName:               "sdp.log",
		OutputFileRotationEnable:     false,
		OutputFileRotationMaxSize:    100,
		OutputFileRotationMaxBackups: 3,
		OutputFileRotationMaxAge:     30,
		OutputFileRotationCompress:   false,
	}

	currCfg *Config
)

type Config struct {
	Level      string
	WithCaller bool

	// Text formatter configuration
	FormatterType            string
	FormatterTimestampEnable bool
	FormatterTimestampFormat string
	FormatterPadLevelText    bool

	// stdout output configuration.
	OutputStdoutEnable bool

	// file output configuration.
	OutputFileEnable bool
	OutputFileDir    string
	OutputFileName   string

	// file rotation configuration
	OutputFileRotationEnable     bool
	OutputFileRotationMaxSize    int
	OutputFileRotationMaxBackups int
	OutputFileRotationMaxAge     int
	OutputFileRotationCompress   bool
}

func GetDefaultConfigs() *Config {
	return defCfg
}

func GetConfigs() *Config {
	return currCfg
}

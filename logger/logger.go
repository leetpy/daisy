package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

func InitLogger() {
	rawJSON := []byte(`{
		"level":"debug",
		"encoding":"json",
		"outputPaths": ["stdout", "api.log"],
		"errorOutputPaths": ["stderr", "error.log"],
		"initialFields":{},
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
}

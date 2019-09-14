package util

import (
	graylog "github.com/gemnasium/logrus-graylog-hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	formatter "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"runtime"
	"strings"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Out = os.Stdout
	logFormatter := new(formatter.TextFormatter)
	logFormatter.ForceColors = true
	logFormatter.ForceFormatting = true

	logger.Formatter = logFormatter
	logger.SetLevel(logrus.TraceLevel + 1)

	gelf := os.Getenv("GELF_UDP")
	if gelf != "" {
		hook := graylog.NewAsyncGraylogHook(gelf, map[string]interface{}{"project": "spectacle"})
		logger.AddHook(hook)
		logger.Info("Graylog registered")
	}

}

func GetLogger(module ...string) *logrus.Entry {
	switch len(module) {
	case 0:
		return logger.WithField("prefix", "default")
	case 1:
		return logger.WithField("prefix", module[0])
	default:
		sb := strings.Builder{}
		len := len(module)
		for i, val := range module {
			sb.WriteString(val)
			if i < (len - 1) {
				sb.WriteByte(',')
			}
		}
		return logger.WithField("prefix", sb.String())
	}
}

func LocalLogger(entry *logrus.Entry, ctx *gin.Context) (logger *logrus.Entry) {
	logger = entry
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		details := runtime.FuncForPC(pc)
		logger = logger.
			WithField("caller_function", details.Name())
	}
	if ctx != nil {
		logIdVal := ctx.Value("LogId")
		if logIdStr, ok := logIdVal.(string); ok {
			logger = logger.WithField("SPEC-Log-Id", logIdStr)
		}
	}
	return logger
}

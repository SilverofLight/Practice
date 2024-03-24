package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core) // logger 全局定义
	sugarLogger = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"error fetching url..",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		logger.Info(
			"Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}
}

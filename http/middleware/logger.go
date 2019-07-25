package middleware

import (
	"fmt"
	"ginDemo/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.LOG_FILE_PATH
	logFileName := config.LOG_FILE_NAME

	fileName := path.Join(logFilePath, logFileName)

	handleFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = handleFile
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		endTime := time.Now()
		fmt.Println()
		latencyTime := endTime.Sub(startTime) // 执行时间

		reqMethod := context.Request.Method
		reqUri := context.Request.RequestURI
		statusCode := context.Writer.Status()
		clientIP := context.ClientIP()

		//logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//	statusCode,
		//	latencyTime,
		//	clientIP,
		//	reqMethod,
		//	reqUri,
		//)
		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency_time": latencyTime,
			"client_ip": clientIP,
			"req_method": reqMethod,
			"req_uri": reqUri,
		}).Info()
	}
}

func LoggerToMongo() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

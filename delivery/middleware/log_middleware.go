package middleware

import (
	"final-project-kelompok-1/utils/common"
	modelutil "final-project-kelompok-1/utils/common/model_util"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type LogMiddleware interface {
	LogRequest() gin.HandlerFunc
}

type logMiddleware struct {
	logService common.MyLogger
}

func (l *logMiddleware) LogRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := l.logService.InitializeLogger(); err != nil {
			log.Fatal("err : ", err.Error())
		}
		t := time.Now()

		logString := modelutil.RequestLog{
			AccessTime: t,
			Latency:    time.Since(t),
			ClientIP:   ctx.ClientIP(),
			Method:     ctx.Request.Method,
			Code:       ctx.Writer.Status(),
			Path:       ctx.Request.URL.Path,
			UserAgent:  ctx.Request.UserAgent(),
		}

		switch {
		case ctx.Writer.Status() >= 500:
			l.logService.LogFatal(logString)
		case ctx.Writer.Status() >= 400:
			l.logService.LogWarn(logString)
		default:
			l.logService.LogInfo(logString)
		}
	}
}

func NewLogMiddleware(logService common.MyLogger) LogMiddleware {
	return &logMiddleware{logService: logService}
}

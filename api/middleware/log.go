package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/longhaoteng/wineglass/logger"
)

type Log struct{}

func (l *Log) Init() ([]gin.HandlerFunc, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return []gin.HandlerFunc{
		func(c *gin.Context) {
			start := time.Now()
			c.Next()
			latency := fmt.Sprintf("%v", time.Now().Sub(start))
			path := c.Request.URL.Path
			raw := c.Request.URL.RawQuery
			if raw != "" {
				path = path + "?" + raw
			}
			statusCode := c.Writer.Status()
			clientIP := c.ClientIP()
			referer := c.Request.Referer()
			dataLength := c.Writer.Size()
			if dataLength < 0 {
				dataLength = 0
			}

			entry := logger.GetEntry().WithFields(
				logrus.Fields{
					"hostname":    hostname,
					"code":        statusCode,
					"latency":     latency,
					"client_ip":   clientIP,
					"method":      c.Request.Method,
					"path":        path,
					"referer":     referer,
					"data_length": dataLength,
				},
			)

			if len(c.Errors) > 0 {
				entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				if statusCode >= http.StatusInternalServerError {
					entry.Error()
				} else if statusCode >= http.StatusBadRequest {
					entry.Warn()
				} else {
					entry.Info()
				}
			}
		},
		gin.Recovery(),
	}, nil
}

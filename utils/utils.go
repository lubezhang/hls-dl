package utils

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog"
)

func Logger() *zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	// consoleWriter.FormatLevel = func(i interface{}) string {
	// 	return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	// }

	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	// multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)
	logger := zerolog.New(consoleWriter).With().Timestamp().Logger()
	return &logger
}

func LoggerDebug(msg string) {
	Logger().Debug().Msg(msg)
}
func LoggerInfo(msg string) {
	Logger().Info().Msg(msg)
}
func LoggerError(msg string) {
	Logger().Error().Msg(msg)
}

func GetBaseUrl(srcUrl string) string {
	u, _ := url.Parse(srcUrl)
	return u.Scheme + "://" + u.Host
}

// 获取临时文件目录
func GetDownloadTmpDir() string {
	dir, _ := os.Getwd()
	return path.Join(dir, CONST_BASE_DATA_DIR, CONST_BASE_TMP_DIR)
}

// 获取下载文件目录
func GetDownloadDataDir() string {
	dir, _ := os.Getwd()
	return path.Join(dir, CONST_BASE_DATA_DIR)
}

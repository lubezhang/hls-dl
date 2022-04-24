package pulldlr

import (
	"flag"

	"github.com/lubezhang/pulldlr/utils"
	"github.com/rs/zerolog"
)

func Command() {
	urlFlag := flag.String("u", "", "m3u8下载地址(http(s)://url/xx/xx/index.m3u8)")
	// nFlag := flag.Int("n", 16, "下载线程数(max goroutines num)")
	oFlag := flag.String("o", "output", "自定义文件名(默认为output)")

	flag.Parse()

	m3u8Url := *urlFlag

	if m3u8Url == "" {
		utils.LoggerError("请输入m3u8地址")
		return
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ShowProtocolInfo(m3u8Url)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	dl, _ := New(m3u8Url)
	dl.SetOpts(DownloaderOption{
		FileName: *oFlag,
	})
	dl.Start()
}

package pulldlr

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestDownloader(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	dl, _ := New("https://cache28.cdanan.xyz/jianghu/vipyk/XNTg1NTg0MDYwOA==.m3u8")
	dl.SetOpts(DownloaderOption{
		FileName: "11.mp4",
	})
	dl.Start()
}

func TestShowProtocolInfo(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ShowProtocolInfo("https://cache28.cdanan.xyz/jianghu/vipyk/XNTg1NTg0MDYwOA==.m3u8")
}

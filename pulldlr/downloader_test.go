package pulldlr

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestDownloader(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	dl, _ := New("https://qq.sd-play.com/20220405/4Si6DIev/hls/index.m3u8")
	// dl, _ := New("https://cache28.cdanan.xyz/jianghu/vipyk/XNTg1NTg0MDYwOA==.m3u8")
	dl.SetOpts(DownloaderOption{
		FileName: "11.mp4",
	})
	dl.Start()
}

func TestShowProtocolInfo(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ShowProtocolInfo("https://qq.sd-play.com/20220405/4Si6DIev/hls/index.m3u8")
}

func TestStartMergeFile(t *testing.T) {
	StartMergeFile()
}

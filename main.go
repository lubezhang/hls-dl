package main

import (
	"github.com/lubezhang/pulldlr/pulldlr"
	"github.com/rs/zerolog"
)

func main() {
	// url := "https://cache28.cdanan.xyz/jianghu/vipyk/XNTg1NTg0MDYwOA==.m3u8"
	url := "https://cache.m3u8.suoyo.cc/duoduo/20220415/354bc72d091d8f1870c5d32442a0664e.m3u8?st=5q0Rwy7G9voBMOYmfaRk4Q&e=1650022114"

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	pulldlr.ShowProtocolInfo(url)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	dl, _ := pulldlr.New(url)
	dl.SetOpts(pulldlr.DownloaderOption{
		FileName: "11.mp4",
	})
	dl.Start()

}

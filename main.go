package main

import (
	"github.com/lubezhang/pulldlr/pulldlr"
	"github.com/rs/zerolog"
)

func main() {
	// url := "https://cache28.cdanan.xyz/jianghu/vipyk/XNTg1NTg0MDYwOA==.m3u8"
	// url := "https://cache.m3u8.bdcdss.com/duoduo/20220417/d09e76b6de22ff8eb7471ea908ce2c00.m3u8?st=LUGo_nfwf3c6akq0XC5N4w&e=1650193965" // 文件头部有多余数据
	// url := "https://v3.dious.cc/20220331/NGxXAbhN/2000kb/hls/index.m3u8"
	// url := "https://cache.m3u8.shenglinyiyang.cn/duoduo/20220416/14f30f497e7899051c2fbf916b7c2b94.m3u8?st=ou6ScY3DyGiCuTNZ_tclow&e=1650103315"
	url := "https://qq.sd-play.com/20220405/4Si6DIev/hls/index.m3u8" // 加密

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	pulldlr.ShowProtocolInfo(url)

	// zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// dl, _ := pulldlr.New(url)
	// dl.SetOpts(pulldlr.DownloaderOption{
	// 	FileName: "11.mp4",
	// })
	// dl.Start()

}

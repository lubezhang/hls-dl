package pulldlr

import (
	"fmt"

	"github.com/lubezhang/hls-parse/protocol"
	"github.com/lubezhang/hls-parse/types"
	"github.com/lubezhang/pulldlr/utils"
)

// 显示协议信息
func ShowProtocolInfo(url string) {
	baseUrl := utils.GetBaseUrl(url)

	data1, _ := utils.HttpGetFile(url)
	strDat1 := string(data1)
	hls, err := protocol.Parse(&strDat1, baseUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	if hls.IsMaster() {
		fmt.Println("******* 主文件 *******")
		fmt.Println("Stream 数量:", len(hls.ExtStreamInf))

		if len(hls.ExtStreamInf) == 0 {
			return
		}

		for idx, extStreamInf := range hls.ExtStreamInf {
			fmt.Printf("Stream%d 分辨率:%s \n", idx+1, extStreamInf.Resolution)
		}
		fmt.Println("******* 主文件 *******")
		fmt.Println("")
		data2, _ := utils.HttpGetFile(hls.ExtStreamInf[0].Url)
		strData2 := string(data2)
		hls2, _ := protocol.Parse(&strData2, baseUrl)
		if hls2.IsVod() {
			showProtocolVod(hls2)
		}

	} else if hls.IsVod() {
		showProtocolVod(hls)
	} else {
		fmt.Println("没有协议")
	}
}

func showProtocolVod(hls types.HLS) {
	fmt.Println("")
	fmt.Println("******* VOD文件 *******")
	fmt.Println("分片数量：", len(hls.Extinf))
	fmt.Println("是否加密：", len(hls.Extkey) > 0, len(hls.Extkey))
	fmt.Println("******* VOD文件 *******")
	fmt.Println("")
}

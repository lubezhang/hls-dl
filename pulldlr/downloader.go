package pulldlr

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/imdario/mergo"
	"github.com/lubezhang/hls-parse/protocol"
	"github.com/lubezhang/hls-parse/types"
	"github.com/lubezhang/pulldlr/utils"
)

const CONST_BASE_SLICE_FILE_EXT = ".ts" // 分片文件扩展名

func New(url string) (result *Downloader, err error) {
	result = &Downloader{
		m3u8Url: url,
	}
	return result, nil
}

// 下载器参数
type DownloaderOption struct {
	FileName string // 文件名
}

// 下载器
type Downloader struct {
	m3u8Url    string            // m3u8文件地址
	hlsMaster  *types.HLS        // 主协议文件对象
	hlsSlave   *types.HLS        // 子协议。vod、live或event
	wg         sync.WaitGroup    // 并发线程管理容器
	opts       DownloaderOption  // 下载器参数
	cache      DownloadCacheData // 下载数据管理器
	sliceCount int               // 下载进度，完成文件合并的分片数量
}

// 设置参数
func (dl *Downloader) SetOpts(opts1 DownloaderOption) {
	dl.opts = opts1
}

// 开始下载m3u8文件
func (dl *Downloader) Start() {
	dl.sliceCount = 0
	// 设置默认参数
	defaultOpts := DownloaderOption{
		FileName: time.Now().Format("2006-01-02$15:04:05") + ".mp4", // 生成临时文件名
	}
	mergo.MergeWithOverwrite(&defaultOpts, dl.opts) // 合并自定义和默认参数
	utils.LoggerInfo(">>>>>>> 下载视频:" + defaultOpts.FileName)
	dl.SetOpts(defaultOpts)

	_, err := dl.getMediaVod()
	if err == nil {
		// dl.startDownload()
		dl.startMergeFile()
	} else {
		utils.LoggerInfo(err.Error())
		return
	}
	utils.LoggerInfo("<<<<<<< 下载视频完成:" + defaultOpts.FileName)
}

// 合并视频分片文件到一个视频文件中
func (dl *Downloader) startMergeFile() {
	sliceTotal := len(dl.hlsSlave.Extinf)
	progress := dl.sliceCount / sliceTotal
	utils.LoggerInfo("******* 视频下载进度：" + strconv.Itoa(dl.sliceCount) + " = " + strconv.Itoa(progress))

	for {
		if dl.sliceCount == sliceTotal {
			break
		}
		sliceFilePath := dl.getTmpFilePath(strconv.Itoa(dl.sliceCount))
		_, err1 := os.Stat(sliceFilePath)
		if err1 != nil {
			break
		}
		file, err2 := os.OpenFile(sliceFilePath, os.O_RDONLY, 0666)
		if err2 != nil {
			break
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		buf, _ := io.ReadAll(reader)
		vodFile, _ := os.OpenFile(dl.getVodFilePath(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		vodFile.Write(buf)
		defer vodFile.Close()

		dl.sliceCount = dl.sliceCount + 1
	}

	// ticker := time.NewTicker(1 * time.Second)
	// count := 0
	// for {
	// 	if count > 10 {
	// 		break
	// 	}
	// 	fmt.Println("ticker ticker ticker ...")
	// 	count++
	// 	<-ticker.C
	// }
	// ticker.Stop()
}

func (dl *Downloader) startDownload() {
	hls := dl.hlsSlave
	if hls.IsVod() && len(hls.Extinf) > 0 {
		dl.wg.Add(len(hls.Extinf)) // 初始化并发线程计数器
	}
	dl.setDecryptKey()
	dl.setDwnloadCache()
	dl.startDownloadVod()
	dl.wg.Wait()
}

// 开始下载Vod文件
func (dl *Downloader) startDownloadVod() {
	for i := 0; i < 10; i++ {
		go dl.downloadVodFile()
		// dl.downloadVodFile()
	}
}

// 将视频片放到数据下载管理器中
func (dl *Downloader) setDwnloadCache() {
	hls := dl.hlsSlave
	if len(hls.Extinf) == 0 {
		return
	}

	var list []DownloadData
	for idx, extinf := range hls.Extinf {
		var decryptKey = ""
		if extinf.EncryptIndex > 0 {
			decryptKey = hls.Extkey[extinf.EncryptIndex].Key
		}
		list = append(list, DownloadData{
			Index:        idx,
			Key:          utils.GetMD5(extinf.Url),
			Title:        extinf.Title,
			Url:          extinf.Url,
			DownloadPath: dl.getTmpFilePath(strconv.Itoa(idx)),
			EncryptKey:   decryptKey,
		})
	}
	dl.cache.Push(list)
}

func (dl *Downloader) downloadVodFile() {
	for {
		data, err := dl.cache.Pop()
		if err != nil {
			break
		}
		utils.DownloadeSliceFile(data.Url, data.DownloadPath, data.EncryptKey)
		// dl.startMergeFile()
		time.Sleep(1000 * time.Millisecond)
		defer dl.wg.Done()
	}
}

// 解析vod类型的协议
func (dl *Downloader) getMediaVod() (result types.HLS, err error) {
	utils.LoggerInfo("获取Vod协议文件对象")
	baseUrl := utils.GetBaseUrl(dl.m3u8Url)

	data1, _ := utils.HttpGetFile(dl.m3u8Url)
	strDat1 := string(data1)
	hls, _ := protocol.Parse(&strDat1, baseUrl)

	if hls.IsMaster() && len(hls.ExtStreamInf) > 0 {
		dl.hlsMaster = &hls
		data2, _ := utils.HttpGetFile(hls.ExtStreamInf[0].Url)
		strData2 := string(data2)
		hls2, _ := protocol.Parse(&strData2, baseUrl)
		if hls.IsVod() {
			dl.hlsSlave = &hls2
			result = hls2
		}
	} else if hls.IsVod() {
		dl.hlsSlave = &hls
		result = hls
	} else {
		return result, errors.New("没有视频回放文件")
	}
	err = nil
	return
}

// 通过链接获取加密密钥，并将密钥填充到加密数据结构中
func (dl *Downloader) setDecryptKey() {
	hls := dl.hlsSlave
	if len(hls.Extkey) == 0 {
		return
	}

	var keys []types.HlsExtKey
	for _, extkey := range hls.Extkey {
		if extkey.Method == "AES-128" {
			tmp := extkey
			data, _ := utils.HttpGetFile(extkey.Uri)
			tmp.Key = string(data)
			keys = append(keys, tmp)
		}
	}

	hls.Extkey = keys
}

func (dl *Downloader) getTmpFilePath(fileName string) string {
	return path.Join(utils.GetDownloadTmpDir(), utils.GetMD5(dl.opts.FileName), fileName+CONST_BASE_SLICE_FILE_EXT)
}
func (dl *Downloader) getVodFilePath() string {
	return path.Join(utils.GetDownloadDataDir(), dl.opts.FileName)

}

package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/lubezhang/hls-parse/common"
)

func HttpGetFile(url string) ([]byte, error) {
	LoggerDebug("HttpGetFile: " + url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// 下载http链接的分片文件到本地磁盘，如果是加密文件请传递加密密钥
// @param url string 链接地址
func DownloadeSliceFile(url string, filePath string, decryptKey string) (result string, err error) {
	LoggerDebug("DownloadeSliceFile: " + filePath)
	err1 := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err1 != nil {
		fmt.Println(err)
	}

	// file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	file, err := CreateTmpFile()
	if err != nil {
		LoggerError(err.Error())
		return filePath, err
	}

	write := bufio.NewWriter(file)
	defer file.Close()

	var decryptData []byte
	data, err := HttpGetFile(url)

	if decryptKey == "" {
		decryptData = data
	} else {
		decryptData, err = common.AesDecrypt(data, decryptKey)
		if err != nil {
			return filePath, err
		}
	}

	write.Write(decryptData)
	write.Flush()

	CopyFile(file.Name(), filePath)
	os.Remove(file.Name())

	return filePath, nil
}

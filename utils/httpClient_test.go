package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpGetFile(t *testing.T) {
	assetObj := assert.New(t)

	httpUrl := "http://mp.zhisland.com/wmp/user/news/2109180006/consult/viewpoint?page=1&t=1649902826233"
	data, _ := HttpGetFile(httpUrl)
	strData := string(data)
	assetObj.Equal(strings.Contains(strData, "errorCode"), true)
}

func TestGetMD5(t *testing.T) {
	assetObj := assert.New(t)
	assetObj.Equal(GetMD5("123456"), "e10adc3949ba59abbe56e057f20f883e")
}

package pulldlr

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var testMergeCount = 0
var wg = sync.WaitGroup{}

func StartMergeFile() {
	fmt.Println("StartMergeFile")
	wg.Add(2)
	// go mergeFile("1")
	go mergeFile("2")

	wg.Wait()
}

func mergeFile(prefixName string) {
	if testMergeCount >= 3 {
		wg.Done()
		func1()
		return
	}
	for i := 0; i < 3; i++ {
		wData := []byte("test" + strconv.Itoa(testMergeCount) + "-" + strconv.Itoa(i))
		fmt.Println("mergeFile: ", string(wData))
		// ioutil.WriteFile("../data/"+string(wData)+".txt", wData, 0644)
		vodFile, _ := os.OpenFile("../data/"+prefixName+"mergeFile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		vodFile.Write(wData)
		vodFile.Close()
		time.Sleep(time.Second * 1)
	}

	testMergeCount++
	mergeFile(prefixName)
}

func func1() {
	time.Sleep(time.Second * 14)
	wg.Done()
}

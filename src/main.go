/*
go run main.go -action=split -src="xxx.exe" -chunkSize=5000000
go run main.go -action=merge -src="./output/" -maxIdx=10
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	Action    string
	InputSrc  string
	ChunkSize int
	MaxIdx    int
)

const (
	Ext           = "lfs"       // 我們用此副檔名來區分是我們程式所要產出或處理的檔案
	DefaultOutVar = "./output/" // 預設的split所輸出的路徑
)

func init() {
	flag.StringVar(&Action, "action", "", "您想要合併還是拆分檔案(merge, split)")
	flag.StringVar(&InputSrc, "src", "", "如果是merge請輸入想要merge的資料夾, 如果是split請輸入檔案路徑")
	flag.IntVar(&ChunkSize, "chunkSize", 6*1024*1024, "[split] default 6MB: 6000000 指的是拆分檔案，每一份檔案的最大檔案大小")
	flag.IntVar(&MaxIdx, "maxIdx", 10, "[merge] merge時最大合併到第幾個檔案標號為止")
	flag.Parse()
	Action = strings.ToLower(Action) // case insensitive
	if Action != "merge" && Action != "split" {
		flag.Usage()
		os.Exit(2)
	}
}

func DoSplit() {
	targetFile := InputSrc
	f, err := os.Open(targetFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf := make([]byte, ChunkSize)

	// 建立輸出資料夾
	outputDir := DefaultOutVar
	if _, err = os.Stat(outputDir); os.IsNotExist(err) {
		err = os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 拆分文件
	for i := 0; ; i++ {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			return
		}
		// 創建一個文件，把當前讀到的內容寫入
		outF, err2 := os.Create(filepath.Join(outputDir, fmt.Sprintf("%d.%s", i, Ext)))
		if err2 != nil {
			log.Fatal(err2)
		}
		_, _ = outF.Write(buf[:n])
		_ = outF.Close()
	}
}

func DoMerge() {
	var err error
	srcDir := InputSrc
	// chunkSize := ChunkSize * MaxIdx // 合併後的總檔案大小上限
	// buf := make([]byte, chunkSize)
	outF, err2 := os.Create(fmt.Sprintf("./result.%s.merge", Ext))
	if err2 != nil {
		log.Fatal(err2)
	}
	defer outF.Close()
	for i := 0; i < MaxIdx; i++ {
		curFilePath := filepath.Join(srcDir, fmt.Sprintf("%d.%s", i, Ext))
		if _, err = os.Stat(curFilePath); os.IsNotExist(err) {
			log.Println("done!") // 理論上編號都是連續的，所以只要該編號的檔案找不到，合併應該就算完成了
			return
		}

		log.Printf("讀取文件: %q\n", curFilePath)
		bytesData, err := os.ReadFile(curFilePath) // 一次讀取該文件的所有資料
		if err != nil {
			log.Fatal(err)
		}
		_, err = outF.Write(bytesData)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("寫入完成: %q\n", curFilePath)
	}
}

func main() {
	actionMap := map[string]func(){
		"merge": DoMerge,
		"split": DoSplit,
	}
	actionMap[Action]()
}

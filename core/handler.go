package core

import (
	"fmt"
	"log"
	"lyric/tool"
)

func HandlerDir(path string) {
	// 遍历文件夹下文件 以及路径
	log.Printf("start searching music in: %s", path)
	lyricer := new(lyric)
	lyricer.fileList, lyricer.counts = lyricer.GetLyricList(path)

	_length := len(fileList)
	if _length == 0 {
		log.Printf("no music found in: %s", path)
		return
	}
	log.Printf("--------- %d file found--------", _length)
	for _, v := range tool.Exts {
		fmt.Printf("%s > %d \n", v, lyricer.counts[v])
	}
	log.Print("-------- start process --------")
	for _, file := range lyricer.fileList {
		lyricer.GetIds(file)
	}
}

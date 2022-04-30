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
	lyricer.InitData()

	lyricer.GetLyricList(path)
	log.Println(lyricer.List, lyricer.Count)

	_length := len(lyricer.List)
	if _length == 0 {
		log.Printf("no music found in: %s", path)
		return
	}
	log.Printf("--------- %d file found--------", _length)
	for _, v := range tool.Exts {
		fmt.Printf("%s > %d \n", v, lyricer.Count[v])
	}
	log.Print("-------- start process --------")
	for _, file := range lyricer.List {
		lyricer.GetIds(file)
	}
}

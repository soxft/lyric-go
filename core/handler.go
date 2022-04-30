package core

import (
	"fmt"
	"log"
	"lyric/tool"
)

func HandlerDir(path string) {
	// 遍历文件夹下文件 以及路径
	log.Printf("-- start searching in: %s --", path)
	lyricer := new(lyric)
	lyricer.InitData()

	lyricer.GetLyricList(path)
	_length := len(lyricer.List)
	if _length == 0 {
		log.Fatalf("no music found in: %s", path)
	}

	log.Printf("--------- %d file found--------", _length)
	for _, v := range tool.Exts {
		fmt.Printf("%s > %d \n", v, lyricer.Count[v])
	}
	log.Print("-------- start get music id --------")
	for _, file := range lyricer.List {
		lyricer.GetIds(file)
	}

}

package core

import (
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
		log.Printf("%s > %d \n", v, lyricer.Count[v])
	}
	log.Print("-------- start get music id --------")
	for _, file := range lyricer.List {
		lyricer.GetIds(file)
	}
	if len(lyricer.Ids) > 0 {
		log.Print("-------- start get lyrics --------")
		for _, id := range lyricer.Ids {
			lyricer.GetLyrics(id)
		}
	}
	log.Print("----------- status -----------")
	_failedNum := len(lyricer.Fail)
	_successNum := _length - _failedNum
	log.Printf("Total: %d", _length)
	log.Printf("Success: %d", _successNum)
	log.Printf("Failed: %d", _failedNum)
	if _failedNum > 0 {
		log.Print("--------- failed list ---------")
		for _, v := range lyricer.Fail {
			log.Printf("%s", v)
		}
	}
	log.Print("---------- complete ----------")
}

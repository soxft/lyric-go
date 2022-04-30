package core

import "log"

func (Lyricer lyric) GetLyric(file FileStruct) {
	log.Println(file.FileName)
}

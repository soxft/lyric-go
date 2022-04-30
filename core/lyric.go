package core

import "log"

func (Lyric *lyric) GetLyric(file FileStruct) {
	log.Println(file.FileName)
}

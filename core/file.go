package core

import (
	"io/ioutil"
	"log"
	"lyric/tool"
	"path/filepath"
)

func (Lyric *lyric) GetLyricList(path string) {
	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			if file.IsDir() {
				Lyric.GetLyricList(path + "/" + file.Name())
			} else {
				fileExt := filepath.Ext(file.Name())
				if tool.IsAudio(fileExt) {
					log.Println(file.Name())
					Lyric.Count[fileExt]++
					Lyric.List = append(Lyric.List, FileStruct{
						FileName: file.Name(),
						FilePath: path + "/" + file.Name(),
						FileExt:  fileExt,
					})
				}
			}
		}
	} else {
		log.Printf("read dir error: %v", err)
	}
}
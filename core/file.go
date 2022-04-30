package core

import (
	"io/ioutil"
	"log"
	"lyric/tool"
	"path/filepath"
)

var fileList []FileStruct
var fileCount = make(map[string]int)

func (Lyricer lyric) GetLyricList(path string) ([]FileStruct, map[string]int) {
	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			if file.IsDir() {
				Lyricer.GetLyricList(path + "/" + file.Name())
			} else {
				fileExt := filepath.Ext(file.Name())
				if tool.IsAudio(fileExt) {
					fileCount[fileExt]++
					fileList = append(fileList, FileStruct{
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
	return fileList, fileCount
}

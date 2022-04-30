package core

import (
	"encoding/json"
	"io"
	"log"
	"lyric/tool"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func (Lyric *lyric) GetIds(file FileStruct) {
	name := strings.Replace(file.Name, filepath.Ext(file.Name), "", -1)
	if tool.FileExists(file.Path + string(os.PathSeparator) + name + ".lrc") {
		log.Println(file.Name, " > already exists")
		return
	}
	postData := url.Values{}
	postData.Add("s", name)
	postData.Add("limit", "1")
	postData.Add("type", "1")
	postData.Add("offset", "0")
	data, err := http.PostForm("https://music.163.com/api/search/get/", postData)
	songId := getId(data, err)
	if songId == -1 {
		Lyric.Fail = append(Lyric.Fail, file.Name)
		log.Println(file.Name, " > failed")
		return
	}
	Lyric.Ids = append(Lyric.Ids, IdStruct{
		Name: name,
		Id:   songId,
		Path: file.Path,
	})
	log.Println(file.Name, " > ", songId)
}

func getId(response *http.Response, err error) int {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return -1
	}
	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return -1
	}
	if int(result["code"].(float64)) != 200 {
		return -1
	}
	log.Println(result)
	if result["result"].(map[string]interface{})["songs"] == nil {
		return -1
	}
	songId := result["result"].(map[string]interface{})["songs"].([]interface{})[0].(map[string]interface{})["id"].(float64)
	return int(songId)
}

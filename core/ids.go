package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func (Lyric *lyric) GetIds(file FileStruct) {
	name := strings.Replace(file.Name, filepath.Ext(file.Name), "", -1)
	postData := url.Values{}
	postData.Add("s", name)
	postData.Add("limit", "1")
	postData.Add("type", "1")
	postData.Add("offset", "0")
	data, err := http.PostForm("https://music.163.com/api/search/get/", postData)
	if err != nil {
		log.Println(err)
		return
	}
	songId := getId(data)
	if songId == -1 {
		Lyric.Fail = append(Lyric.Fail, file.Name)
	}
	Lyric.Ids[file.Name] = songId
	log.Println(file.Name, " > ", songId)
}

func getId(response *http.Response) int {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return -1
	}
	if int(result["code"].(float64)) != 200 {
		return -1
	}
	songId := result["result"].(map[string]interface{})["songs"].([]interface{})[0].(map[string]interface{})["id"].(float64)
	return int(songId)
}

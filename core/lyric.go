package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func (Lyric *lyric) GetLyrics(file IdStruct) {
	getData := url.Values{}
	getData.Set("os", "pc")
	getData.Set("id", strconv.Itoa(file.Id))
	getData.Set("lv", "-1")
	getData.Set("kv", "-1")
	getData.Set("tv", "-1")
	res, err := http.Get("https://music.163.com/api/song/lyric" + "?" + getData.Encode())
	lyrics := getLyric(res, err)
	if lyrics == "" {
		Lyric.Fail = append(Lyric.Fail, file.Name)
		log.Println(file.Name, " > failed")
		return
	}
	if !Lyric.WriteLyric(file.Path+"/"+file.Name+".lrc", lyrics) {
		Lyric.Fail = append(Lyric.Fail, file.Name)
		log.Println(file.Name, " > write file failed")
		return
	}
	log.Println(file.Name, " > success")
}

func getLyric(response *http.Response, err error) string {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return ""
	}
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return ""
	}
	if result["lrc"] == nil {
		return ""
	}
	lyric := result["lrc"].(map[string]interface{})
	if lyric["lyric"] == nil {
		return ""
	}
	return lyric["lyric"].(string)
}

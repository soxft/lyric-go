package core

type FileStruct struct {
	Name string
	Path string
	Ext  string
}

type Lyricer interface {
	GetLyricList(string) []string
	GetIds(file FileStruct)
	GetLyric(file FileStruct)
	InitData()
}

type lyric struct {
	List  []FileStruct
	Count map[string]int
	Ids   map[string]int
	Fail  []string
}

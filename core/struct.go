package core

type FileStruct struct {
	Name string
	Path string
	Ext  string
}

type IdStruct struct {
	Name string
	Path string
	Id   int
}

type Lyricer interface {
	GetLyricList(string) []string
	GetIds(file FileStruct)
	GetLyric(file FileStruct)
	WriteLyric(file FileStruct, lyric string) bool
	InitData()
}

type lyric struct {
	List  []FileStruct
	Count map[string]int
	Ids   []IdStruct
	Fail  []string
}

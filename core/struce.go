package core

type FileStruct struct {
	FileName string
	FilePath string
	FileExt  string
}

type Lyricer interface {
	GetLyricList(string) []string
	GetIds(file FileStruct)
	GetLyric(file FileStruct)
}

type lyric struct {
	fileList []FileStruct
	counts   map[string]int
	ids      map[string]int
}

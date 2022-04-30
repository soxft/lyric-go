package core

func (Lyric *lyric) InitData() {
	Lyric.List = make([]FileStruct, 0)
	Lyric.Count = make(map[string]int)
	Lyric.Ids = make([]IdStruct, 0)
	Lyric.Fail = make([]string, 0)
}

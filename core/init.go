package core

func (Lyric *lyric) InitData() {
	Lyric.List = make([]FileStruct, 0)
	Lyric.Count = make(map[string]int)
}

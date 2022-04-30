package tool

var (
	Exts = []string{
		".flac",
		".mp3",
		".m4a",
		".ncm",
	}
)

func IsAudio(ext string) bool {
	for _, v := range Exts {
		if v == ext {
			return true
		}
	}
	return false
}

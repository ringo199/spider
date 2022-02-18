package constant

const TmpBasePath string = "tmpPath/"
const InputBasePath string = "inputPath/"
const OutputBasePath string = "outputPath/"

const AudioBasePath string = "audios/"
const VideoBasePath string = "videos/"
const DanmakuBasePath string = "danmakus/"
const SubtitleBasePath string = "subtitles/"

const BilibiliPath string = "bilibili/"
const DouyinPath string = "douyin/"

var FileTypePaths = []string{
	// SubtitleBasePath,
	// AudioBasePath,
	// VideoBasePath,
	DanmakuBasePath,
}

var PlatformTypePaths = []string{
	BilibiliPath,
	// DouyinPath,
}

var ASoulPaths = []string{
	"asoul",

	"ava",
	"bella",
	"carol",
	"diana",
	"eileen",
}

const DownloadLimit int = 5

const AudiosUrl string = "https://asoul-rec.herokuapp.com/ASOUL-REC/AAC%E5%BD%95%E6%92%AD%E9%9F%B3%E8%BD%A8/"
const SubtitleUrl string = "https://asoul-rec.herokuapp.com/ASOUL-REC/SRT%E8%AF%AD%E9%9F%B3%E8%BD%AC%E5%AD%97%E5%B9%95%E6%96%87%E4%BB%B6/"

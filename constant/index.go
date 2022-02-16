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
	AudioBasePath,
	SubtitleBasePath,
}

var PlatformTypePaths = []string{
	BilibiliPath,
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

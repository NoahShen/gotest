package speech2text

import (
	"testing"
)

func NoTestMp3ToFlac(t *testing.T) {
	url, e := Mp3ToFlac("https://imo.im/fd/E/uxy2s3jjrv/voiceim.mp3")
	if e != nil {
		t.Log(e)
	}
	t.Log(url)
}

func TestSpeech2Text(t *testing.T) {
	text, c, e := Speech2Text("/home/noah/output.flac")
	if e != nil {
		t.Log(e)
	}
	t.Log(text)
	t.Log(c)
}

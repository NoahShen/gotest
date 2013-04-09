package speech2text

import (
	"path/filepath"
	"testing"
)

func TestSpeech2Text(t *testing.T) {
	path, err := Mp3ToFlac("https://imo.im/fd/E/864hokp84r/voiceim.mp3")
	if err != nil {
		t.Fatal(err)
	}
	text, c, e := Speech2Text(path)
	if e != nil {
		t.Fatal(e)
	}
	t.Log(text)
	t.Log(c)
}

func NoTestDownloadFile(t *testing.T) {
	f, err := downloadFile("https://imo.im/fd/E/864hokp84r/voiceim.mp3")
	if err != nil {
		t.Fatal(err)
	}
	path, _ := filepath.Abs(f.Name())
	t.Log(path)
}

func NoTestMp3ToFlac(t *testing.T) {
	path, err := Mp3ToFlac("https://imo.im/fd/E/864hokp84r/voiceim.mp3")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
}

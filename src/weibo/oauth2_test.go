package weibo

import (
	"testing"
)

func TestSpeech2Text(t *testing.T) {
	text, c, e := Speech2Text("https://imo.im/fd/E/864hokp84r/voiceim.mp3")
	if e != nil {
		t.Fatal(e)
	}
	t.Log(text)
	t.Log(c)
}

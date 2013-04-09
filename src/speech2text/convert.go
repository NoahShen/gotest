package speech2text

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	SPEECH_URL = "http://www.google.com/speech-api/v1/recognize?xjerr=1&client=chromium&lang=zh-CN"
)

type speechRespJson struct {
	Id         string       `json:"id,omitempty"`
	Status     int          `json:"status,omitempty"`
	Hypotheses []hypotheses `json:"hypotheses,omitempty"`
}

type hypotheses struct {
	Utterance  string  `json:"utterance,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
}

func Speech2Text(voiceFile string) (string, float64, error) {
	bodyBuf := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuf)
	defer bodyWriter.Close()
	fileWriter, _ := bodyWriter.CreateFormFile("uploadfile", voiceFile)

	fh, openErr := os.Open(voiceFile)
	if openErr != nil {
		return "", 0, openErr
	}
	io.Copy(fileWriter, fh)
	speechReq, _ := http.NewRequest("POST", SPEECH_URL, bodyBuf)
	speechReq.Header.Set("Content-Type", "audio/x-flac; rate=16000")
	speechResp, postErr := http.DefaultClient.Do(speechReq)
	if postErr != nil {
		return "", 0, postErr
	}
	defer speechResp.Body.Close()
	bytes, readErr := ioutil.ReadAll(speechResp.Body)
	if readErr != nil {
		return "", 0, readErr
	}
	log.Println(string(bytes))
	speechResult := &speechRespJson{}
	if unmarshalErr := json.Unmarshal(bytes, speechResult); unmarshalErr != nil {
		return "", 0, unmarshalErr
	}
	if speechResult.Status != 0 {
		return "", 0, errors.New("speech2text errror!")
	}
	h := speechResult.Hypotheses[0]
	return h.Utterance, h.Confidence, nil
}

func Mp3ToFlac(mp3Url string) (string, error) {
	mp3File, downloadFileErr := downloadFile(mp3Url)
	if downloadFileErr != nil {
		return "", downloadFileErr
	}
	flacPath, _ := filepath.Abs("voiceim.flac")
	mp3Path, _ := filepath.Abs(mp3File.Name())
	output, convertErr := exec.Command("ffmpeg", "-i", mp3Path, flacPath).Output()
	if convertErr != nil {
		return "", convertErr
	}
	log.Println("convert output: ", string(output))
	os.Remove(mp3Path)
	return flacPath, nil
}

func downloadFile(fileUrl string) (*os.File, error) {
	mp3File := "voiceim.mp3"
	f, createFileErr := os.Create(mp3File)
	defer f.Close()
	if createFileErr != nil {
		return nil, createFileErr
	}

	resp, downloadFileErr := http.Get(fileUrl)
	defer resp.Body.Close()
	if downloadFileErr != nil {
		return nil, downloadFileErr
	}
	_, writeFileErr := io.Copy(f, resp.Body)
	if writeFileErr != nil {
		return nil, writeFileErr
	}
	return f, nil
}

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
	return "", nil

	//requestXml := fmt.Sprintf(InsertQueueUrl, API_KEY, mp3Url)

	//buf := new(bytes.Buffer)
	//w := multipart.NewWriter(buf)
	//w.WriteField("queue", requestXml)
	//convertReq, _ := http.NewRequest("POST", CONVERT_URL, buf)

	////convertReq, _ := http.NewRequest("POST", CONVERT_URL, nil)
	////convertReq.Form = make(url.Values)
	////convertReq.Form.Set("queue", requestXml)

	//fmt.Println(convertReq)
	//convertResp, convertReqErr := http.DefaultClient.Do(convertReq)
	//defer convertResp.Body.Close()
	//if convertReqErr != nil {
	//	return "", convertReqErr
	//}
	//convertRespBytes, readErr := ioutil.ReadAll(convertResp.Body)
	//if readErr != nil {
	//	return "", readErr
	//}
	//return string(convertRespBytes), nil
	//keyId := utils.RandomString(14)
	//buf := new(bytes.Buffer)
	//w := multipart.NewWriter(buf)
	//w.WriteField("APC_UPLOAD_PROGRESS", keyId)
	//w.WriteField("storedOpt", "70")
	//w.WriteField("FileOrURLFlag", "url")
	//w.WriteField("download_url", fileUrl)
	//w.WriteField("youtube_mode", "default")
	//w.WriteField("input_format", ".mp3")
	//w.WriteField("output_format", ".flac")

	//convertReq, _ := http.NewRequest("POST", CONVERT_URL, buf)
	//convertReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.172 Safari/537.22")
	//convertReq.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//convertReq.Header.Set("Accept-Language", "en-US,en;q=0.5")
	//convertReq.Header.Set("Accept-Encoding", "gzip, deflate")

	//log.Println("convertReq:", convertReq)

	//log.Println("Send convert request!")
	//convertResp, convertReqErr := http.DefaultClient.Do(convertReq)
	//defer convertResp.Body.Close()
	//if convertReqErr != nil {
	//	return "", convertReqErr
	//}
	//log.Println("Convert Response Status:", convertResp.Status)

	//if convertResp.StatusCode != 200 {
	//	return "", errors.New(convertResp.Status)
	//}
	//convertRespBytes, convertRespErr := ioutil.ReadAll(convertResp.Body)
	//if convertRespErr != nil {
	//	return "", convertRespErr
	//}
	//log.Println(string(convertRespBytes))

	//for {
	//	progressReq, _ := http.NewRequest("GET", PROGRESS_URL+keyId, nil)
	//	progressReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.172 Safari/537.22")
	//	progressResp, progReqErr := http.DefaultClient.Do(progressReq)
	//	defer progressResp.Body.Close()
	//	if progReqErr != nil {
	//		return "", progReqErr
	//	}
	//	bytes, readErr := ioutil.ReadAll(progressResp.Body)
	//	if readErr != nil {
	//		return "", readErr
	//	}
	//	progress := string(bytes)
	//	//log.Println("Convert progress:", progress)
	//	if progress == "100" {
	//		break
	//	}
	//	time.Sleep(5000)
	//}

	//getFileReq, _ := http.NewRequest("GET", GET_FILE_URL+keyId, nil)
	//getFileResp, getFileReqErr := http.DefaultClient.Do(getFileReq)
	//defer getFileResp.Body.Close()
	//if getFileReqErr != nil {
	//	return "", getFileReqErr
	//}
	//bytes, readErr := ioutil.ReadAll(getFileResp.Body)
	//if readErr != nil {
	//	return "", readErr
	//}
	//html := string(bytes)
	//log.Println("File download_html:", html)
	//index := strings.Index(html, "http://dw4.convertfiles.com")
	//if index == -1 {
	//	return "", errors.New("Can not find downlaod file url")
	//}
	//endindex := strings.Index(html, ".flac")
	//if endindex == -1 {
	//	return "", errors.New("Can not find downlaod file url")
	//}
	//flacFileUrl := html[index : endindex+5]
	//return flacFileUrl, nil
}

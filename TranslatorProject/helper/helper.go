package helper

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/AllenOps08/TranslatorBot/constant"
	"github.com/AllenOps08/TranslatorBot/model"
)

// Common functions which can be used

func Translate(obj model.Translate) ([]byte, error) {
	//Building query parameter, converting query to string
	str2 := "q=" + obj.Text + "&source=" + obj.SourceLanguage + "&target=" + obj.TargetLanguage

	payLoad := strings.NewReader(str2)

	req, err := http.NewRequest(http.MethodPost, constant.TranslateURL, payLoad)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("GoogleApiKey"))
	req.Header.Add("X-RapidAPI-Host", "google-translate113.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

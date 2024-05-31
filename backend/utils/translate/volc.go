package translate

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Volc struct{}

var (
	kAccessKey      = os.Getenv("VOLC_AK") // https://console.volcengine.com/iam/keymanage/
	kSecretKey      = os.Getenv("VOLC_SK")
	kServiceVersion = "2020-06-01"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "translate.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "translate"},
	}
	ApiInfoList = map[string]*base.ApiInfo{
		"TranslateText": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TranslateText"},
				"Version": []string{kServiceVersion},
			},
		},
	}
)
var client *base.Client

func init() {
	fmt.Printf("=========\n %v \n %v \n=================\n", kAccessKey, kSecretKey)
	client = base.NewClient(ServiceInfo, ApiInfoList)
	client.SetAccessKey(kAccessKey)
	client.SetSecretKey(kSecretKey)
}

type Req struct {
	SourceLanguage string   `json:"SourceLanguage"`
	TargetLanguage string   `json:"TargetLanguage"`
	TextList       []string `json:"TextList"`
}
type Res struct {
	TranslationList []struct {
		Translation            string      `json:"Translation"`
		DetectedSourceLanguage string      `json:"DetectedSourceLanguage"`
		Extra                  interface{} `json:"Extra"`
	} `json:"TranslationList"`
}

var LangMap = map[string]string{
	"EN": "en",
	"ZH": "zh",
	"JP": "ja",
}

func (v *Volc) Translate(text string, sourceLang string, targetLang string) (string, error) {
	// 实现 Volc 的翻译逻辑
	req := Req{
		SourceLanguage: LangMap[sourceLang],
		TargetLanguage: LangMap[targetLang],
		TextList:       []string{text},
	}
	body, _ := json.Marshal(req)
	resp, _, err := client.Json("TranslateText", nil, string(body))
	fmt.Println("resp:", string(resp))
	respStruct := Res{}
	err = json.Unmarshal(resp, &respStruct)
	if err != nil {
		return "", err
	}
	fmt.Println("respS:", respStruct)
	return respStruct.TranslationList[0].Translation, nil
}

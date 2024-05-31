package translate

import (
	"fmt"
	"github.com/imroc/req/v3"
)

type DeepL struct{}

type DeepLRequest struct {
	Text       string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}
type DeepLResponse struct {
	Alternatives []string `json:"alternatives"`
	Code         int      `json:"code"`
	Data         string   `json:"data"`
	ID           int64    `json:"id"`
	Method       string   `json:"method"`
	SourceLang   string   `json:"source_lang"`
	TargetLang   string   `json:"target_lang"`
}

func (d *DeepL) Translate(text string, sourceLang string, targetLang string) (string, error) {
	// 实现 DeepL 的翻译逻辑
	//url := "http://localhost:1188/translate"
	url := "http://deeplx:1188/translate"
	requestBody := DeepLRequest{
		Text:       text,
		SourceLang: sourceLang,
		TargetLang: targetLang,
	}
	// 发送请求
	client := req.C().DevMode()
	response, err := client.R().SetBody(requestBody).Post(url)
	if err != nil {
		return "", err
	}
	// 解析响应
	var deepLResponse DeepLResponse
	err = response.UnmarshalJson(&deepLResponse)
	if err != nil {
		return "", err
	}
	// 返回翻译结果
	fmt.Println(deepLResponse)
	return deepLResponse.Data, nil
}

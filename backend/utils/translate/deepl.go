package translate

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/imroc/req/v3"
	"go.uber.org/zap"
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
	url := "http://deeplx:1188/translate"
	requestBody := DeepLRequest{
		Text:       text,
		SourceLang: sourceLang,
		TargetLang: targetLang,
	}
	// 发送请求
	client := req.C()
	response, err := client.R().SetBody(requestBody).Post(url)
	if err != nil {
		global.GVA_LOG.Error("DeepL 翻译失败, 请求出错", zap.Error(err))
		return "", err
	}
	// 解析响应
	var deepLResponse DeepLResponse
	err = response.UnmarshalJson(&deepLResponse)
	if err != nil {
		global.GVA_LOG.Error("DeepL 翻译失败, JSON 解析出错", zap.Error(err))
		return "", err
	}
	// 返回翻译结果
	global.GVA_LOG.Debug("DeepL 翻译成功",
		zap.String("text", text), zap.String("sourceLang", sourceLang),
		zap.String("targetLang", targetLang), zap.String("data", deepLResponse.Data))
	return deepLResponse.Data, nil
}

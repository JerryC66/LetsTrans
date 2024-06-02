package document

import (
	"bytes"
	"encoding/json"
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"github.com/firwoodlin/letstrans/utils"
	"github.com/imroc/req/v3"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type ProcessResponse struct {
	Segments []letstrans.Segment `json:"segments"`
}

// ProcessDocument uploads a document to the processing endpoint and returns the segments.
func ProcessDocument(doc letstrans.Document) ([]letstrans.Segment, error) {
	filepath := doc.FilePath
	url := "http://doc-processor:5000/process_document"
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//var segments []letstrans.Segment
	var res ProcessResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res.Segments, nil
}

// ModifyDocument uploads the document along with the replacements and retrieves the modified document.
func ModifyDocument(filePath string, replacements []letstrans.Segment) (outFilePath string, err error) {
	url := "http://doc-processor:5000/modify_document"
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	replacementsJSON, err := json.Marshal(replacements)
	if err != nil {
		return
	}

	// 发送请求
	client := req.C()
	resp, err := client.R().
		SetFile("file", filePath).
		SetFormData(map[string]string{
			"replacements": string(replacementsJSON),
		}).
		Post(url)
	if err != nil {
		global.GVA_LOG.Error("ModifyDocument failed", zap.Error(err))
	}
	outFileName := "modified_" + filepath.Base(filePath)
	outFileDir := filepath.Join(filepath.Dir(filePath), "modified")
	if err = utils.CreateDir(outFileDir); err != nil {
		global.GVA_LOG.Error("Failed to create directory", zap.Error(err))
		return "", err
	}

	outFilePath = filepath.Join(outFileDir, outFileName)
	outFile, err := os.Create(outFilePath)
	if err != nil {
		global.GVA_LOG.Error("Failed to create output file", zap.Error(err))
		return "", err
	}
	defer outFile.Close()

	if _, err := outFile.Write(resp.Bytes()); err != nil {
		global.GVA_LOG.Error("Failed to write modified document", zap.Error(err))
		return "", err
	}
	return outFilePath, err
}

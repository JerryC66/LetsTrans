package document

import (
	"bytes"
	"encoding/json"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"io"
	"mime/multipart"
	"net/http"
	"os"
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
func ModifyDocument(filepath string, replacements []letstrans.Segment) error {
	url := "http://doc-processor:5000/modify_document"

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	replacementsJSON, err := json.Marshal(replacements)
	if err != nil {
		return err
	}

	_ = writer.WriteField("replacements", string(replacementsJSON))
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	outFile, err := os.Create("modified_" + filepath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	return err
}

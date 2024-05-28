package document // Doc2Seg Document to Segments
import (
	"bufio"
	"fmt"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"os"
)

type TxtProcessor struct{}

func (tdp TxtProcessor) Doc2Seg(doc letstrans.Document) ([]letstrans.Segment, error) {
	// 从文档中提取段落
	segments := []letstrans.Segment{}
	filePath := doc.FilePath
	// 1. 读取文档
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("file close failed")
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		segments = append(segments,
			letstrans.Segment{
				SourceText: scanner.Text(),
				DocumentID: doc.ID,
			})
	}
	return segments, nil
}

// Seg2Doc converts segments to a text document.
func (tdp TxtProcessor) Seg2Doc(segments []letstrans.Segment, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, segment := range segments {
		_, err := writer.WriteString(segment.SourceText + "\n")
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

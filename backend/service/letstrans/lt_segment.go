package letstrans

import (
	"bufio"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"os"
)

type SegmentService struct{}

// Doc2Seg Document to Segments
func Doc2Seg(doc letstrans.Document) ([]letstrans.Segment, error) {
	// 从文档中提取段落
	segments := []letstrans.Segment{}
	filePath := doc.FilePath
	// 1. 读取文档
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		segments = append(segments,
			letstrans.Segment{
				SourceText: scanner.Text(),
				DocumentID: doc.ID,
			})
	}
	//global.GVA_DB.Model(&letstrans.Segment{}).Save(&segments)
	//return lines, scanner.Err()
	// 2. 分割文档
	// 3. 生成段落
	// 4. 返回段落
	return segments, nil
}

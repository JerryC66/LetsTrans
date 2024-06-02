package utils

import (
	"encoding/csv"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"mime/multipart"
)

func CSV2Terms(file *multipart.FileHeader, glossaryID uint) (terms []letstrans.Term, err error) {
	// 打开文件
	f, err := file.Open()
	if err != nil {
		return
	}
	defer f.Close()

	// 创建 CSV 读取器
	reader := csv.NewReader(f)
	reader.Comma = ',' // 设定分隔符
	reader.FieldsPerRecord = -1

	// 读取所有行
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	srcLang := records[0][0]
	tgtLang := records[0][1]
	// 假设表头是 EN 和 ZH
	for _, record := range records[1:] {

		term := letstrans.Term{
			SourceLang: srcLang,
			TargetLang: tgtLang,
			SourceText: record[0],
			TargetText: record[1],
			GlossaryID: glossaryID,
		}

		terms = append(terms, term)
	}
	return terms, nil
}

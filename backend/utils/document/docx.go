package document

import "github.com/firwoodlin/letstrans/model/letstrans"

type DocxProcessor struct{}

func (DocxProcessor) Doc2Seg(doc letstrans.Document) ([]letstrans.Segment, error) {
	//TODO implement me
	panic("implement me")
}

func (DocxProcessor) Seg2Doc(segments []letstrans.Segment, filePath string) error {
	//TODO implement me
	panic("implement me")
}

package document

import "github.com/firwoodlin/letstrans/model/letstrans"

type PptxProcessor struct{}

func (PptxProcessor) Doc2Seg(doc letstrans.Document) ([]letstrans.Segment, error) {
	//TODO implement me
	panic("implement me")
}

func (PptxProcessor) Seg2Doc(segments []letstrans.Segment, filePath string) error {
	//TODO implement me
	panic("implement me")
}

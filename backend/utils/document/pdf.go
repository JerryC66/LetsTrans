package document

import "github.com/firwoodlin/letstrans/model/letstrans"

type PdfProcessor struct{}

func (PdfProcessor) Doc2Seg(doc letstrans.Document) ([]letstrans.Segment, error) {
	//TODO implement me
	panic("implement me")
}

func (PdfProcessor) Seg2Doc(segments []letstrans.Segment, filePath string) error {
	//TODO implement me
	panic("implement me")
}

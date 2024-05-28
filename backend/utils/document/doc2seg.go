package document

import (
	"github.com/firwoodlin/letstrans/model/letstrans"
)

type DocProcessor interface {
	Doc2Seg(doc letstrans.Document) ([]letstrans.Segment, error)
	Seg2Doc(segments []letstrans.Segment, filePath string) error
}

func NewDocProcessor(fileType string) DocProcessor {
	switch fileType {
	case "txt":
		return &TxtProcessor{}
	default:
		return &TxtProcessor{}
	}
}

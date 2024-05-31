package letstrans

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
)

type SegmentService struct{}

func (s *SegmentService) GetSegmentList(docID uint) (segments []letstrans.Segment, err error) {
	err = global.GVA_DB.Model(&letstrans.Segment{}).Where("document_id = ?", docID).Order("id ASC").Find(&segments).Error
	return segments, err
}

func (s *SegmentService) UpdateSegment(segID uint, seg letstrans.Segment) (err error) {
	// 注意此处对于空字段的处理
	vals := map[string]interface{}{}
	if seg.Finished != nil {
		vals["finished"] = seg.Finished
	}
	if seg.TargetText != "" {
		vals["target_text"] = seg.TargetText
	}

	err = global.GVA_DB.Model(&letstrans.Segment{}).Where("id= ?", segID).Updates(&vals).Error
	return err
}

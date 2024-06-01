package letstrans

import (
	"github.com/agnivade/levenshtein"
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"go.uber.org/zap"
	"sort"
)

// TranslateMemoryService 定义翻译记忆库服务结构体
type TranslateMemoryService struct{}

// SimilarSegment 表示一个包含相似度的翻译段
type SimilarSegment struct {
	Segment    letstrans.Segment
	Similarity int
}

func (ts *TranslateMemoryService) GetSuggestions(sourceText string) (suggestions []letstrans.TranslationMemory, err error) {
	segments := []letstrans.Segment{}
	//err := global.GVA_DB.Joins("JOIN documents ON documents.id = segments.document_id").
	//	Where("documents.source_lang = ? AND documents.target_lang = ? documents.author_id = ? AND segments.finished = true",
	//		sourceLang, targetLang, authorID).
	//	Find(&segments).Error
	err = global.GVA_DB.Model(&letstrans.Segment{}).Where(" segments.finished = true").Find(&segments).Error
	if err != nil {
		return
	}
	var similarSegments []SimilarSegment
	for _, segment := range segments {
		distance := levenshtein.ComputeDistance(sourceText, segment.SourceText)
		similarSegments = append(similarSegments, SimilarSegment{
			Segment:    segment,
			Similarity: distance,
		})
	}

	// 按相似度排序
	sort.Slice(similarSegments, func(i, j int) bool {
		return similarSegments[i].Similarity < similarSegments[j].Similarity
	})

	// 获取相似度最高的三条记录
	for i := 0; i < len(similarSegments) && i < 3; i++ {
		simSeg := similarSegments[i]
		suggestions = append(suggestions, letstrans.TranslationMemory{
			DocumentID: simSeg.Segment.DocumentID,
			SourceText: simSeg.Segment.SourceText,
			TargetText: simSeg.Segment.TargetText,
			SimRank:    int64(i + 1),
		})
		global.GVA_LOG.Info("Segment similarity",
			zap.String("source_text", simSeg.Segment.SourceText),
			zap.Int("similarity", simSeg.Similarity))
	}
	return
}

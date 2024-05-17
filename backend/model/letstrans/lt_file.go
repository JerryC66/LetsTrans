package letstrans

// FileRecord 定义文件记录结构体
type FileRecord struct {
	BaseModel
	FileName string `gorm:"type:varchar(255)"`
	FilePath string `gorm:"type:varchar(255)"`
}

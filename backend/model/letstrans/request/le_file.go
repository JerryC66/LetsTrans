package request

type FileDownloadReq struct {
	IDs  []uint `json:"ids" binding:"required"`
	Type string `json:"type" binding:"required"`
}

package model

type ChatFile struct {
	Model
	FileUrl string `json:"file_url" gorm:"comment:文件URL"`
}

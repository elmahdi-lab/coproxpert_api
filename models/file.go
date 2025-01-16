package models

import (
	"regexp"

	"github.com/google/uuid"
)

type Provider string

const DocumentExtensionsRegex = "application/(pdf|doc|docx|xls|xlsx)"
const ImageExtensionsRegex = "image/(jpeg|png)"

const (
	Aws Provider = "aws"
	Gcp Provider = "gcp"
)

var AllowedFileTypes = []string{
	"application/pdf",
	"image/jpeg",
	"image/png",
	"application/doc",
	"application/docx",
	"application/xls",
	"application/xlsx",
}

type File struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User       User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	BucketName string    `json:"bucket_name"`
	Provider   Provider  `json:"provider" gorm:"not null"`
	FileType   string    `json:"file_type" gorm:"not null"`
	PublicUrl  string    `json:"public_url"`
	PrivateUrl string    `json:"private_url"`

	BaseModel
}

func (f *File) IsImage() bool {
	imageRegex := regexp.MustCompile(ImageExtensionsRegex)
	return imageRegex.MatchString(f.FileType)
}

func (f *File) IsDocument() bool {
	docRegex := regexp.MustCompile(DocumentExtensionsRegex)
	return docRegex.MatchString(f.FileType)
}

package files

import (
	"fmt"
	"io"
)

type CloudProvider string

const (
	AWS CloudProvider = "aws"
	GCP CloudProvider = "gcp"
)

type FileManager interface {
	Upload(file io.Reader, filename string) error
	Link(filename string) (string, error)
	Delete(filename string) error
}

func NewFileManager(providerType CloudProvider) (FileManager, error) {
	switch providerType {
	case AWS:
		return &AWSFileManager{}, nil
	case GCP:
		return &GCPFileManager{}, nil
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", providerType)
	}
}

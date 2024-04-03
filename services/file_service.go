package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers/gcp"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
	"ithumans.com/coproxpert/types"
	"mime/multipart"
)

func isFileTypeAllowed(fileType string) bool {
	for _, allowedFileType := range models.AllowedFileTypes {
		if fileType == allowedFileType {
			return true
		}
	}
	return false
}

func CreateFile(h *multipart.FileHeader) (*models.File, types.ErrorResponse) {
	// Upload the file using FileManager
	fileManager, err := gcp.NewFileManager()
	errResponse := types.ErrorResponse{
		Code: "file_not_uploaded",
	}

	if err != nil {
		errResponse.Message = "Failed to upload the file"
		return nil, errResponse
	}

	contentType := h.Header.Get("Content-Type")

	// Check if the file type is allowed
	if !isFileTypeAllowed(contentType) {
		errResponse.Message = "File type not allowed"
		return nil, errResponse
	}

	// Open the file
	fileContent, err := h.Open()
	if err != nil {
		errResponse.Message = "Failed to open the uploaded file"
		return nil, errResponse
	}
	defer func(fileContent multipart.File) {
		err := fileContent.Close()
		if err != nil {
			return
		}
	}(fileContent)

	// Upload the file
	err = fileManager.Upload(fileContent, h.Filename)
	if err != nil {
		errResponse.Message = "Failed to upload the file"
		return nil, errResponse
	}

	// Create a new file record
	fileRecord := models.File{
		BucketName: h.Filename,
		Provider:   models.Gcp,
		FileType:   contentType,
		PublicUrl:  "", // You might want to set these values appropriately if available
		PrivateUrl: "", // You might want to set these values appropriately if available
	}

	// Save the file record
	fileRepository := repositories.NewFileRepository()
	if err != nil {
		errResponse.Message = "Failed to get file repository"
		return nil, errResponse
	}

	err = fileRepository.Create(&fileRecord)
	if err != nil {
		errResponse.Message = "Failed to save the uploaded file record"
		return nil, errResponse
	}

	return &fileRecord, types.ErrorResponse{}
}

func GetFileByID(id uuid.UUID) (*models.File, error) {
	fileRepository := repositories.NewFileRepository()
	if fileRepository == nil {
		return nil, nil
	}

	file, err := fileRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func UpdateFile(f *models.File) (*models.File, error) {
	fileRepository := repositories.NewFileRepository()
	if fileRepository == nil {
		return nil, nil
	}

	err := fileRepository.Update(f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func DeleteFileByID(id uuid.UUID) bool {
	fileRepository := repositories.NewFileRepository()
	if fileRepository == nil {
		return false
	}

	return fileRepository.DeleteByID(id)
}

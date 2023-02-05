package upload

import (
	"mime/multipart"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader, pathStr string, nameStr string) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法 后期可以扩展上传第三方
func NewOss() OSS {
	return &Local{}
}

// UploadFile 上传文件
func UploadFile(header *multipart.FileHeader, pathStr string, nameStr string) (string, string, error) {
	oss := NewOss()
	filePath, key, uploadErr := oss.UploadFile(header, pathStr, nameStr)
	if uploadErr != nil {
		return "", "", uploadErr
	}
	return filePath, key, nil
}

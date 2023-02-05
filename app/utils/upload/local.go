package upload

import (
	"dcr-gin/app/global"
	"dcr-gin/app/utils"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct{}

//UploadFile
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*Local) UploadFile(file *multipart.FileHeader, pathStr string, nameStr string) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := nameStr
	if nameStr == "" {
		name = strings.TrimSuffix(file.Filename, ext)
		name = utils.Md5x(name)
	}
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	t := time.Now()
	tmpPath := fmt.Sprintf("/%v/%v/%v/%v", t.Year(), int(t.Month()), t.Day(), pathStr)
	mkdirErr := os.MkdirAll(global.ServerConfig.Local.StorePath+tmpPath, os.ModePerm)
	if mkdirErr != nil {
		global.Logger.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.ServerConfig.Local.StorePath + tmpPath + filename
	filepath := global.ServerConfig.Local.Path + tmpPath + filename
	fmt.Println(p)
	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.Logger.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	out, createErr := os.Create(p)
	if createErr != nil {
		global.Logger.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close()             // 创建文件 defer 关闭
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.Logger.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

//DeleteFile
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error
func (*Local) DeleteFile(key string) error {
	p := global.ServerConfig.Local.StorePath + "/" + key
	if strings.Contains(p, global.ServerConfig.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}

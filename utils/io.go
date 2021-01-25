package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type IOUtils struct {
}

var ioInstance *IOUtils
var ioOnce sync.Once

func IO() *IOUtils {
	ioOnce.Do(func() {
		ioInstance = &IOUtils{}
	})
	return ioInstance
}

// 判断指定路径文件是否存在
func (i *IOUtils) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 复制文件到目标路径
func (i *IOUtils) CopyFile(src string, dst string) error {
	srcFile, errSrc := os.Open(src)
	if errSrc != nil {
		return errSrc
	}
	return i.CopyFileFromReader(srcFile, dst)
}

// 通过Reader复制文件
func (i *IOUtils) CopyFileFromReader(srcFile io.Reader, dst string) error {
	dirPath, _ := path.Split(dst)
	if !i.PathExists(dirPath) {
		dstDirErr := os.MkdirAll(dirPath, os.ModePerm)
		if dstDirErr != nil {
			return dstDirErr
		}
	}
	dstFile, errDstCreate := os.Create(dst)
	if errDstCreate != nil {
		return errDstCreate
	}
	defer dstFile.Close()
	_, errDestCopy := io.Copy(dstFile, srcFile)
	if errDestCopy != nil {
		return errDestCopy
	}
	return nil
}

// 复制整个路径到目标位置
func (i *IOUtils) CopyDir(src string, dest string) error {
	err := filepath.Walk(src, func(currentSrc string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		aimPath := strings.Replace(currentSrc, src, dest, 1)
		if !f.IsDir() {
			i.CopyFile(currentSrc, aimPath)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// 将本地的JSON文件读取并转换成指定的结构体
func (i *IOUtils) JsonFileToStruct(jsonSrc string, obj interface{}) error {
	data, readErr := ioutil.ReadFile(jsonSrc)
	if readErr != nil {
		return readErr
	}
	parseErr := json.Unmarshal(data, obj)
	if parseErr != nil {
		return parseErr
	}
	return nil
}

// 将结构体序列化成JSON并保存到文件
func (i *IOUtils) StructToJsonFile(jsonSrc string, obj interface{}) error {
	data, parseErr := json.Marshal(obj)
	if parseErr != nil {
		return parseErr
	}
	return i.ReplaceStrToFile(string(data), jsonSrc)
}

// 获取运行时目录
func (i *IOUtils) GetRuntimePath(filename string) string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(dir, filename)
}

// 将内容覆盖到指定文件，如果文件不存在那么创建，如果文件存在那么覆盖
func (i *IOUtils) ReplaceStrToFile(content, path string) error {
	if i.PathExists(path) {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, wErr := file.WriteString(content)
	if wErr != nil {
		return wErr
	}
	return nil
}

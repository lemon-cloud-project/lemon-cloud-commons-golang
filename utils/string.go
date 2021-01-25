package utils

import (
	"errors"
	"github.com/satori/go.uuid"
	"strings"
	"sync"
)

type StringUtils struct {
}

var stringInstance *StringUtils
var stringOnce sync.Once

func String() *StringUtils {
	stringOnce.Do(func() {
		stringInstance = &StringUtils{}
	})
	return stringInstance
}

// 生成UUID
func (i *StringUtils) Uuid(withoutLine bool) string {
	if withoutLine {
		return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	}
	return uuid.NewV4().String()
}

// 获取文件的后缀名
func (i *StringUtils) GetFileSuffixName(fileName string) (error, string) {
	var fileFormatArr = [...]string{".cpio.gz", ".cpio.cgz", ".tar.gz", ".tar.bz", ".tar.bz2", ".tar.xz", ".tar.z"}
	for _, v := range fileFormatArr {
		if strings.Index(fileName, v) != -1 {
			return nil, v
		}
	}
	if strings.Index(fileName, ".") == -1 {
		return errors.New("Unable to get file type"), ""
	}
	endIndex := strings.LastIndex(fileName, ".")
	fileType := fileName[endIndex:len(fileName)]
	return nil, fileType
}

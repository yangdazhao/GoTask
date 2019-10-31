package Job

import (
	"GoTask/LogExt"
	"github.com/beego/bee/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CleanDir(path string) {
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".dat") &&
			strings.HasPrefix(info.Name(), "20") {
			Last := time.Now().AddDate(0, -1, 0)
			if Last.After(info.ModTime()) {
				_ = os.Remove(path)
			}
		}
		return nil
	})
}

type CleanJob struct {
	SrcPath string
}

func (this CleanJob) Run() {
	LogExt.LogDebug("BackFileJob")
	if utils.IsExist(this.SrcPath) {
		CleanDir(this.SrcPath)
	} else {
		LogExt.LogWarn(this.SrcPath + "文件不存在")
	}
}

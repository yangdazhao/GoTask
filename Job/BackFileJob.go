package Job

import (
	"GoTask/LogExt"
	"fmt"
	"github.com/beego/bee/utils"
	"io"
	"os"
	"time"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	LogExt.LogInfo(dstFileName)
	LogExt.LogInfo(srcFileName)
	srcFile, err := os.Open(srcFileName)

	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer srcFile.Close()

	//打开dstFileName

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)

}

type BackFileJob struct {
	SrcPath string
}

func (this BackFileJob) Run() {
	LogExt.LogDebug("BackFileJob")
	fmt.Println("testJob2...")
	if utils.IsExist(this.SrcPath) {
		//2006-01-02 15:04:05
		fmt.Println(CopyFile(time.Now().Format("20060102T150405")+"_"+this.SrcPath, this.SrcPath))
	} else {
		LogExt.LogWarn(this.SrcPath + "文件不存在")
	}
}

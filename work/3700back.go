package work

import (
	"GoTask/Job"
	"encoding/xml"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"strings"
	"time"
)

type Config struct {
	XMLName xml.Name `xml:"Config"`
	Group   []string `xml:"GroupName"`
}

func AR3700BackDat() {
	fmt.Println("Yang")
	c := cron.New()
	spec := "0 0 4 * * ?"

	var _config Config
	content, _ := ioutil.ReadFile(".\\Config.xml")
	_ = xml.Unmarshal(content, &_config)
	for _, v := range _config.Group {

		Group := v + ".dat"
		fmt.Println(Group)
		_, _ = Job.CopyFile(
			strings.ReplaceAll(time.Now().Format("20060102T150405_GROUP.dat"), "GROUP", v), Group)
		_ = c.AddJob(spec, Job.BackFileJob{v + ".dat"})
	}
	//启动计划任务
	c.Start()
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()
	select {}
}

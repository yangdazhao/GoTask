package work

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestAR3700BackDat(t *testing.T) {
	var _config Config
	content, _ := ioutil.ReadFile("..\\Config.xml")
	_ = xml.Unmarshal(content, &_config)

	for _, v := range _config.Group {
		t.Log(time.Now().Format("20060102T150405_GROUPNAME.dat"))
		t.Log(strings.ReplaceAll(time.Now().Format("20060102T150405_GROUPNAME.dat"), "GROUPNAME", v))
	}
}

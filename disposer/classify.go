package disposer

import (
	"encoding/base64"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/cns1rius/imgstore/config"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	types []string
)

func classify(path string) string {
	img, _ := os.ReadFile(path)
	b64Img := base64.StdEncoding.EncodeToString(img)
	payload := strings.NewReader(fmt.Sprintf("image=%s", b64Img))

	url := "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general?access_token=" + config.Conf.GetString("set.baidu_aip")
	res, err := http.Post(url, "application/x-www-form-urlencoded", payload)
	if !(err == nil) {
		return ""
	}
	body, _ := io.ReadAll(res.Body)

	root, _ := jsonparser.GetString(body, "result", "[0]", "root")
	for _, i := range types {
		if strings.Contains(root, i) {
			root = i
		}
	}
	// get 类型:root

	return path
	// todo ai.baidu.com api 识别 -> 匹配关键字 -> tmp转存(mv! not cp)目录(固定部分类别，未在分类列表内的放入其他)
}

package disposer

import (
	"encoding/base64"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/cns1rius/imgstore/config"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func classify(c *gin.Context, filePath []string, errors []string) {
	var typeTag string
	types := config.Types
	for _, path := range filePath {
		// ai.baidu.com api 识别 -> 匹配关键字
		img, _ := os.ReadFile(path)
		b64Img := base64.StdEncoding.EncodeToString(img)
		payload := strings.NewReader(fmt.Sprintf("image=%s", b64Img))

		url := "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general?access_token=" + config.Conf.GetString("set.baidu_aip")
		res, err := http.Post(url, "application/x-www-form-urlencoded", payload)
		if !(err == nil) {
			errors = append(errors, path)
			continue
		}
		body, _ := io.ReadAll(res.Body)

		resultNum, _ := jsonparser.GetInt(body, "result_num")
		for i := 0; i <= int(resultNum); i++ {
			if typeTag != "" {
				break
			}
			root, _ := jsonparser.GetString(body, "result", "[0]", "root")
			for _, j := range types {
				if strings.Contains(root, j) {
					typeTag = j
					break
				} else if strings.Contains(root, j) {
					typeTag = "物品"
					break
				}
			}
		}
		if typeTag == "" {
			typeTag = "其他"
		}
		// get 图片类型变量 typeTag
		// tmp转存(mv! not cp)目录
		newPath := path
		strings.Replace(newPath, "tmp", typeTag, 1)
		_ = os.Rename(path, newPath)
		// 传库
		_ = config.ImgUpdate(newPath, config.GetCookieId(c))
	}
	c.HTML(http.StatusOK, "user/login.tmpl", gin.H{"error": errors})
}

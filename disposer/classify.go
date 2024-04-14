package disposer

import (
	"encoding/base64"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/cns1rius/imgstore/config"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func classify(c *gin.Context, filePaths []string, errors []string) {
	var (
		typeTag string
		tmpTag  string
	)
	types := config.Types
	for _, path := range filePaths {
		// ai.baidu.com api 识别 -> 匹配关键字
		img, _ := os.ReadFile(path)
		b64Img := url.QueryEscape(base64.StdEncoding.EncodeToString(img))
		payload := strings.NewReader(fmt.Sprintf("image=%s", b64Img))

		Url := "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general?access_token=" + config.Conf.GetString("set.baidu_aip")
		res, err := http.Post(Url, "application/x-www-form-urlencoded", payload)
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
			root, _ := jsonparser.GetString(body, "result", "["+strconv.Itoa(i)+"]", "root")
			for _, j := range types {
				if strings.Contains(root, j) {
					typeTag = j
					break
				} else if strings.Contains(root, "商品") {
					tmpTag = "物品"
					break
				}
			}
		}
		if typeTag == "" {
			if tmpTag != "" {
				typeTag = tmpTag
			} else {
				typeTag = "其他"
			}
		}
		// get 图片类型变量 typeTag
		// tmp转存(mv! not cp)目录
		newPath := strings.Replace(path, "tmp", typeTag, 1)
		_ = os.Rename(path, newPath)
		// 传库
		_ = config.ImgUpdate(newPath, config.GetCookieId(c))
	}
	errorStr := strings.Join(errors, "\n")
	if errorStr == "" {
		errorStr = "Succeed!"
	}
	c.HTML(http.StatusOK, "root/redirect.tmpl", gin.H{"error": errorStr, "href": "返回主页"})
}

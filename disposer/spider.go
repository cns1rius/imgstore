package disposer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

func Spider(c *gin.Context) {
	var (
		imgUrls  []string
		errors   []string
		filePath []string
	)
	url := c.PostForm("url")
	matched, _ := regexp.MatchString(`\.(jpg|jpeg|png|gif|bmp|svg|ico|webp)$`, url)

	if matched {
		imgUrls = append(imgUrls, url)
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return
		}
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			src, exists := s.Attr("src")
			if exists {
				if !strings.Contains(src, "http") {
					src = url[0:len(url)-1] + src
				}
				if strings.Contains(src, "?") {
					src = strings.Split(src, "?")[0]
				}
				imgUrls = append(imgUrls, src)
			}
		})
	}

	for _, value := range imgUrls {
		tmpPath, err := downLoad("./img/tmp/", value)
		if err != nil {
			errors = append(errors, value)
		}
		filePath = append(filePath, tmpPath)
		c.JSON(http.StatusOK, gin.H{"已下载": filePath, "失败列表": errors})
	}
	// c.HTML(http.StatusOK, "user/login.tmpl", gin.H{"失败列表": errors})
	// 调用classify 然后传库
	classify(c, filePath, errors)
}

func downLoad(pwd string, url string) (string, error) {
	filePath := pwd + path.Base(url)
	v, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http get [%v] failed! %v", url, err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(v.Body)
	content, _ := io.ReadAll(v.Body)
	err = os.WriteFile(filePath, content, 0666)
	if err != nil {
		fmt.Printf("Save to file failed! %v", err)
		return "", err
	}
	return filePath, nil
}

// todo 存在任意文件上传 可能无法利用

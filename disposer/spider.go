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
			c.HTML(http.StatusOK, "root/redirect.tmpl", gin.H{"error": err, "href": "返回主页"})
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			c.HTML(http.StatusOK, "root/redirect.tmpl", gin.H{"error": err, "href": "返回主页"})
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

	w := c.Writer
	w.Header().Set("Transfer-Encoding", "chunked") // 设置响应头以启用 chunked transfer encoding
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusFound)
	_, _ = w.Write([]byte("<script>\nalert(\"Downloading...\");\n</script>"))
	w.(http.Flusher).Flush()

	for _, value := range imgUrls {
		tmpPath, err := downLoad("./img/tmp/", value)
		if err != nil {
			errors = append(errors, value)
			_, _ = w.Write([]byte(fmt.Sprintf("%v failed", value)))
			w.(http.Flusher).Flush()
			continue
		}
		filePath = append(filePath, tmpPath)

		_, _ = w.Write([]byte(fmt.Sprintf("%v succeed", filePath)))
		w.(http.Flusher).Flush()
	}
	// 调用classify 然后传库
	classify(c, filePath, errors)
}

func downLoad(pwd string, url string) (string, error) {
	filePath := pwd + path.Base(url)
	v, err := http.Get(url)
	if err != nil {
		_, _ = fmt.Fprint(gin.DefaultWriter, fmt.Sprintf("[GIN] Http get [%v] failed! %v", url, err))
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(v.Body)
	content, _ := io.ReadAll(v.Body)
	err = os.WriteFile(filePath, content, 0666)
	if err != nil {
		_, _ = fmt.Fprint(gin.DefaultWriter, fmt.Sprintf("[GIN] Save to file failed! %v", err))
		return "", err
	}
	return filePath, nil
}

// todo 存在任意文件上传 可能无法利用

package disposer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func spider(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)

	re := regexp.MustCompile(`https?://\s+?\.(?:jpg|jpeg|png|gif|bmp|svg)`)
	imgUrls := re.FindAllString(string(body), -1)
	for _, value := range imgUrls {
		path, err := downLoad("./pictures/tmp", value)
		if err != nil {
			return "", err
		} else {
			return path, nil
		}
	}
	return "", err
}

func downLoad(pwd string, url string) (string, error) {
	path := pwd
	idx := strings.Index(url, "/")
	if idx < 0 {
		path += "/" + url
	} else {
		path += url[idx:]
	}
	v, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http get [%v] failed! %v", url, err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(v.Body)
	content, _ := io.ReadAll(v.Body)
	if err != nil {
		fmt.Printf("Read http response failed! %v", err)
		return "", err
	}
	err = os.WriteFile(path, content, 0666)
	if err != nil {
		fmt.Printf("Save to file failed! %v", err)
		return "", err
	}
	return path, nil
}

// todo 存在任意文件上传 可能无法利用

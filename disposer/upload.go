package disposer

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Upload(c *gin.Context) {
	var (
		errs      []string
		filePaths []string
	)
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := "./img/tmp/" + file.Filename
		// todo 存在任意目录上传 可通过文件名覆盖 复制时能把关键文件复制过来
		if err := c.SaveUploadedFile(file, dst); err != nil {
			errs = append(errs, err.Error())
		}
		filePaths = append(filePaths, dst)
	}

	w := c.Writer
	w.Header().Set("Transfer-Encoding", "chunked") // 设置响应头以启用 chunked transfer encoding
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusFound)
	_, _ = w.Write([]byte("<script>\n  alert(\"Hello World!\");\n</script>\n"))
	w.(http.Flusher).Flush()
	//c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	classify(c, filePaths, errs)
}

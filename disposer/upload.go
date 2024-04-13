package disposer

import (
	"fmt"
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
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	// classify(c, filePaths, errs)
	classify(c, filePaths, errs)
}

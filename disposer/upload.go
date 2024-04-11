package disposer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Upload(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := "./img/tmp/"
		// todo 存在任意目录上传 可通过文件名覆盖 复制时能把关键文件复制过来
		if err := c.SaveUploadedFile(file, dst); err != nil {
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	// todo disposer.classify(path) path
	// todo Create ImgTable{ path, id/username}
}

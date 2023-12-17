package route

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UNDER CONSTRUCTION // // UNDER CONSTRUCTION // // UNDER CONSTRUCTION // // UNDER CONSTRUCTION //
func DownloadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileID := c.Param("id")
		fileName := c.Param("filename")

		filePath := filepath.Join("./uploads", fmt.Sprintf("%s-%s", fileID, fileName))
		file, err := os.Open(filePath)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}
		defer file.Close()
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Length", fmt.Sprintf("%d", getFileSize(file)))
		io.Copy(c.Writer, file)
	}
}

func getFileSize(file *os.File) int64 {
	stat, err := file.Stat()
	if err != nil {
		return 0
	}
	return stat.Size()
}

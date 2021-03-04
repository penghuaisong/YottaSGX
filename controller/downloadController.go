package controller

import (
	"github.com/gin-gonic/gin"
)

//DownloadFile 下载
func DownloadFile(g *gin.Context) {
	// bucketName := g.Query("bucketName")

	// fileName := g.Query("fileName")

	// publicKey := g.Query("publicKey")

	// // savePath := g.Query("path")

	// content := publicKey[3:]
	// c := api.GetClient(content)

	// download, err := c.NewDownloadFile(bucketName, fileName, primitive.NilObjectID)
	// if err != nil {
	// 	// logrus.Errorf("[DownloadFile ]AuthSuper ERR:%s\n", err)
	// }

	// if err != nil {
	// 	// logrus.Errorf("[DownloadFile ]AuthSuper ERR:%s\n", err)
	// }

	// errn := download.SaveToPath(savePath + "/" + fileName)
	// if errn != nil {
	// 	// logrus.Errorf("[DownloadFile ]AuthSuper ERR:%s\n", errn)
	// } else {
	// 	// logrus.Infof("[ " + fileName + " ]" + " is Download Success")
	// }

}

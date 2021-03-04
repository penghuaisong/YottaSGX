package controller

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"github.com/yottachain/YTCoreService/api"
	"github.com/yottachain/YTCoreService/env"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var download_progress_CACHE = cache.New(time.Duration(600000)*time.Second, time.Duration(600000)*time.Second)

//DownloadFile 下载
func DownloadFile(g *gin.Context) {
	defer env.TracePanic("DownloadFile")
	bucketName := g.Query("bucketName")

	fileName := g.Query("fileName")

	publicKey := g.Query("publicKey")

	savePath := g.Query("path")

	content := publicKey[3:]
	c := api.GetClient(content)

	download, err := c.NewDownloadFile(bucketName, fileName, primitive.NilObjectID)
	if err != nil {
		logrus.Errorf("[DownloadFile ]AuthSuper ERR:%s\n", err)
	}

	putDownloadObject(bucketName, fileName, publicKey, download)

	if err != nil {
		logrus.Errorf("[DownloadFile ]AuthSuper ERR:%s\n", err)
	}

	errn := download.SaveToPath(savePath + "/" + fileName)
	if errn != nil {
		logrus.Errorf("[DownloadFile ]AuthSuper ERR:%s\n", errn)
	} else {
		logrus.Infof("[ " + fileName + " ]" + " is Download Success")
	}

}

//putUploadObject 将上传实例加入到缓存中 用于进度查询
func putDownloadObject(bucketName, fileName, publicKey string, upload *api.DownloadObject) {

	key := bucketName + fileName + publicKey + "download"

	data := []byte(key)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	logrus.Infof("md5str set : %s", md5str)
	download_progress_CACHE.SetDefault(md5str, upload)
}

//GetDownloadProgress 查询上传进度
func GetDownloadProgress(g *gin.Context) {
	defer env.TracePanic("GetDownloadProgress")
	publicKey := g.Query("publicKey")
	bucketName := g.Query("bucketName")
	fileName := g.Query("fileName")

	ii := getDownloadProgress(bucketName, fileName, publicKey)

	g.String(http.StatusOK, strconv.FormatInt(int64(ii), 10))
}

//getDownloadProgress 查询进度
func getDownloadProgress(bucketName, fileName, publicKey string) int32 {
	var num int32
	key := bucketName + fileName + publicKey + "download"

	data := []byte(key)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	logrus.Infof("md5str get : %s", md5str)
	v, found := download_progress_CACHE.Get(md5str)

	logrus.Infof("key is value : \n", found)
	if found {
		ii := v.(*api.DownloadObject).GetProgress()
		num = ii
	} else {
		num = 0
	}
	return num
}

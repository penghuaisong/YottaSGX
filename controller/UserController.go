package controller

// import (
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sirupsen/logrus"
// 	"github.com/yottachain/YTCoreService/api"
// 	"github.com/yottachain/YTCoreService/env"
// )

// //User 用户注册
// type User struct {
// 	UserName   string `form:"userName" json:"userName" binding:"required"`
// 	PrivateKey string `form:"privateKey" json:"privateKey" xml:"privateKey" binding:"required"`
// }

// //Register 用户注册
// func Register(g *gin.Context) {
// 	// defer env.TracePanic("Register")
// 	var json User
// 	ii := 1

// 	if err := g.Bind(&json); err != nil {
// 		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	userName := json.UserName

// 	privateKey := json.PrivateKey

// 	var client *api.Client
// 	var err2 error
// 	for {
// 		// client, err2 = api.NewClient(userName, privateKey)
// 		client, err2 = api.NewClientV2(&env.UserInfo{
// 			UserName: userName,
// 			Privkey:  []string{privateKey}}, 3)
// 		if err2 != nil {
// 			ii++
// 			if ii <= 3 {
// 				time.Sleep(time.Second * 5)
// 			} else {
// 				logrus.Infof("err:%s\n", err2)
// 				break
// 			}
// 		} else {
// 			break
// 		}
// 	}
// 	if err2 != nil {
// 		logrus.Errorf("User Register Failed, %s\n", err2)
// 		g.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "Msg": "Register Failed!Please checked userName and privateKey "})
// 	} else {

// 		logrus.Infof("PrivateKey::::%s\n", client.AccessorKey)
// 		logrus.Infof("User Register Success,UserName: %s\n", userName)
// 		g.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Msg": "Register success " + userName})
// 	}
// }

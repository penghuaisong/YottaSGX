package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//User 用户注册
type User struct {
	UserName   string `form:"userName" json:"userName" binding:"required"`
	PrivateKey string `form:"privateKey" json:"privateKey" xml:"privateKey" binding:"required"`
}

//Register 用户注册
func Register(g *gin.Context) {
	// defer env.TracePanic("Register")
	var json User
	// ii := 1

	if err := g.Bind(&json); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userName := json.UserName

	privateKey := json.PrivateKey

	// var client *api.Client
	// var err2 error
	// for {
	// 	// client, err2 = api.NewClient(userName, privateKey)
	// 	client, err2 = api.NewClientV2(&env.UserInfo{
	// 		UserName: userName,
	// 		Privkey:  []string{privateKey}}, 3)
	// 	if err2 != nil {
	// 		ii++
	// 		if ii <= 3 {
	// 			time.Sleep(time.Second * 5)
	// 		} else {
	// 			logrus.Infof("err:%s\n", err2)
	// 			break
	// 		}
	// 	} else {
	// 		break
	// 	}
	// }
	// if err2 != nil {
	// 	logrus.Errorf("User Register Failed, %s\n", err2)
	// 	g.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "Msg": "Register Failed!Please checked userName and privateKey "})
	// } else {

	// 	logrus.Infof("PrivateKey::::%s\n", client.AccessorKey)
	// 	logrus.Infof("User Register Success,UserName: %s\n", userName)
	// 	g.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Msg": "Register success " + userName})
	// }
	g.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Msg": "Register success " + userName, "privateKey": privateKey})

}

func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func SaveFile(filename string, data []byte) bool {
	var err error
	fmt.Println("saveFile::::", filename)
	if len(filename) > 0 && data != nil {
		dir := filepath.Dir(filename)
		if MakeDir(dir) != nil {
			return false
		}

		err = ioutil.WriteFile(filename, data, 0666)
		if err != nil {
			fmt.Println("SaveFile err:", err)
		} else {
			return true
		}
	} else {
		fmt.Println("SaveFile err: wrong params")
	}
	return false

}

func ReadFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Fail to open file:", err)
	}
	defer file.Close()

	fmt.Println("fileName:", filename)
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			fmt.Println("Read EOF")
			break
		} else if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		fmt.Println(string(buf[:n]))
	}
	return buf
}

func GetInfo(g *gin.Context) {
	var msg string
	msg = "Test golang router...."
	g.JSON(http.StatusOK, gin.H{"msg": msg})
}

//WriteFile  test write file
func WriteFile(g *gin.Context) {
	directory := g.Query("path")
	fmt.Printf("path::::" + directory)

	isExist := CheckFileIsExist(directory)

	if isExist {
		data := ReadFile(directory)

		SaveFile("/mnt/test/mm", data)
	}

	g.JSON(http.StatusOK, gin.H{"msg": "test"})
}

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
	fmt.Println("文件路径：" + directory)

	isExist := CheckFileIsExist(directory)
	fmt.Println("文件是否存在:", isExist)

	if !isExist {
		data := ReadFile(directory)

		fmt.Println("file length:::", len(data))

		var d1 = []byte("sdssddssdsdssdsdsddsodos")
		err2 := ioutil.WriteFile("/mnt/test/m1", d1, 0666)
		if err2 != nil {

		}

		// SaveFile("./run/mount/tmp/m1",d1)
	}

	g.JSON(http.StatusOK, gin.H{"msg": "test"})
}

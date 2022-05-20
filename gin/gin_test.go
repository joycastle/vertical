package gin

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func Test_GIN(t *testing.T) {

	ret := make(chan bool, 1)

	//InitGinServer(gin.ReleaseMode)
	InitGinServer(gin.DebugMode)

	cfg := GinConf{
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	RegisterControler(GIN_GET, "/hi", func(c *gin.Context) {
		ret <- true
		c.JSON(200, gin.H{
			"message": "hi ok",
		})
	})

	RegisterControler(GIN_GET, "/hii", func(c *gin.Context) {
		panic("me panic")
	})

	RegisterControler(GIN_POST, "/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello ok",
		})
	})

	go StartGin(cfg, 8080)

	time.Sleep(time.Second)

	go http.Get("http://127.0.0.1:8080/hi")

	select {
	case <-ret:
	case <-time.NewTimer(time.Second * 2).C:
		t.Logf("Time Out Gin Test")
		t.Fail()
	}
}

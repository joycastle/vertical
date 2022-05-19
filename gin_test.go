package vertical

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func Test_GIN(t *testing.T) {
	return

	ret := make(chan bool, 1)

	InitGinServer(gin.ReleaseMode)

	SetGinMiddware(func(c *gin.Context) {
		fmt.Println("1111")
		c.Next()
		fmt.Println("111111111111")
	})

	SetGinMiddware(func(c *gin.Context) {
		fmt.Println("2222")
		c.Next()
		fmt.Println("22222222222")
	})

	SetGinMiddware(func(c *gin.Context) {
		fmt.Println("3333")
		c.Next()
		fmt.Println("33333333")
	})

	SetGinMiddware(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "参数错苏",
		})
		c.Abort()
		fmt.Println("444444")
	})

	cfg := GinConf{
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	RegisterControler(GIN_GET, "/hi", func(c *gin.Context) {
		fmt.Println("AAAAAA")
	}, func(c *gin.Context) {
		fmt.Println("555555555555555")
		c.JSON(200, gin.H{
			"message": "hi ok",
		})
		fmt.Println("666666666666666")
	})

	RegisterControler(GIN_GET, "/hii", func(c *gin.Context) {
		ret <- true
		panic("me panic")
	})

	RegisterControler(GIN_GET, "/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello ok",
		})
	})

	go StartGin(cfg, 8080)

	<-ret

	time.Sleep(time.Second * 1)

	go http.Get("http://127.0.0.1:8080/hi")

	select {
	case <-ret:
	case <-time.NewTimer(time.Second * 2).C:
		t.Logf("Time Out Gin Test")
		t.Fail()
	}
}

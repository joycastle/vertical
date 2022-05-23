package gin

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joycastle/vertical/log"
)

const (
	GIN_POST   = "POST"
	GIN_GET    = "GET"
	GIN_DELETE = "DELETE"
	GIN_PUT    = "PUT"
	GIN_HEAD   = "HEAD"
)

type GinConf struct {
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

type GinController struct {
	Method       string
	RelativePath string
	Handlers     []gin.HandlerFunc
}

var (
	ginRouter           *gin.Engine
	ginServer           *http.Server
	ginControllers      []GinController
	ginNoRouterFunc     gin.HandlerFunc //noRouter process method
	ginPanicFunc        gin.HandlerFunc //panic process method
	ginMiddwareHandlers []gin.HandlerFunc
)

func InitGinServer(runMode string) {

	gin.SetMode(runMode)

	ginNoRouterFunc = func(context *gin.Context) {
		if strings.Contains(context.Request.URL.Path, "/system/ping") {
			context.String(200, "Ping")
		} else {
			context.String(404, "page not found")
		}
	}

	ginPanicFunc = gin.RecoveryWithWriter(log.GetLogger("error").Fptr)

	if runMode == gin.DebugMode {
		ginMiddwareHandlers = append(ginMiddwareHandlers, gin.Logger())
	}

	ginMiddwareHandlers = append(ginMiddwareHandlers, ginPanicFunc)
}

func RegisterControler(method string, relativePath string, handlers ...gin.HandlerFunc) {
	var ginController GinController

	ginController.Method = method
	ginController.RelativePath = relativePath
	ginController.Handlers = append(ginController.Handlers, handlers...)

	ginControllers = append(ginControllers, ginController)
}

func registerControler(method string, relativePath string, handlers ...gin.HandlerFunc) {
	switch method {
	case GIN_POST:
		ginRouter.POST(relativePath, handlers...)
	case GIN_GET:
		ginRouter.GET(relativePath, handlers...)
	case GIN_DELETE:
		ginRouter.DELETE(relativePath, handlers...)
	case GIN_PUT:
		ginRouter.PUT(relativePath, handlers...)
	case GIN_HEAD:
		ginRouter.HEAD(relativePath, handlers...)
	default:
		panic(fmt.Sprintf("not existe method:[%s]", method))
	}
}

func SetGinMiddware(middleware ...gin.HandlerFunc) {
	ginMiddwareHandlers = append(ginMiddwareHandlers, middleware...)
}

func StartGin(ginCfg GinConf, port int) {
	ginRouter = gin.New()

	ginRouter.Use(ginMiddwareHandlers...)

	ginRouter.NoRoute(ginNoRouterFunc)

	for _, controller := range ginControllers {
		registerControler(controller.Method, controller.RelativePath, controller.Handlers...)
	}

	ginServer = &http.Server{
		Addr:           "0.0.0.0:" + fmt.Sprintf("%d", port),
		Handler:        ginRouter,
		ReadTimeout:    ginCfg.ReadTimeout,
		WriteTimeout:   ginCfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.GetLogger("run").Infof("Gin Runing %s", "0.0.0.0:"+fmt.Sprintf("%d", port))
		if err := ginServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.GetLogger("error").Warnf("Gin Starting Failed: %s", err)
		}
	}()

	quit := make(chan os.Signal)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.GetLogger("run").Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ginServer.Shutdown(ctx); err != nil {
		log.GetLogger("run").Infof("Server Shutdown: %s", err.Error())
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.GetLogger("run").Infof("timeout of 5 seconds.")
	}

	log.GetLogger("run").Infof("Server exiting")
}

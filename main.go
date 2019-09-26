package main

import (
	"fmt"
	"github.com/ECEHive/myHive-backend/controller"
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/middleware"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logger = util.GetLogger("MAIN")

func main() {
	e := godotenv.Load()
	if e != nil {
		logger.Warning(".env file not existing or bad format", e)
	}
	// Connect to DB
	_ = db.GetDB()
	// DB: Migrate Entities
	entity.MigrateEntities()

	// Gin Setup
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler) // Handle ctx.error

	// Logging Setup
	httpLogger := util.GetLogger("Gin")
	gin.DefaultWriter = httpLogger.WriterLevel(logrus.InfoLevel)
	gin.DefaultErrorWriter = httpLogger.WriterLevel(logrus.ErrorLevel)
	r.Use(func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		end := time.Now()
		execDuration := end.Sub(start)
		path := ctx.Request.URL.Path
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		responseStatus := ctx.Writer.Status()
		httpLogger.WithField("HIVE-Log-Id", ctx.GetString("LogId")).
			WithField("Duration", execDuration.Seconds()*1000).
			WithField("UA", ctx.Request.Header.Get("User-Agent")).
			WithField("RequestPath", path).
			WithField("RequestMethod", method).
			WithField("ResponseStatus", responseStatus).
			Infof("%s %s | %d | %s - (%.4fms)",
				method, path, responseStatus, clientIP, execDuration.Seconds()*1000)
	})

	// Setup low priority middlewares
	r.Use(middleware.PaginationResolver) // This resolves pagination

	// Setup Routes
	controller.ConfigureUserController(r.Group("/user"))
	controller.ConfigureWorkbenchRoutes(r.Group("/workbench"))
	controller.ConfigureInventoryRouter(r.Group("/inventory"))

	// Setup 404 - No Route
	r.NoRoute(func(c *gin.Context) {
		c.Set("error", model.NotFound(fmt.Sprintf("route '%s' has no matching handler", c.Request.URL.Path)))
	})

	port := os.Getenv("serverport")
	if port == "" {
		port = "9000"
	}

	if err := r.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		logger.Error("error running the server", err)
	} // listen and serve on :port
}

package main

import (
	controllers "api-menu/controllers"
	"api-menu/middlewares"
	"api-menu/models"
	"api-menu/utils"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	//newrelic "github.com/newrelic/go-agent"
	_ "github.com/newrelic/go-agent/v3/integrations/nrgin"
)

// var newRelicApp *newrelic.Application

//var runMode string

//var isInitializeDB = flag.Bool("initDB", false, "initialize data")
//var isAutoMigrateDB = flag.Bool("autoDB", false, "auto migration database")

// var isTest = flag.Bool("isTest", false, "flag for testing")

// func init(){
//  	flag.Parse()

// }
//var newRelicApp *newrelic.Application

func ConvertObjectToJSONString(data interface{}) string {
	byteData, _ := json.Marshal(data)
	return string(byteData)
}

// func setupRoutes(){
// 	http.HandleFunc("/upload", controllers.UploadImage)
// }

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	//r.Use(nrgin.Middleware(newRelicApp))
	fmt.Println("main")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorMessage("endpoint not found", 404))
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, utils.ErrorMessage("method not allow", 405))
	})

	//initial
	utils.InitialDB()
	//if *isAutoMigrateDB == true {
	// a := controllers.App{}
	// a.Init()
	models.DoautoMigration(utils.MySQLDB)
	//models.AddForeignKey(utils.MySQLDB)
	controllers.InitialFirebaseStorage()
	//controllers.Init()
	controllers.PromotionEndpoint(r.Group("/pro"))
	controllers.ImageFirebaseEndpoint(r.Group("/image"))
	controllers.CreateFoodEndpoints(r.Group("/food"))
	controllers.CreateBillingEndpoints(r.Group("/bill"))
	controllers.TableEnpoint(r.Group("/table"))
	//controllers.CreateImageEnd(r.Group("/img"))
	controllers.CreateReviewEndpoints(r.Group("/reviews"))
	//fmt.Println("auto")
	// r.
	// http.HandleFunc("/up",models.UploadImage)
	//}
	
	//setupRoutes()
	// controllers.CreateFoodEndpoints(r.Group("/food"))
	// controllers.CreateIngredientsEndpoints(r.Group("/ingredients"))
	// controllers.CreateOderDetailEndpoints(r.Group("/Order"))
	// controllers.CreateTypeFoodEndpoints(r.Group("/type"))
	// controllers.CreatebillingEndpoints(r.Group("/Billing"))

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(rw, "kuay")
	// })
	// http.ListenAndServe(":8081", nil)
	{
		// r.GET("/sleep", func(c *gin.Context) {
		// 	time.Sleep(10 * time.Second)
		// 	c.String(http.StatusOK, "Welcome Gin Server")
		// })

		targetPort, isSet := os.LookupEnv("PORT")
		if len(targetPort) == 0 || isSet == false {
			targetPort = "8081"
		}
		srv := &http.Server{
			Addr:    ":" + targetPort,
			Handler: r,
		}

		// Initializing the server in a goroutine so that
		// it won't block the graceful shutdown handling below
		go func() {
			fmt.Println("Running on Port: ", targetPort)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}

		log.Println("Server exiting")
	}

}

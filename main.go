package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/agstyogottulen/clean-arc-lion/common"
	courierHttp "github.com/agstyogottulen/clean-arc-lion/courier/delivery/http"
	_courierRepository "github.com/agstyogottulen/clean-arc-lion/courier/repository"
	_courierService "github.com/agstyogottulen/clean-arc-lion/courier/service"
	"github.com/agstyogottulen/clean-arc-lion/models"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := gorm.Open("postgres", connStr)
	if err != nil {
		logrus.Error(err)
	}

	err = dbConn.DB().Ping()
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println("ping from db")

	defer func() {
		err = dbConn.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	dbConn.Debug().AutoMigrate(
		&models.Courier{},
	)

	r := mux.NewRouter()

	r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := common.Message(true, "Welcome to courier service api")
		common.Response(w, response)
		return
	}))).Methods(http.MethodGet)

	r.Handle("/ping", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := common.Message(true, "pong from courier service api")
		common.Response(w, response)
		return
	}))).Methods(http.MethodGet)

	courierRepository := _courierRepository.NewCourierRepository(dbConn)
	courierService := _courierService.NewCourierService(courierRepository)
	courierHttp.NewCourierHandler(r, courierService)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := common.Message(false, "url not found")
		w.WriteHeader(http.StatusNotFound)
		common.Response(w, response)
		return
	})

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"})

	logrus.Fatal(http.ListenAndServe(viper.GetString("server.address"), handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}

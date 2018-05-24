package main

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"go-it-asset-management/model"
	"go-it-asset-management/service"
	"go-it-asset-management/repository"
	"github.com/gorilla/mux"
	"go-it-asset-management/controller"
	"net/http"
	"log"
)

func main(){
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=assetMonitor password=example sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	db.AutoMigrate(&model.Application{})
	applicationService := service.ApplicationService{ApplicationRepository:repository.ApplicationRepository{Db:db}}
	applicationController := &controller.ApplicationController{ApplicationService:applicationService}

	defer db.Close()

	router := mux.NewRouter()
	router.Path("/api/applications").Methods("GET").Name("Get").HandlerFunc(applicationController.Get)
	router.Path("/api/applications").Methods("POST").Name("Sync").HandlerFunc(applicationController.Sync)

	log.Fatal(http.ListenAndServe(":8000", router))
}



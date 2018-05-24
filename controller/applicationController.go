package controller

import (
	"go-it-asset-management/service"
	"net/http"
	"encoding/json"
)

type ApplicationController struct{
	ApplicationService service.ApplicationService
}

func(c ApplicationController) Get(w http.ResponseWriter, req *http.Request){
	result := c.ApplicationService.GetAll()
	data, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func(c ApplicationController) Sync(w http.ResponseWriter, req *http.Request){
	c.ApplicationService.CreateOrUpdateApplicationInstalled()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("done"))
	return
}



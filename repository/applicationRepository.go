package repository

import (
	"go-it-asset-management/model"
	"github.com/jinzhu/gorm"
)

type ApplicationRepository struct {
	Db *gorm.DB
}


func (a ApplicationRepository) Insert(application *model.Application)  model.Application{
	a.Db.Create(application)
	return *application
}

func (a ApplicationRepository) GetAll() model.Applications{
	applications := model.Applications{}
	a.Db.Find(&applications)
	return applications
}
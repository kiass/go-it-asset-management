package service

import (
	"go-it-asset-management/repository"
	"go-it-asset-management/model"
	"fmt"
)

type ApplicationService struct {
	ApplicationRepository repository.ApplicationRepository
}

func(a ApplicationService) GetAll() model.Applications{
	return a.ApplicationRepository.GetAll()
}

func (a ApplicationService) CreateOrUpdateApplicationInstalled(){
	for _, d := range GetInstalledApp()[2:]{
		applicationDetail := CsvByLine(d)
		fmt.Println(d)
		a.ApplicationRepository.Insert(&model.Application{
			Name : applicationDetail[2],
			Version: applicationDetail[4],
			InstalledDate: applicationDetail[1],
		});
	}
}
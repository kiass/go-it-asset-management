package model

import "github.com/jinzhu/gorm"

type Application struct {
	gorm.Model
	Name string `json:name`
	Version string `json:version`
	InstalledDate string `json:installedDate`
}



type Applications []Application
package database

import (
	"presence-app-backend/configs"
	"presence-app-backend/models"
)

func CreateDepartment(department *models.Department) error {
	if e := configs.DB.Save(department).Error; e != nil {
		return e
	}
	return nil
}

func GetDepartments() (interface{}, error) {
	var deparments models.Department
	if e := configs.DB.Find(&deparments).Error; e != nil {
		return nil, e
	}
	return deparments, nil
}

func UpdateDepartment(department *models.Department) error {
	if e := configs.DB.Save(department).Error; e != nil {
		return e
	}
	return nil
}

func DeleteDepartment(department *models.Department) error {
	if e := configs.DB.Delete(department).Error; e != nil {
		return e
	}
	return nil
}

package departments

import (
	"context"
	"gorm.io/gorm"
	dpt "presence-app-backend/business/departments"
)

type MysqlDepartmentRepository struct {
	Conn *gorm.DB
}

func NewMysqlDepartmentRepository(conn *gorm.DB) dpt.Repository {
	return &MysqlDepartmentRepository{Conn: conn}
}

func (repo *MysqlDepartmentRepository) GetAll(ctx context.Context) ([]dpt.Domain, error) {
	var departments []Department
	result := repo.Conn.Find(&departments)

	if result.Error != nil {
		return []dpt.Domain{},result.Error
	}

	var convertedDepartment []dpt.Domain

	for _, val := range departments {
		convertedDepartment = append(convertedDepartment, val.ToDomain())
	}
	return convertedDepartment, nil
}

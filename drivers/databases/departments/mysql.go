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
	return &MysqlDepartmentRepository{
		Conn: conn,
	}
}

func (repo *MysqlDepartmentRepository) GetAll(ctx context.Context) ([]dpt.Domain, error) {
	var departments []Department
	result := repo.Conn.Find(&departments)

	if result.Error != nil {
		return []dpt.Domain{}, result.Error
	}

	var convertedDepartment []dpt.Domain

	for _, val := range departments {
		convertedDepartment = append(convertedDepartment, val.ToDomain())
	}
	return convertedDepartment, nil
}

func (repo *MysqlDepartmentRepository) Store(ctx context.Context, department *dpt.Domain) (dpt.Domain, error) {
	var result = FromDomain(department)
	if err := repo.Conn.Save(&result).Error; err != nil {
		return dpt.Domain{}, err
	}
	return result.ToDomain(), nil
}

func (repo *MysqlDepartmentRepository) GetById(ctx context.Context, id int) (dpt.Domain, error) {
	var department Department

	if err := repo.Conn.First(&department, id).Error; err != nil {
		return dpt.Domain{}, err
	}
	return department.ToDomain(), nil
}

func (repo *MysqlDepartmentRepository) Update(ctx context.Context, department *dpt.Domain) (dpt.Domain, error) {
	departmentFromDb := FromDomain(department)
	if err := repo.Conn.Save(departmentFromDb).Error; err != nil {
		return dpt.Domain{}, err
	}
	return departmentFromDb.ToDomain(), nil

}

func (repo *MysqlDepartmentRepository) Delete(ctx context.Context, id int) error {
	err := repo.Conn.Delete(&Department{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

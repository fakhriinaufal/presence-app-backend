package schedules

import (
	"context"
	"gorm.io/gorm"
	"presence-app-backend/business/schedules"
)

type MysqlScheduleRepository struct {
	Conn *gorm.DB
}

func NewMysqlScheduleRepository(conn *gorm.DB) schedules.Repository {
	return &MysqlScheduleRepository{
		Conn: conn,
	}
}

func (repo *MysqlScheduleRepository) Store(ctx context.Context, domain *schedules.Domain) (schedules.Domain, error) {
	result := FromDomain(domain)
	if err := repo.Conn.Create(&result).Error; err != nil {
		return schedules.Domain{}, err
	}
	return result.ToDomain(), nil
}

func (repo *MysqlScheduleRepository) GetAll(ctx context.Context) ([]schedules.Domain, error) {
	var schedulesFromDB []Schedule
	if err := repo.Conn.Find(&schedulesFromDB).Error; err != nil {
		return []schedules.Domain{}, err
	}

	var result []schedules.Domain
	for _, val := range schedulesFromDB {
		result = append(result, val.ToDomain())
	}
	return result, nil
}

func (repo *MysqlScheduleRepository) GetById(ctx context.Context, id int) (schedules.Domain, error) {
	var result Schedule
	if err := repo.Conn.Find(&result, id).Error; err != nil {
		return schedules.Domain{}, err
	}
	return result.ToDomain(), nil
}

func (repo *MysqlScheduleRepository) Update(ctx context.Context, domain *schedules.Domain) (schedules.Domain, error) {
	result := FromDomain(domain)
	if err := repo.Conn.Save(&result).Error; err != nil {
		return schedules.Domain{}, err
	}
	return result.ToDomain(), nil
}
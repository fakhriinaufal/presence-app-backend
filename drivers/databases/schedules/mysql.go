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

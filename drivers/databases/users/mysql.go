package users

import (
	"gorm.io/gorm"
	"presence-app-backend/business/users"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		conn,
	}
}

func (m MysqlUserRepository) Store(domain *users.Domain) (users.Domain, error) {
	var user = FromDomain(domain)
	err := m.Conn.Create(&user).Error

	if err != nil {
		return users.Domain{}, err
	}
	return user.ToDomain(), nil
}


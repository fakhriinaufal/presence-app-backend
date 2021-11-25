package users

import (
	"context"
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

func (m MysqlUserRepository) Store(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	var user = FromDomain(domain)
	err := m.Conn.Create(&user).Error

	if err != nil {
		return users.Domain{}, err
	}
	return user.ToDomain(), nil
}

func (m MysqlUserRepository) GetAll() ([]users.Domain, error) {
	var usersFromDB []User

	err := m.Conn.Find(&usersFromDB).Error

	if err != nil {
		return []users.Domain{}, err
	}

	return ToArrayOfDomain(&usersFromDB), nil
}

func (m MysqlUserRepository) GetById(ctx context.Context, id int) (users.Domain, error) {
	var user User
	if err := m.Conn.First(&user, id).Error; err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (m MysqlUserRepository) Update(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	user := FromDomain(domain)
	if err := m.Conn.Save(&user).Error; err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil

}


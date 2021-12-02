package presences

import (
	"context"
	"presence-app-backend/business/presences"

	"gorm.io/gorm"
)

type MysqlPresenceRepository struct {
	Conn *gorm.DB
}

func NewMysqlPresenceRepository(conn *gorm.DB) presences.Repository {
	return &MysqlPresenceRepository{
		Conn: conn,
	}
}

func (repo *MysqlPresenceRepository) Store(ctx context.Context, domain *presences.Domain) (presences.Domain, error) {
	result := FromDomain(domain)
	if err := repo.Conn.Create(&result).Error; err != nil {
		return presences.Domain{}, err
	}
	return result.ToDomain(), nil
}

func (repo *MysqlPresenceRepository) GetAll(ctx context.Context) ([]presences.Domain, error) {
	var result []Presence
	if err := repo.Conn.Preload("User").Preload("Schedule").Find(&result).Error; err != nil {
		return []presences.Domain{}, err
	}
	var domainResult []presences.Domain
	for _, val := range result {
		domainResult = append(domainResult, val.ToDomain())
	}
	return domainResult, nil
}

func (repo *MysqlPresenceRepository) GetById(ctx context.Context, id int) (presences.Domain, error) {
	var result Presence
	if err := repo.Conn.Preload("User").Preload("Schedule").First(&result, id).Error; err != nil {
		return presences.Domain{}, err
	}
	return result.ToDomain(), nil
}

func (repo *MysqlPresenceRepository) Update(ctx context.Context, domain *presences.Domain) (presences.Domain, error) {
	result := FromDomain(domain)
	if err := repo.Conn.Save(&result).Error; err != nil {
		return presences.Domain{}, err
	}
	return result.ToDomain(), nil
}

func (repo *MysqlPresenceRepository) Delete(ctx context.Context, id int) error {
	if err := repo.Conn.Delete(&Presence{}, id).Error; err != nil {
		return err
	}
	return nil
}

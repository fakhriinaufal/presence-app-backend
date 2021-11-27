package presences

import (
	"context"
	"gorm.io/gorm"
	"presence-app-backend/business/presences"
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
	if err := repo.Conn.Find(&result).Error; err != nil {
		return []presences.Domain{}, err
	}
	var domainResult []presences.Domain
	for _, val := range result {
		domainResult = append(domainResult, val.ToDomain())
	}
	return domainResult, nil
}

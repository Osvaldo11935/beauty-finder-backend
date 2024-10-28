package repositories_common

import (
	interfaces_repositories_common "src/internal/domain/interfaces_repositories/common"
	"src/internal/persistence/database"
	"gorm.io/gorm"
)

type GormBaseRepository struct {
	db  *gorm.DB
}

func NewGormBaseRepository() interfaces_repositories_common.IGormBaseRepository[gorm.DB]{
	db, _ := database.Connect()
   return &GormBaseRepository{db: db}
}

func(g * GormBaseRepository) Query() *gorm.DB{
	return g.db
}

func(g *GormBaseRepository) Insert(entity interface{}) error{
     return g.db.Create(entity).Error
}

func(g *GormBaseRepository) Update(entity interface{}) error{
	return g.db.Save(entity).Error
}

func(g *GormBaseRepository) Remove(entity interface{}) error{
	return g.db.Delete(entity).Error
}
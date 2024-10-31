package repositories_common

import (
	interfaces_repositories_common "src/internal/domain/interfaces_repositories/common"
	"gorm.io/gorm"
)

type GormBaseRepository struct {
	db  *gorm.DB
}

func NewGormBaseRepository(db *gorm.DB) interfaces_repositories_common.IGormBaseRepository[gorm.DB]{
  
	if db ==  nil {
	 return nil
   }

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
package database

import (
	"log"
	"src/internal/domain/entities"
	"src/internal/domain/object_values"
	"src/internal/domain/primitives"
	"time"
	"gorm.io/gorm"
)

func Seed() {
	db, connectionErr := Connect()
	if connectionErr != nil {
		return
	}
	SeedRole(db)
	SeedRatingType(db)
	SeedStatusType(db)
}

func SeedRole(db *gorm.DB) {
	roles := []entities.Role{
		{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.ROLE_ADMIN_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: object_values.ROLE_ADMIN_NAME,
		},
		{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.ROLE_CLIENT_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: object_values.ROLE_CLIENT_NAME,
		},
		{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.ROLE_SERVICE_PROVIDER_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: object_values.ROLE_SERVICE_PROVIDER_NAME,
		},
	}

	for _, rl := range roles {
		err := db.FirstOrCreate(&rl).Error
		if err != nil {
			log.Printf("Failed to seed status type %s: %v", rl.Name, err)
		}
	}
}

func SeedRatingType(db *gorm.DB) {
	roles := []entities.RatingType{
		{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.RATING_TYPE_BAD_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Type: object_values.RATING_TYPE_BAD_NAME,
		},
		{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.RATING_TYPE_NORMAL_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Type: object_values.RATING_TYPE_NORMAL_NAME,
		},
		{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.RATING_TYPE_EXCELLENT_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Type: object_values.RATING_TYPE_EXCELLENT_NAME,
		},
	}

	for _, rl := range roles {
		err := db.FirstOrCreate(&rl).Error
		if err != nil {
			log.Printf("Failed to seed status type %s: %v", rl.Type, err)
		}
	}
}

func SeedStatusType(db *gorm.DB) {

	statusTypes := []entities.AppointmentStatus{
		{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.STATUS_CANCELLED_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Type: object_values.STATUS_CANCELLED_NAME,
		},
		{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.STATUS_COMPLETED_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Type: object_values.STATUS_COMPLETED_NAME,
		},{
			BaseAuditableEntity: primitives.BaseAuditableEntity{
				BaseEntity: primitives.BaseEntity{
					ID: object_values.STATUS_PENDING_ID,
				},
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Type: object_values.STATUS_PENDING_NAME,
		},
	}

	for _, statusType := range statusTypes {
		err := db.FirstOrCreate(&statusType).Error
		if err != nil {
			log.Printf("Failed to seed status type %s: %v", statusType.Type, err)
		}
	}
}

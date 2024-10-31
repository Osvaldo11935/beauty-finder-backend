package database

import (
	"log"
	"src/internal/domain/entities"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	if db == nil {
		return
	}
	migrationErr := db.AutoMigrate(
		entities.AppointmentStatus{},
		entities.ServiceCategory{},
		entities.Service{},
		entities.Appointment{},
		entities.Role{},
		entities.User{},
		entities.FcmToken{},
		entities.Address{},
		entities.AttachmentType{},
		entities.Attachment{},
		entities.Message{},
		entities.Person{},
		entities.ServicePrice{},
		entities.ServiceProvider{},
		entities.RatingType{},
		entities.UserRating{},
	)

	if migrationErr != nil {
		log.Printf("Erro ao executar migrations", migrationErr)
		return
	}

}

package database

import (
	"log"
	"src/internal/domain/entities"
)

func RunMigration() {
	db, connectErr := Connect()

	if connectErr != nil {
		log.Panic("Erro ao conectar no banco de dados", connectErr)
		return
	}
	migrationErr := db.AutoMigrate(
		entities.AppointmentStatus{},
		entities.ServiceCategory{},
		entities.Service{},
		entities.Appointment{},
		entities.Role{},
		entities.User{},
		entities.Address{},
		entities.AttachmentType{},
		entities.Attachment{},
		entities.Message{},
		entities.Person{},
		entities.ServicePrice{},
		entities.ServiceProvider{},
	)

	if migrationErr != nil {
		log.Panic("Erro ao executar migrations", migrationErr)
		return
	}


}
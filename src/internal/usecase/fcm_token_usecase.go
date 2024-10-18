package usecase

import (
	"context"
	"log"

	"firebase.google.com/go/v4/messaging"
	fcm "github.com/appleboy/go-fcm"
	"github.com/google/uuid"
)

type FcmTokenUseCase struct {
	UserUseCase UserUseCase
}

func (s *FcmTokenUseCase) DispatchServiceNotification(ctx context.Context, serviceId uuid.UUID, appointmentId uuid.UUID) {

	users, findServiceErr := s.UserUseCase.FindUserByServiceId(serviceId)

	if findServiceErr != nil {
		return
	}
	for _, user := range users {
		for _, fcmToken := range user.FcmToken {
			client, clientErr := fcm.NewClient(ctx,
				fcm.WithCredentialsFile("../testnotification-6595e-fe94b46df7d9.json"))
			if clientErr != nil {
				log.Fatal(clientErr)
			}

			token := fcmToken.TokenFcm
			resp, err := client.Send(
				ctx,
				&messaging.Message{
					Token: token,
					Notification: &messaging.Notification{
						Title: "Confirmação de Serviço",
						Body:  "Você deseja confirmar o agendamento?",
					},
					Data: map[string]string{
						"appointment_id": appointmentId.String(),
						"action_type":    "confirm_service",  
					},
					Android: &messaging.AndroidConfig{
						Priority: "high",
					},
				},
			)
			if err != nil {
				log.Fatal(err)
			}
			print(resp)
		}
	}
}
package usecase

import (
	"context"
	"log"
	"src/internal/configs"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"

	"firebase.google.com/go/v4/messaging"
	fcm "github.com/appleboy/go-fcm"
	"github.com/google/uuid"
)

type FcmTokenUseCase struct {
	UserUseCase UserUseCase
}

func (s *FcmTokenUseCase) DispatchAssessmentNotification(ctx context.Context, userEvaluatorId uuid.UUID,
	request models_requests_posts.DispatchAssessmentNotification) {
	config, loadConfigErr := configs.LoadConfig()

	if loadConfigErr != nil {
		return
	}

	userEvaluator, findUserEvaluatorErr := s.UserUseCase.FindUserById(userEvaluatorId)

	if findUserEvaluatorErr != nil {
		return
	}

	if request.UserAvaluatedId != nil {
		return
	}

	userAvaluated, findUserAvaluatedErr := s.UserUseCase.FindUserById(*request.UserAvaluatedId)

	if findUserAvaluatedErr != nil {
		return
	}

	for _, fcmToken := range userEvaluator.FcmToken {
		client, clientErr := fcm.NewClient(ctx,
			fcm.WithCredentialsFile(config.FileConfigFirebase))
		if clientErr != nil {
			log.Fatal(clientErr)
		}

		token := fcmToken.TokenFcm
		resp, err := client.Send(
			ctx,
			&messaging.Message{
				Token: token,
				Notification: &messaging.Notification{
					Title: "Avaliaçã de Serviço",
					Body:  "Avalia o serviço do profissional",
				},
				Data: map[string]string{
					"user_evaluator_id": userEvaluator.ID.String(),
					"user_avaluated_id": userAvaluated.ID.String(),
					"action_type":       "confirm_service",
				},
				Android: &messaging.AndroidConfig{
					Priority: "high",
				},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		print(&resp)

	}
}

func (s *FcmTokenUseCase) DispatchServiceNotification(ctx context.Context, serviceId uuid.UUID, appointmentId uuid.UUID,
	request models_requests_posts.DispatchServiceNotification) {
	config, loadConfigErr := configs.LoadConfig()

	if loadConfigErr != nil {
		return
	}
	users, findServiceErr := s.UserUseCase.FindUsersNearBy(serviceId, request.Latitude, request.Longitude, request.RadiusKm)

	if findServiceErr != nil {
		return
	}

	for _, user := range users {
		for _, fcmToken := range user.User.FcmToken {
			client, clientErr := fcm.NewClient(ctx,
				fcm.WithCredentialsFile(config.FileConfigFirebase))
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

package usecase

import (
	"context"
	"log"
	"src/internal/configs"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	service_interface "src/internal/services/interface_services"

	"firebase.google.com/go/v4/messaging"
	fcm "github.com/appleboy/go-fcm"
	"github.com/google/uuid"
)

type FcmTokenUseCase struct {
	UserUseCase        UserUseCase
	FileManagerService service_interface.IFileManager
}

func (s *FcmTokenUseCase) DispatchAssessmentNotification(ctx context.Context, userEvaluatorId uuid.UUID,
	request models_requests_posts.DispatchAssessmentNotification) {
	loadEnv, loadConfigErr := configs.LoadConfig()

	if loadConfigErr != nil {
		return
	}
	
	file, getFileErr := s.FileManagerService.Download(loadEnv.SupaBaseFileConfigFireBase)
	if getFileErr != nil {
		log.Fatal("Falha ao carregar arquivo de configuração do Firebase:", getFileErr)
	}

	client, clientErr := fcm.NewClient(ctx, fcm.WithCredentialsJSON(file))
	if clientErr != nil {
		log.Fatal("Erro ao criar cliente FCM:", clientErr)
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
		if resp.FailureCount > 0 {
			log.Printf("Erro na resposta do FCM: %s, para o token: %s", resp.Responses, token)
		} else {
			log.Printf("Notificação enviada com sucesso para o token: %s, ID da mensagem: %s", token)
		}

	}
}

func (s *FcmTokenUseCase) DispatchServiceNotification(ctx context.Context, serviceId uuid.UUID, appointmentId uuid.UUID,
	request models_requests_posts.DispatchServiceNotification) {
	loadEnv, loadConfigErr := configs.LoadConfig()

	if loadConfigErr != nil {
		return
	}
	file, getFileErr := s.FileManagerService.Download(loadEnv.SupaBaseFileConfigFireBase)
	if getFileErr != nil {
		log.Fatal("Falha ao carregar arquivo de configuração do Firebase:", getFileErr)
	}

	users, findServiceErr := s.UserUseCase.FindUsersNearBy(serviceId, request.Latitude, request.Longitude, request.RadiusKm)
	if findServiceErr != nil {
		log.Println("Erro ao encontrar serviços:", findServiceErr)
		return
	}

	client, clientErr := fcm.NewClient(ctx, fcm.WithCredentialsJSON(file))
	if clientErr != nil {
		log.Fatal("Erro ao criar cliente FCM:", clientErr)
	}

	for _, user := range users {
		for _, fcmToken := range user.User.FcmToken {
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
				log.Println("Erro ao enviar notificação:", err)
				continue
			}

			if resp.FailureCount > 0 {
				log.Printf("Erro na resposta do FCM: %s, para o token: %s", resp.Responses, token)
			} else {
				log.Printf("Notificação enviada com sucesso para o token: %s, ID da mensagem: %s", token)
			}
		}
	}
}

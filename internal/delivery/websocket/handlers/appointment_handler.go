package handlers

import (
	"log"
	"src/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	 models_responses "src/internal/delivery/http/models/responses"
)

func FindAppointmentWebsocketHandler(ctx *gin.Context, _useCase usecase.AppointmentUseCase) {
	appointmentId:= uuid.MustParse(ctx.Query("appointmentId")) 
	
	conn, updgradeErr := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if updgradeErr != nil {
		log.Println("Erro ao fazer upgrade para WebSocket:", updgradeErr)
		return
	}
	defer conn.Close()

	data, _ := _useCase.FindAppointmentById(appointmentId)

	 resp := models_responses.ToAppointmentResponse(data)

	if err := conn.WriteJSON(resp); err != nil {
		log.Println("Erro ao enviar dados pelo WebSocket:", err)
	}
}
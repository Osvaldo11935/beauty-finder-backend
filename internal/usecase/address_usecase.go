package usecase

import (
	"context"
	"encoding/json"
	err "errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"src/internal/configs"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	models_responses "src/internal/delivery/http/models/responses"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressUseCase struct {
	HttpClientUseCase HttpClientUseCase
	Repo              interfaces_repositories.IAddressRepository
}

func (uc *AddressUseCase) InsertAddress(userId uuid.UUID, request models_requests_posts.CreateAddressRequest) (*uuid.UUID, error) {
	req := entities.NewAddress(userId, request)

	createErr := uc.Repo.Insert(&req)

	if createErr != nil {
		return nil, errors.UnknownCreateAddressError(createErr.Error())
	}

	return &req.ID, nil
}
func (uc *AddressUseCase) InsertAddressAppointment(appointmentId uuid.UUID, request models_requests_posts.CreateAddressRequest) (*uuid.UUID, error) {
	req := entities.NewAddressAppointment(appointmentId, request)

	createErr := uc.Repo.Insert(&req)

	if createErr != nil {
		return nil, errors.UnknownCreateAddressError(createErr.Error())
	}

	return &req.ID, nil
}
func (uc *AddressUseCase) SearchForAddressOnGoogleByLatitudeAndLongitude(ctx context.Context, lat float64, lng float64) (*models_responses.GeoCodeResponse, error) {
	var address models_responses.GeoCodeResponse

	_, configErr := configs.LoadConfig()

	if configErr != nil {
		fmt.Println("Erro ao carregar as configurações:", configErr)
		return nil, configErr
	}
	latStr := strconv.FormatFloat(lat, 'f', 6, 64)
	lngStr := strconv.FormatFloat(lng, 'f', 6, 64)

	headers := map[string]string{"Content-Type": "application/json"}

	response, requestExternalApiErr := uc.HttpClientUseCase.Get(ctx, "https://maps.googleapis.com/maps/api/geocode/json?latlng="+latStr+","+lngStr+"&key=AIzaSyDQL9W4ReDELEv-TluATH6G05sr0-1Dj5w", headers)

	if requestExternalApiErr != nil {
		return nil, errors.UnknownFindPersonError(requestExternalApiErr.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, readBodyErr := ioutil.ReadAll(response.Body)
	if readBodyErr != nil {
		return nil, errors.UnknownFindPersonError(readBodyErr.Error())

	}

	if deserializeErr := json.Unmarshal(body, &address); deserializeErr != nil {
		return nil, errors.UnknownFindPersonError(deserializeErr.Error())
	}
	return &address, nil
}
func (uc *AddressUseCase) FindAddressByUserId(userId uuid.UUID) (*entities.Address, error) {
	var data entities.Address

	findErr := uc.Repo.Query().First(&data, "UserId", userId).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundFindAddressError()
		}

		return nil, errors.UnknownFindAddressError(findErr.Error())
	}

	return &data, nil
}
func (uc *AddressUseCase) FindAddressByAppointmentId(appointmentId uuid.UUID) (*entities.Address, error) {
	var data entities.Address

	findErr := uc.Repo.Query().First(&data, "AppointmentId", appointmentId).Error

	if findErr != nil {
		if err.Is(findErr, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundFindAddressError()
		}

		return nil, errors.UnknownFindAddressError(findErr.Error())
	}

	return &data, nil
}
func (uc *AddressUseCase) UpdateAddress(userId uuid.UUID, request models_requests_puts.UpdateAddressRequest) error {

	Address, findErr := uc.FindAddressByUserId(userId)

	if findErr != nil {
		return findErr
	}

	Address.Update(request)

	updateErr := uc.Repo.Update(Address)

	if updateErr != nil {
		return errors.UnknownUpdateAddressError(updateErr.Error())
	}

	return nil
}
func (uc *AddressUseCase) DeleteAddress(userId uuid.UUID) error {

	Address, findErr := uc.FindAddressByUserId(userId)

	if findErr != nil {
		return findErr
	}

	removeErr := uc.Repo.Remove(Address)

	if removeErr != nil {
		return errors.UnknownDeleteAddressError(removeErr.Error())
	}

	return nil
}
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	latDistance := toRadians(lat2 - lat1)
	lonDistance := toRadians(lon2 - lon1)
	a := math.Sin(latDistance/2)*math.Sin(latDistance/2) +
		math.Cos(toRadians(lat1))*math.Cos(toRadians(lat2))*
			math.Sin(lonDistance/2)*math.Sin(lonDistance/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
func toRadians(degree float64) float64 {
	return degree * (math.Pi / 180)
}

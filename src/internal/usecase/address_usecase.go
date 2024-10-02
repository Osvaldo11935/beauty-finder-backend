package usecase

import (
	"math"
	models_requests_posts "src/internal/delivery/http/models/requests/posts"
	models_requests_puts "src/internal/delivery/http/models/requests/put"
	"src/internal/domain/entities"
	"src/internal/domain/errors"
	"src/internal/domain/interfaces_repositories"

	"github.com/google/uuid"
)

type AddressUseCase struct {
	Repo interfaces_repositories.IAddressRepository
}

func(uc *AddressUseCase) InsertAddress(userId uuid.UUID, request models_requests_posts.CreateAddressRequest) (*uuid.UUID, error){
	 req := entities.NewAddress(userId, request)

	 createErr := uc.Repo.Insert(req)

	 if createErr != nil {
		return nil, errors.UnknownCreateAddressError(createErr.Error())
	 }

	 return &req.ID, nil
}

func(uc *AddressUseCase) FindAddressByUserId(userId uuid.UUID) (*entities.Address, error){
	var data entities.Address

	findErr := uc.Repo.Query().First(&data, "UserId", userId).Error

	if findErr != nil {
	   return nil, errors.UnknownFindAddressError(findErr.Error())
	}

	return &data, nil
}

func(uc *AddressUseCase) UpdateAddress(userId uuid.UUID, request models_requests_puts.UpdateAddressRequest) (error){
	
	Address, findErr := uc.FindAddressByUserId(userId)

	if findErr !=nil {
		return findErr
	}

	Address.Update(request)
	
	updateErr := uc.Repo.Update(Address)

	if updateErr != nil {
	   return errors.UnknownUpdateAddressError(updateErr.Error())
	}

	return nil
}

func(uc *AddressUseCase) DeleteAddress(userId uuid.UUID) error{
	
	Address, findErr := uc.FindAddressByUserId(userId)

	if findErr !=nil {
		return findErr
	}

    removeErr := uc.Repo.Remove(Address)

	if removeErr !=nil {
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
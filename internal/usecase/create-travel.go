package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type CreateTravelUseCase struct {
	TravelRepository database.TravelRepository
}

func NewCreateTravelUseCase(
	TravelRepository database.TravelRepository,
) *CreateTravelUseCase {
	return &CreateTravelUseCase{
		TravelRepository: TravelRepository,
	}
}

func (c *CreateTravelUseCase) Execute(input *dto.TravelInputDTO) error {
	entity := entity.Travel{
		Price: 								 	input.Price,
  	AccountID: 							input.AccountID,
  	BusID: 								 	input.BusID,
  	DepartureTime: 					input.DepartureTime,
  	DepartureCityID: 				input.DepartureCityID,
  	ArrivalTime: 						input.ArrivalTime,
  	ArrivalCityID: 					input.ArrivalCityID,
	}

	err := c.TravelRepository.Save(&entity)
	if err != nil {
		return err
	}

	return nil
}

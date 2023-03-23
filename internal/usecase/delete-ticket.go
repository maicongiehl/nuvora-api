package usecase

import (
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type DeleteTicketUseCase struct {
	TicketRepository database.TicketRepository
}

func NewDeleteTicketUseCase(
	TicketRepository database.TicketRepository,
) *DeleteTicketUseCase {
	return &DeleteTicketUseCase{
		TicketRepository: TicketRepository,
	}
}

func (c *DeleteTicketUseCase) Execute(id int) error {
	entity := entity.Ticket{
		ID:								id,
	}

	err := c.TicketRepository.Delete(&entity)
	if err != nil {
		return err
	}

	return nil
}

package vehicle

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type TruckApp struct {
	topic      string
	repository IVehicleRepository
	publisher  *kafka.Publisher
}

func NewVehicleApp(topicName string, db *sqlx.DB, publisher *kafka.Publisher) *TruckApp {
	return &TruckApp{
		topic:      topicName,
		repository: NewVehicleRepository(db),
		publisher:  publisher,
	}
}

func (t *TruckApp) RegisterVehicle(ctx *fiber.Ctx) error {
	return nil
}

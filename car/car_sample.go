package car

import (
	"log"
	"my-protobuf/protogen/car"

	"github.com/google/uuid"
)

func ValidateCar() {
	car := car.Car{
		CarId: uuid.New().String(),
		Brand: "BMW",
		Model: "745i Sport",
		// Brand:           "Chevrolet",
		// Model:           "745i Sport xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Price:           204000,
		ManufactureYear: 2024,
	}

	err := car.ValidateAll()

	if err != nil {
		log.Fatalln("Validation failed", err)
	}

	log.Println(&car)
}

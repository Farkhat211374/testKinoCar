package routes

import (
	"context"
	"github.com/Farkhat211374/api-gateway/pkg/car/pb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Car struct {
	ID           primitive.ObjectID `json:"_id,omitempty"`
	Brand        string             `json:"brand,omitempty"`
	Model        string             `json:"model,omitempty"`
	Year         int32              `json:"year,omitempty"`
	Body         string             `json:"body,omitempty"`
	BrakeSystem  string             `json:"brake_system,omitempty"`
	Aspiration   string             `json:"aspiration,omitempty"`
	Horsepower   float32            `json:"horsepower,omitempty"`
	Mpg          float32            `json:"mpg,omitempty"`
	Cylinders    int32              `json:"cylinders,omitempty"`
	Acceleration float32            `json:"acceleration,omitempty"`
	Displacement float32            `json:"displacement,omitempty"`
	Country      string             `json:"country,omitempty"`
}

type CreateCarRequest struct {
	Car Car `json:"car"`
}

func CreateCar(ctx *gin.Context, c pb.CarServiceClient) {
	body := CreateCarRequest{}

	car := &pb.Car{
		Brand:        "Toyota",
		Model:        "Camry",
		Year:         2023,
		Body:         "Sedan",
		BrakeSystem:  "Disc brakes",
		Aspiration:   "Naturally aspirated",
		Horsepower:   350.0,
		Mpg:          32.0,
		Cylinders:    4,
		Acceleration: 6.1,
		Displacement: 2.5,
		Country:      "Japan",
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateCar(context.Background(), &pb.CreateCarRequest{
		Car: car,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

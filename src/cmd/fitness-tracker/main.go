package main

import (
	"fmt"
	"time"

	"github.com/aliot-bot/fitness-tracker/internal/models"
)

func main() {
	training := models.Training{
		ID:       1,
		Date:     time.Now(),
		Type:     "Бег",
		Duration: 30,
		Notes:    "Тестовая тренировкаaaaaa",
	}
	fmt.Println(training)
}

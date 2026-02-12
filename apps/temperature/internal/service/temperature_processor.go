package service

import (
	"fmt"
	"time"
	"math/rand"

	"github.com/google/uuid"

	"temperature/internal/entity"
)

type RandomTemperatureProcessor struct {
	
}


func (p *RandomTemperatureProcessor) TemperatureByLocation(location string) (*entity.TempreatureStats, error) {
	if location == "unknown" {
		return nil, nil
	}
	if location == "illegal" {
		return nil, fmt.Errorf("illegal location")
	}
	result := randomStats()
	result.Location = location

	return result, nil
}

func (p *RandomTemperatureProcessor) TemperatureBySensorID(id string) (*entity.TempreatureStats, error) {
	if id == "unknown" {
		return nil, nil
	}
	if id == "illegal" {
		return nil, fmt.Errorf("illegal location")
	}
	result := randomStats()
	result.SensorID = id

	return result, nil
}

func randomStats() *entity.TempreatureStats {
	return &entity.TempreatureStats{
		Description: "temperature stats",
		Location: "location",
		SensorType: "temperature",
		SensorID: uuid.NewString(),
		Timestamp: time.Now(),
		Status: "OK",
		Unit: "°C",
		Value: randomNumber(-40.0, 50.0),

	}
}

func randomNumber(min, max float64) float64 {
	return min + rand.Float64() * (max - min)
}

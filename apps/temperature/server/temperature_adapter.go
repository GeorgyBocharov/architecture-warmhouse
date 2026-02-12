package server

import (
	"fmt"
	"temperature/internal/entity"
	"encoding/json"

	"net/http"
)

type (
	TemperatureAdapter struct {
		service TemperatureService
	}
	TemperatureService interface {
		TemperatureByLocation(location string) (*entity.TempreatureStats, error)
		TemperatureBySensorID(id string) (*entity.TempreatureStats, error)
	}
)

func NewTemperatureAdapter(service TemperatureService) *TemperatureAdapter {
	return &TemperatureAdapter {
		service: service,
	}
}

func (t *TemperatureAdapter) TemperatureByLocation(w http.ResponseWriter, r *http.Request, params TemperatureByLocationParams) {
	location := params.Location

	result, err := t.service.TemperatureByLocation(location)
	if err == nil && result == nil {
		t.handleErrorResponseJSON(w, fmt.Errorf("sensor not found"), 1, http.StatusNotFound, fmt.Sprintf("no sensors in location %s", location))
	} else if err != nil {
		t.handleErrorResponseJSON(w, err, 2, http.StatusInternalServerError, "unknown error")
	}
	t.handleSucces(w, result)
}


func (t *TemperatureAdapter) TemperatureBySensor(w http.ResponseWriter, r *http.Request, sensor Sensor) {
	result, err := t.service.TemperatureBySensorID(sensor)
	if err == nil && result == nil {
		t.handleErrorResponseJSON(w, fmt.Errorf("sensor not found"), 1, http.StatusNotFound, fmt.Sprintf("no sensors by id %s", sensor))
	} else if err != nil {
		t.handleErrorResponseJSON(w, err, 2, http.StatusInternalServerError, "unknown error")
	}
	t.handleSucces(w, result)
}

func (t *TemperatureAdapter) handleSucces(w http.ResponseWriter, result *entity.TempreatureStats) {
	response := TemperatureResponse {
		Description: result.Description,
		Location: result.Location,
		SensorId: result.SensorID,
		SensorType: result.SensorType,
		Status: TemperatureResponseStatus(result.Status),
		Timestamp: result.Timestamp,
		Unit: result.Unit,
		Value: result.Value,
	}
	body, err := json.Marshal(response)
	statusCode := http.StatusOK
	if err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

func (t *TemperatureAdapter) handleErrorResponseJSON(w http.ResponseWriter, err error, errorCode, statusCode int, message string) {
  details := fmt.Sprintf("err: %v", err)
  resp := ErrorResponse{
	Code: errorCode,
	Message: message,
    Details: &details,
  }
  body, err := json.Marshal(resp)
  if err != nil {
    http.Error(w, http.StatusText(statusCode), statusCode)

    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)
  _, _ = w.Write(body)
}
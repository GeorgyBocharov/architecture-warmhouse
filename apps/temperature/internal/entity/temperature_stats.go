package entity

import "time"

type TempreatureStats struct {
	Value       float64 
	Unit        string    
	Timestamp   time.Time 
	Location    string    
	Status      string    
	SensorID    string   
	SensorType  string   
	Description string   
}
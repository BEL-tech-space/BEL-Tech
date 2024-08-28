package pkg

import (
	"time"

	"github.com/google/uuid"
)

 type WeightedWaste struct {
	TimeStamp time.Time	
	BinID uuid.UUID
	WasteReports []WasteReport

	// User ID of the user who reported the waste
	
	
 }

 type WasteReport struct {
    WasteType WasteType
    Weight    float64 // Weight in kilograms
}
 type WasteType string

 const (
	Glass WasteType = "glass"
	Metal WasteType = "metal" 	
	Plastics WasteType = "plastics"
	Paper WasteType = "paper"
 )

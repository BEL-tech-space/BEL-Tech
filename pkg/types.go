package pkg

import (
	"time"

	"github.com/google/uuid"
)

 type WeightedWaste struct {
	TimeStamp time.Time	
	BinID uuid.UUID
	UserID string
	GlassWeighed float64
	MetalWeighed float64
	PlasticsWeighed float64
	PaperWeighed float64
 }


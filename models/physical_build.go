package models

import (
	"gorm.io/gorm"
)

type PhysicalBuild struct {
	gorm.Model
    Weight     float32
    Height     float32
    BicepsSize float32
    HipsSize   float32
    WaistSize  float32
    ChestSize  float32
}

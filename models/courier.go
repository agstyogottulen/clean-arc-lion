package models

import "github.com/jinzhu/gorm"

type Courier struct {
	gorm.Model
	Role          string `json:"role"`
	FullName      string `json:"full_name"`
	PhoneNumber   string `json:"phone_number"`
	VehicleType   string `json:"vehicle_type"`
	VehiclePlate  string `json:"vehicle_plate"`
	PartnerPlate  string `json:"partner_plate"`
	PartnerName   string `json:"partner_name"`
	CourierType   string `json:"courier_type"`
	CourierStatus string `json:"courier_status"`
}

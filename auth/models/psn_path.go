// To indicate a user could see which buses
package models

import "gorm.io/gorm"

type Passenger_Path struct {
	gorm.Model
	//Auto ID
	FUserID string `json:"f_user_id"`
	FPathID string `json:"f_path_id"`
}

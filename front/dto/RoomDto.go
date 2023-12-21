package dto

import "Cgo/front/models"

type RoomDto struct {
	models.RoomApply
	//是否同意
	IsArgument bool
}

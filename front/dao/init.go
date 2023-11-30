package dao

import (
	"gorm.io/gorm"
	"reflect"
)

func DaoErrorHandle[T any](callback func() (*gorm.DB, T)) T {
	if r, result := callback(); r.RowsAffected > 0 {
		return result
	} else {
		return reflect.Zero(reflect.TypeOf(result)).Interface().(T)
	}
}

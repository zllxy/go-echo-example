package db

import "gorm.io/gorm"

type Driver interface {
	ParseDSN() string
	Dialector() gorm.Dialector
}

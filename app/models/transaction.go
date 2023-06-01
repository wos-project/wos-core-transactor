package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Uid       string `gorm:"column:uid; index"`
	Kind      string `gorm:"column:kind"`
	Status    string `gorm:"column:status"`
	TxHashHex string `gorm:"column:tx_hash_hex"`
	Cost      string `gorm:"column:cost"`
}

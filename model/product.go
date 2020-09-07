package model

import (
	"go-gin-demo/database"
)

type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Intro     string    `json:"Intro"`
}

var db = database.DB

func (p Product) Insert() (int64, error) {
	res := db.Create(&p)

	id := p.ID

	if res.Error != nil {
		return id, res.Error
	}

	return id, nil
}

//func (p *Product) List() (products []Product, err error) {
//	if err := db.Find(&products).Error; err != nil {
//		return
//	}
//
//	return
//}

//func (p Product) Update() (updatedProduct Product, err error) {
//
//}

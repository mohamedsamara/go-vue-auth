package handlers

import "gorm.io/gorm"

type BaseHandler struct {
	db *gorm.DB
}

func New(db *gorm.DB) BaseHandler {
	return BaseHandler{db}
}

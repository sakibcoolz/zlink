package controller

import (
	"zlink/log"
	"zlink/service"

	"github.com/go-playground/validator/v10"
)

type Controller struct {
	log      *log.Log
	validate *validator.Validate
	service  service.IService
}

func NewController(log *log.Log, service service.IService, validate *validator.Validate) *Controller {
	return &Controller{
		log:      log,
		validate: validate,
		service:  service,
	}
}

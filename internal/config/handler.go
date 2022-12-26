package config

import (
	handler "final-project/internal/handler"
	usecase "final-project/internal/usecase"
)

type Handler struct {
	AllHandler handler.AllHandler
}

func InitHandler(Usecase usecase.AllUsecase) Handler {
	return Handler{
		AllHandler: handler.NewHandler(Usecase),
	}
}

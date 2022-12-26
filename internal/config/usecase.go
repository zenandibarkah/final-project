package config

import (
	repository "final-project/internal/repository"
	usecase "final-project/internal/usecase"
)

type Usecase struct {
	AllUsecase usecase.AllUsecase
}

func InitUsecase(Repository repository.AllRepository) Usecase {
	return Usecase{
		AllUsecase: usecase.NewUsecase(Repository),
	}
}

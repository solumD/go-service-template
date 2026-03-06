package handler

type handler struct {
	entityUsecase EntityUsecase
}

func New(uc EntityUsecase) *handler {
	return &handler{
		entityUsecase: uc,
	}
}

package repository

type (
	Repository interface {
		Hello() string
	}
	repository struct{}
)

func NewRepository() Repository {
	return &repository{}
}

func (t *repository) Hello() string {
	return "hello from repository"
}

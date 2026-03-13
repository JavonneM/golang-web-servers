package repository

type Repository interface {
	GetSomething() error
}	

func NewRepository() Repository {
	return &RepositoryImpl{}
}

type RepositoryImpl struct {
}

func (r  *RepositoryImpl) GetSomething() error {
	return nil
}

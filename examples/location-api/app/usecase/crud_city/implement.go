package crud_city_usecase

import (
	"errors"
	"fmt"

	city_repository "github.com/herryg91/cdd/examples/location-api/app/repository/city"
	"github.com/herryg91/cdd/examples/location-api/entity"
)

type usecase struct {
	city_repo city_repository.Repository
}

func New(city_repo city_repository.Repository) UseCase {
	return &usecase{
		city_repo: city_repo,
	}
}
func (uc *usecase) GetByPrimaryKey(id int) (*entity.City, error) {
	data, err := uc.city_repo.GetById(id)
	if err != nil {
		if errors.Is(err, city_repository.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) GetAll() ([]*entity.City, error) {
	data, err := uc.city_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Create(in entity.City) (*entity.City, error) {
	data, err := uc.city_repo.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Update(in entity.City) (*entity.City, error) {
	data, err := uc.city_repo.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Delete(id int) error {
	err := uc.city_repo.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return nil
}

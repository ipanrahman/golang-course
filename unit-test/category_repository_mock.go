package unittest

import (
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	mock mock.Mock
}

func (r *CategoryRepositoryMock) FindByID(id string) *Category {
	calls := r.mock.Called(id)
	if calls.Get(0) == nil {
		return nil
	}
	category := calls.Get(0).(Category)
	return &category
}

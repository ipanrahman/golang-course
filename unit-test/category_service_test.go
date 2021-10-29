package unittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository = &CategoryRepositoryMock{mock: mock.Mock{}}
	categoryService    = CategoryService{Repository: categoryRepository}
)

func TestCategoryService_FindByIDNotFound(t *testing.T) {
	categoryRepository.mock.On("FindByID", "1").Return(nil)

	category, err := categoryService.FindByID("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_FindByIDFound(t *testing.T) {
	category := Category{
		ID:   "1",
		Name: "Elektronik",
	}
	categoryRepository.mock.On("FindByID", "2").Return(category)

	result, err := categoryService.FindByID("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, category.ID, category.ID)
	assert.Equal(t, category.Name, category.Name)
}

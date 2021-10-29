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

func TestCategoryService_FindByID(t *testing.T) {
	categoryRepository.mock.On("FindByID", "1").Return(nil)

	category, err := categoryService.FindByID("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

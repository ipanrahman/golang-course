package unittest

type CategoryRepository interface {
	FindByID(id string) *Category
}

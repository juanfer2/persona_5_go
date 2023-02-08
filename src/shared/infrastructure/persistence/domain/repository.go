package persistence_domain

import "gorm.io/gorm"

type Repository[T any] interface {
	FindBy(id int) T
	Create(newRecord T) T
	All() []T
	Delete(record T) T
	Where(query any, arg ...any) *gorm.DB
	Preload(relations ...string) *gorm.DB
	WhereBy(query any, arg ...any) T
	// FindWithAssociation(associations []string, query any, arg ...any) T
	// WhereWithAssociation(associations []string, query any, arg ...any) []T
	Join(relations ...string) *gorm.DB
	CreateInBatches(records []T)
}

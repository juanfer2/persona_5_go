package postgres

import (
	"log"

	persistence_domain "github.com/juanfer2/persona_5/src/shared/infrastructure/persistence/domain"
	"gorm.io/gorm"
)

type PostgresRepository[T any] struct {
	PostgresClient *PostgresClient
}

func NewPostgresRepository[T any](
	postgresClient *PostgresClient,
) persistence_domain.Repository[T] {
	return &PostgresRepository[T]{
		PostgresClient: postgresClient,
	}
}

func (pr *PostgresRepository[T]) Create(newRecord T) T {
	pr.PostgresClient.Conn().Create(&newRecord)

	return newRecord
}

func (pr *PostgresRepository[T]) FindBy(id int) T {
	var record T
	pr.PostgresClient.Conn().First(&record, id)

	return record
}

func (pr *PostgresRepository[T]) All() []T {
	var records []T
	pr.PostgresClient.Conn().Find(&records)

	return records
}

func (pr *PostgresRepository[T]) Delete(record T) T {
	pr.PostgresClient.Conn().Delete(&record)

	return record
}

func (pr *PostgresRepository[T]) Save(record T) T {
	pr.PostgresClient.Conn().Save(&record)

	return record
}

func (pr *PostgresRepository[T]) Update(record T, updateRecord T) T {
	pr.PostgresClient.Conn().Model(&record).Updates(updateRecord)

	return record
}

func (pr *PostgresRepository[T]) Where(query any, arg ...any) *gorm.DB {
	var table T
	return pr.PostgresClient.Conn().Model(&table).Where(query, arg)
}

func (pr *PostgresRepository[T]) WhereBy(query any, arg ...any) T {
	var table T
	pr.PostgresClient.Conn().Where(query, arg).First(&table)

	return table
}

func (pr *PostgresRepository[T]) Not(query any, arg ...any) []T {
	var table []T
	pr.PostgresClient.Conn().Not(query, arg).Find(&table)

	return table
}

func (pr *PostgresRepository[T]) CreateInBatches(data []T) {
	pr.PostgresClient.Conn().CreateInBatches(data, 100)
}

func (pr *PostgresRepository[T]) Join(relations ...string) *gorm.DB {
	var table T
	model := pr.PostgresClient.Conn().Model(&table)

	for _, relationName := range relations {
		joinRelation()(model, relationName)
	}

	if model.Error != nil {
		log.Fatalln(model.Error)
	}

	return model
}

type JoinRelation = func(g *gorm.DB, name string) *gorm.DB

func joinRelation() JoinRelation {
	return func(db *gorm.DB, name string) *gorm.DB {
		return db.Joins(name)
	}
}

func (pr *PostgresRepository[T]) Preload(relations ...string) *gorm.DB {
	var table T
	model := pr.PostgresClient.Conn().Model(&table)

	for _, relationName := range relations {
		preloadRelation()(model, relationName)
	}

	if model.Error != nil {
		log.Fatalln(model.Error)
	}

	return model
}

type RelationPreload = func(g *gorm.DB, name string) *gorm.DB

func preloadRelation() RelationPreload {
	return func(db *gorm.DB, name string) *gorm.DB {
		return db.Preload(name)
	}
}

func (pr *PostgresRepository[T]) FindWithAssociation(
	associations []string, query any, arg ...any,
) T {
	var record T
	var table T
	model := pr.PostgresClient.Conn().Model(&table).Where(query, arg)

	for _, association := range associations {
		associationAppend()(model, association)
	}

	model.First(&record)

	return record
}

func (pr *PostgresRepository[T]) WhereWithAssociation(
	associations []string, query any, arg ...any,
) []T {
	var records []T
	var table T
	model := pr.PostgresClient.Conn().Model(&table).Where(query, arg)

	for _, association := range associations {
		associationAppend()(model, association)
	}

	model.Find(&records)

	return records
}

type AssociationFunction = func(g *gorm.DB, name string) *gorm.Association

func associationAppend() AssociationFunction {
	return func(db *gorm.DB, name string) *gorm.Association {
		return db.Association(name)
	}
}

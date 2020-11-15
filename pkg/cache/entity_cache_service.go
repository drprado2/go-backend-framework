package cache

import (
	"github.com/drprado2/go-backend-framework/pkg/entities"
	"reflect"
)

type EntityCacheServiceInterface interface {
	GetOrAdd(id entities.ID) (*entities.Entity, error)
	GetOrNil(id entities.ID) (*entities.Entity, error)
	CacheAll(entityType reflect.Type) error
	Add(entity *entities.Entity) error
	AddByID(id entities.ID) error
	Delete(id entities.ID) error
	DeleteAll(entityType reflect.Type) error
}

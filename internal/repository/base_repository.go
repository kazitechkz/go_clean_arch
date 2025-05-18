package repository

import (
	"fmt"
	"football_licence/internal/filter"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Get(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) GetAll(orderBy string, direction string) ([]T, error) {
	var entities []T
	query := r.db.Model(new(T))
	if orderBy != "" {
		query = query.Order(fmt.Sprintf("%s %s", orderBy, direction))
	}
	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepository[T]) GetWithFilters(filters map[string]interface{}, orderBy string, direction string) ([]T, error) {
	var entities []T
	query := r.db.Model(new(T)).Where(filters)
	if orderBy != "" {
		query = query.Order(fmt.Sprintf("%s %s", orderBy, direction))
	}
	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepository[T]) GetFirstWithFilters(filters map[string]interface{}, orderBy, direction string) (*T, error) {
	var entity T
	query := r.db.Model(new(T)).Where(filters)

	if orderBy != "" {
		query = query.Order(fmt.Sprintf("%s %s", orderBy, direction))
	}

	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepository[T]) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(new(T)).Where("id = ?", id).Updates(updates).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	result := r.db.Delete(new(T), id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

type PaginatedResult[T any] struct {
	Items      []T
	Page       int
	PerPage    int
	TotalItems int64
	TotalPages int
}

func (r *BaseRepository[T]) Paginate(page, perPage int, filters map[string]interface{}, orderBy, direction string) (*PaginatedResult[T], error) {
	var entities []T
	var totalItems int64

	query := r.db.Model(new(T)).Where(filters)
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	if orderBy != "" {
		query = query.Order(fmt.Sprintf("%s %s", orderBy, direction))
	}

	offset := (page - 1) * perPage
	if err := query.Limit(perPage).Offset(offset).Find(&entities).Error; err != nil {
		return nil, err
	}

	totalPages := int((totalItems + int64(perPage) - 1) / int64(perPage))
	return &PaginatedResult[T]{
		Items:      entities,
		Page:       page,
		PerPage:    perPage,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}, nil
}

func (r *BaseRepository[T]) GetFirstWithDynamicFilters(filters []filter.GenericFilter, orderBy, direction string) (*T, error) {
	var entity T
	query := r.db.Model(new(T))

	for _, f := range filters {
		// безопасный SQL через ? placeholder
		switch f.Op {
		case "=", "!=", ">", "<", ">=", "<=":
			query = query.Where(fmt.Sprintf("%s %s ?", f.Field, f.Op), f.Value)
		case "ILIKE", "LIKE":
			query = query.Where(fmt.Sprintf("%s %s ?", f.Field, f.Op), f.Value)
		case "IN":
			query = query.Where(fmt.Sprintf("%s IN ?", f.Field), f.Value)
		case "BETWEEN":
			rangeVals, ok := f.Value.([2]interface{})
			if ok {
				query = query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", f.Field), rangeVals[0], rangeVals[1])
			}
		}
	}

	if orderBy != "" {
		query = query.Order(fmt.Sprintf("%s %s", orderBy, direction))
	}

	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

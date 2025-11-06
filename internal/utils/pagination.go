package utils

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"service-register/internal/models"
	"service-register/internal/validators"
)

type PaginationOptions struct {
	AllowedOrderValues []string
	MaxLimit           int
}

func ParsePagination(c *gin.Context, options *PaginationOptions) (*models.CollectionQueryOptions, error) {

	limitStr := c.Query("limit")
	pageStr := c.Query("page")
	orderBy := c.Query("orderBy")

	if orderBy != "" {
		if !validators.ValidateOrderBy(orderBy, options.AllowedOrderValues) {
			return nil, errors.New("unsupported orderBy: allowed values are: " + strings.Join(options.AllowedOrderValues, ", ") + "but got" + orderBy)
		}
	}

	if limitStr == "" || pageStr == "" {
		return nil, errors.New("limit and page are required")
	}

	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {
		return nil, errors.New("can't parse limit as int64")
	}

	if int(limit) > options.MaxLimit {
		return nil, errors.New("limit above the maximum limit allowed")
	}

	page, err := strconv.ParseUint(pageStr, 10, 32)
	if err != nil {
		return nil, errors.New("can't parse offset as int64")
	}
	if page < 1 {
		return nil, errors.New("pagination starts from 1")
	}
	return &models.CollectionQueryOptions{
		Limit:   int(limit),
		Page:    int(page),
		OrderBy: orderBy,
	}, nil
}

func CalculatePagination(currentPage, limit int, totalRecords int64) models.Pagination {
	var pagination models.Pagination

	pagination.TotalRecords = totalRecords
	pagination.CurrentPage = currentPage
	pagination.TotalPages = int(math.Ceil(float64(totalRecords) / float64(limit)))
	if pagination.CurrentPage < pagination.TotalPages {
		nextPage := pagination.CurrentPage + 1
		pagination.NextPage = &nextPage
	}
	if pagination.CurrentPage > 1 {
		previousPage := pagination.CurrentPage - 1
		pagination.PreviousPage = &previousPage
	}
	return pagination
}

func ApplyCollectionQueryOptions(query *gorm.DB, options *models.CollectionQueryOptions) (*gorm.DB, int64, error) {
	if options.OrderBy != "" {
		orderQueries := parseOrderBy(options.OrderBy)
		for _, orderBy := range orderQueries {
			query = query.Order(orderBy)
		}
	}
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if options.Limit > 0 {
		query = query.Limit(int(options.Limit))
	}
	offset := options.Limit * (options.Page - 1)
	if offset > 0 {
		query = query.Offset(int(offset))
	}
	return query, count, nil
}

func parseOrderBy(orderBy string) (orderQueries []string) {
	if orderBy == "" {
		return make([]string, 0)
	}
	orderBy = strings.ReplaceAll(orderBy, "gas", "gas_consumption")
	orderBy = strings.ReplaceAll(orderBy, "time", "submitted_at")
	orderBy = strings.ReplaceAll(orderBy, "asc", "ASC")
	orderBy = strings.ReplaceAll(orderBy, "desc", "DESC")
	orderByParts := strings.Split(orderBy, " ")
	for i, part := range orderByParts {
		orderByParts[i] = strings.ReplaceAll(part, ":", " ")
	}
	return orderByParts
}

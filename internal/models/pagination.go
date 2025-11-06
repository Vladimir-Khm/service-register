package models

type CollectionQueryOptions struct {
	Limit   int
	Page    int
	OrderBy string
}

type PageResponse[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	TotalRecords int64 `json:"totalRecords"`
	CurrentPage  int   `json:"currentPage"`
	TotalPages   int   `json:"totalPages"`
	NextPage     *int  `json:"nextPage"`
	PreviousPage *int  `json:"previousPage"`
}

type MessagesQueryFilter struct {
	CollectionQueryOptions
	UserID   *int64
	ChatID   *int64
	ThreadID *int64
}

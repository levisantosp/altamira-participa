package utils

type PaginatedResponse[T any] struct {
	Page        int  `json:"page"`
	HasNextPage bool `json:"hasNextPage"`
	Items       []T  `json:"items"`
}

type CursorPaginatedResponse[T any] struct {
	HasNextPage bool `json:"hasNextPage"`
	Items       []T  `json:"items"`
}

func PaginatedResponseFrom[T any](
	items []T,
	page int,
	limit int,
) PaginatedResponse[T] {
	hasNextPage := len(items) > limit

	if hasNextPage {
		items = items[:limit]
	}

	return PaginatedResponse[T]{
		Page:        page,
		HasNextPage: hasNextPage,
		Items:       items,
	}
}

func CursorPaginatedResponseFrom[T any](
	items []T,
	limit int,
) CursorPaginatedResponse[T] {
	hasNextPage := len(items) > limit

	if hasNextPage {
		items = items[:limit]
	}

	return CursorPaginatedResponse[T]{
		HasNextPage: hasNextPage,
		Items:       items,
	}
}

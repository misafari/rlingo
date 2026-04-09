package dto

import "github.com/google/uuid"

type Page struct {
	Limit  int
	Cursor uuid.UUID
}

func DefaultPage() Page {
	return Page{Limit: 20}
}

func (p Page) LimitOrDefault() int {
	if p.Limit <= 0 || p.Limit > 100 {
		return 20
	}
	return p.Limit
}

type PagedResult[T any] struct {
	Items      []T
	NextCursor *uuid.UUID // nil means no more pages
	Total      int
}

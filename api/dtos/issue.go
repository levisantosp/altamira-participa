package dtos

import (
	"time"

	"github.com/levisantosp/altamira-participa/api/ent/generated"
	"github.com/levisantosp/altamira-participa/api/ent/generated/issue"
)

type Issue struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      issue.Status `json:"status"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

func IssueFrom(e *generated.Issue) Issue {
	return Issue{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Status:      e.Status,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

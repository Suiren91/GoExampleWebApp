package service

import (
	"context"
	"fmt"

	"github.com/Suiren91/GoExampleWebApp/internal/entity"
	"github.com/Suiren91/GoExampleWebApp/internal/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}

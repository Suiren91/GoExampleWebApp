package store

import (
	"context"

	"github.com/Suiren91/GoExampleWebApp/internal/entity"
)

func (r *Repository) ListTasks(ctx context.Context, db Queryer) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT 
	id, title, status, created, modified 
	FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}

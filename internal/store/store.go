package store

import (
	"errors"

	"github.com/Suiren91/GoExampleWebApp/internal/entity"
)

var (
	Tasks       = &TaskStore{Tasks: map[entity.TaskID] /*本ではintだけど多分TaskIDが正しい*/ *entity.Task{}}
	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskID // 動作確認用なのでexportしてる
	Tasks  map[entity.TaskID]*entity.Task
}

// Add はTaskStoreに新しいタスクを加え，そのタスクのIDを返す
func (ts *TaskStore) Add(t *entity.Task) (int, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return int(t.ID), nil
}

// All はソート済みのタスク一覧を返す
func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}
	return tasks
}

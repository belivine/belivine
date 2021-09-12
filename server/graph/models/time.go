package models

import (
	"database/sql"
	"time"
)

type Time struct {
	ID    int64        `json:"id"`
	Start time.Time    `json:"start"`
	End   sql.NullTime `json:"end"`
	Task  Task         `json:"task"`
}

type Task struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	EstHours  sql.NullFloat64 `json:"est_hours"`
	EstDate   sql.NullTime    `json:"est_date"`
	GroupTask GroupTask       `json:"group_task"`
}

type GroupTask struct {
	ID   sql.NullInt64  `json:"id"`
	Name sql.NullString `json:"name"`
}

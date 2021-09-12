package times

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/neurocome/ine_go/graph/models"
	"github.com/neurocome/ine_go/internal/pkg/db/pgsql"
)

type InputTime struct {
	Task_id int
	User_id int
}

type UpdateTime struct {
	Time_id int
}

func (val *InputTime) Create() (models.Time, error) {
	var id int
	err := pgsql.Db.QueryRow(context.Background(), "INSERT INTO times (time_start,task_id,user_id) VALUES($1,$2,$3) returning id", time.Now().Format("2006-01-02 15:04:05"), val.Task_id, val.User_id).Scan(&id)
	if err != nil {
		return models.Time{}, err
	}
	data, err := GetDetail(id)
	if err != nil {
		return models.Time{}, err
	}
	return data, nil
}

func (val *UpdateTime) Update() (models.Time, error) {
	commandTag, err := pgsql.Db.Exec(context.Background(), "UPDATE times set time_end = $1 WHERE id = $2", time.Now().Format("2006-01-02 15:04:05"), val.Time_id)
	if err != nil {
		log.Fatal(err)
		return models.Time{}, err
	}
	if commandTag.RowsAffected() != 1 {
		return models.Time{}, errors.New("No row found to delete")
	}
	data, err := GetDetail(val.Time_id)
	if err != nil {
		return models.Time{}, err
	}
	return data, nil
}

func GetDetail(id int) (models.Time, error) {
	var value models.Time
	err := pgsql.Db.QueryRow(context.Background(), `
		SELECT 
		times.id as time_id,
		times.time_start,
		times.time_end,
		tasks.id tasks_id,
		tasks.name tasks_name,
		tasks.est_hours,
		tasks.est_date,
		task_group.id group_id,
		task_group.name group_name
		from times
		LEFT JOIN tasks ON times.task_id = tasks.id
		LEFT JOIN task_group ON tasks.task_group_id = task_group.id
		WHERE times.id = $1
	`, id).Scan(&value.ID, &value.Start, &value.End, &value.Task.ID, &value.Task.Name, &value.Task.EstHours, &value.Task.EstDate, &value.Task.GroupTask.ID, &value.Task.GroupTask.Name)
	if err != nil {
		return value, err
	}
	return value, nil
}

func GetAll(id int64) ([]*models.Time, error) {
	var value []*models.Time
	rows, err := pgsql.Db.Query(context.Background(), `
		SELECT 
		times.id as time_id,
		times.time_start,
		times.time_end,
		tasks.id tasks_id,
		tasks.name tasks_name,
		tasks.est_hours,
		tasks.est_date,
		task_group.id group_id,
		task_group.name group_name
		from times
		LEFT JOIN tasks ON times.task_id = tasks.id
		LEFT JOIN task_group ON tasks.task_group_id = task_group.id
		WHERE times.user_id = $1
	`, id)
	if err != nil {
		return value, err
	}

	for rows.Next() {
		var row models.Time
		err = rows.Scan(&row.ID, &row.Start, &row.End, &row.Task.ID, &row.Task.Name, &row.Task.EstHours, &row.Task.EstDate, &row.Task.GroupTask.ID, &row.Task.GroupTask.Name)
		if err != nil {
			return []*models.Time{}, err
		}
		value = append(value, &row)
	}
	return value, nil
}

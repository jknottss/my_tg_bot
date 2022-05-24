package database

const (
	taskCount = `
		SELECT COUNT (*)
		FROM tasks 
		WHERE user_id=$1
	`

	addTask = `
		INSERT INTO tasks (
		user_id,
		task,
		priority,
		done )
		VALUES ($1, $2, $3 + 1, $4)
	`
	doneTask = `
		UPDATE tasks
		SET done = true
		WHERE user_id = $1
		AND priority = $2
	`
	getAll = `
		SELECT priority,
		task,
		done
		FROM tasks
		WHERE user_id= $1
	`
)

type Model struct {
	Priority int    `db:"priority"`
	Task     string `db:"task"`
	Done     bool   `db:"done"`
}

func (c *Connection) AddTask(userId string, task string) (int, error) {
	conn, err := c.Pool.Acquire()
	if err != nil {
		return 0, err
	}
	defer c.Pool.Release(conn)
	var count int
	err = conn.QueryRow(taskCount, userId).Scan(&count)
	if _, err := conn.Exec(addTask, userId, task, count, false); err != nil {
		return 0, err
	}
	return count + 1, nil
}

func (c *Connection) DoneTask(userId string, taskNbr int) error {
	conn, err := c.Pool.Acquire()
	if err != nil {
		return err
	}
	defer c.Pool.Release(conn)
	if _, err := conn.Exec(doneTask, userId, taskNbr); err != nil {
		return err
	}
	return nil
}

func (c *Connection) GetAllTasks(userId string) ([]Model, error) {
	conn, err := c.Pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer c.Pool.Release(conn)
	var count, i int
	err = conn.QueryRow("SELECT COUNT (*) FROM tasks WHERE user_id=$1", userId).Scan(&count)
	modArr := make([]Model, count)
	query, err := conn.Query(getAll, userId)
	if err != nil {
		return nil, err
	}
	for query.Next() {
		err := query.Scan(&modArr[i].Priority, &modArr[i].Task, &modArr[i].Done)
		i++
		if err != nil {
			return nil, err
		}
	}
	return modArr, nil
}

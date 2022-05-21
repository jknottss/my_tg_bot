package database

const (
	addTask = `
		INSERT INTO tasks (
		user_id,
		task,
		priority,
		done )
		VALUES ($1, $2, (SELECT COUNT (*) FROM tasks WHERE user_id=$1) + 1, $3)
	`
	doneTask = `
		UPDATE tasks
		SET done = true
		WHERE user_id = $1
		AND priority = $2
	`
	getAll = `
		SELECT priority,
		task
		FROM tasks
		WHERE user_id= $1
	`
)

type Model struct {
	Priority int    `db:"priority"`
	Task     string `db:"task"`
}

func (c *Connection) AddTask(userId string, task string) error {
	conn, err := c.Pool.Acquire()
	if err != nil {
		return err
	}
	defer c.Pool.Release(conn)
	if _, err := conn.Exec(addTask, userId, task, false); err != nil {
		return err
	}
	return nil
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
		err := query.Scan(&modArr[i].Priority, &modArr[i].Task)
		i++
		if err != nil {
			return nil, err
		}
	}
	return modArr, nil
}

package database

const (
	addTask = `
		INSERT INTO tasks
		(user_id, task_nbr, task)
		VALUES ($1, $2, $3)
	`
	dropTask = `
		DELETE FROM tasks
		WHERE user_id='$1'
		AND task_nbr='$2'
	`
	showAll = `
		SELECT *
		FROM tasks
		WHERE user_id='$1'
	`
)

func (c *Connection) AddTask(userId, taskNbr int, task string) error {
	conn, err := c.Pool.Acquire()
	if err != nil {
		return err
	}
	defer c.Pool.Release(conn)
	if _, err := conn.Exec(addTask, userId, taskNbr, task); err != nil {
		return err
	}
	return nil
}

func (c *Connection) DropTask(userId, taskNbr int) error {
	conn, err := c.Pool.Acquire()
	if err != nil {
		return err
	}
	defer c.Pool.Release(conn)
	if _, err := conn.Exec(dropTask, userId, taskNbr); err != nil {
		return err
	}
	return nil
}

func (c *Connection) ShowAllTasks(userId int) (string, error) {
	conn, err := c.Pool.Acquire()
	if err != nil {
		return "", err
	}
	defer c.Pool.Release(conn)
	var task string
	if err := conn.QueryRow(showAll, userId).Scan(&task); err != nil {
		return "", err
	}
	return task, nil
}

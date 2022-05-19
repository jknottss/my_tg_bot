package database

const (
	addTask = `
		INSERT INTO tasks
		(user_id, task)
		VALUES ($1, $2)
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

func (c *Connection) AddTask(userId, task string) error {
	conn, err := c.Pool.Acquire()
	if err != nil {
		return err
	}
	defer c.Pool.Release(conn)
	if _, err := conn.Exec(addTask, userId, task); err != nil {
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
	var tasks string
	if err := conn.QueryRow(showAll, userId).Scan(&tasks); err != nil {
		return "", err
	}
	return tasks, nil
}

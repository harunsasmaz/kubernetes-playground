package storage

var (
	CreateTable = `
	CREATE TABLE IF NOT EXISTS todos (
		id VARCHAR(40) UNIQUE NOT NULL,
		title VARCHAR(50) NOT NULL,
		description TEXT NOT NULL,
		status VARCHAR(20) NOT NULL,
		created_at TIME NOT NULL DEFAULT NOW(),
		due_date TIME NOT NULL
	)`

	GetAll = `
	SELECT
	id,
	title,
	description,
	status,
	created_at,
	due_date 
	FROM todos`

	GetRemaining = `
	SELECT
	id,
	title,
	description,
	status,
	created_at,
	due_date 
	FROM todos
	WHERE status = 'waiting'`

	GetDone = `
	SELECT
	id,
	title,
	description,
	status,
	created_at,
	due_date 
	FROM todos
	WHERE status = 'done'
	`

	Get = `
	SELECT 
	id,
	title,
	description,
	status,
	created_at,
	due_date  
	FROM todos 
	WHERE id = $1`

	Insert = `
	INSERT INTO todos
	VALUES ($1, $2, $3, $4, $5, $6)`

	UpdateStatus = `
	UPDATE todos
	SET status = $2
	WHERE id = $1
	`

	Delete = `
	DELETE FROM todos
	WHERE id = $1`
)

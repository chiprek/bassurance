package store

import (
	"database/sql"
	"fmt"
)

// CreateProcessWithSteps creates a new process with steps
func CreateProcessWithSteps(db *sql.DB, p Process, steps []Step) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	processQuery := `
	INSERT INTO process (name, description, created_by)
	VALUES (?, ?, ?)
	RETURNING id;
	`
	var processID int
	err = tx.QueryRow(processQuery, p.Name, p.Description, p.CreatedBy).Scan(&processID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert process: %w", err)
	}

	stepQuery := `
	INSERT INTO step (process_id, name, description, required, critical, step_order)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	for _, s := range steps {
		_, err := tx.Exec(stepQuery, processID, s.Name, s.Description, s.Required, s.Critical, s.Order)
		if err != nil {
			return 0, fmt.Errorf("failed to insert step '%s' (order %d): %w", s.Name, s.Order, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to comit transaction: %w", err)
	}

	return processID, nil
}

// fetch a complete blueprint and all its required steps.
func GetProcess(db *sql.DB, id int) (Process, error) {
	var p Process
	//fetch parent process
	processQuery := `
	SELECT id, name, description, created_at, created_by
	FROM process
	WHERE id = ?;
	`
	// scan returned sql colums directly into process structs fields
	err := db.QueryRow(processQuery, id).Scan(
		&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.CreatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			//catch if someone entered an id that does not exist.
			return p, fmt.Errorf("process ID %d not found", err)
		}
		return p, fmt.Errorf("failed to fetch process: %w", err)
	}

	stepsQuery := `
	SELECT id, process_id, name, description, required, critical, step_order
	FROM step
	WHERE process_id = ?
	ORDER BY step_order ASC;
	`
	rows, err := db.Query(stepsQuery, id)
	if err != nil {
		return p, fmt.Errorf("failed to query for process %d: %w", id, err)
	}
	defer rows.Close()

	//fetch child steps associated to process`
	// order by how admin built process sequence
	for rows.Next() {
		var s Step
		err := rows.Scan(
			&s.ID, &s.ProcessID, &s.Name, &s.Description, &s.Required, &s.Critical, &s.Order,
		)
		if err != nil {
			return p, fmt.Errorf("failed to scan step data: %w", err)
		}

		p.Steps = append(p.Steps, s)
	}
	if err := rows.Err(); err != nil {
		return p, fmt.Errorf("error iterating over steps: %w", err)
	}
	//return the asembled blueprint
	return p, nil
}

// GetAllProcesses returns a list of all processes
func GetAllProcesses(db *sql.DB) ([]Process, error) {
	query := `
	SELECT id, name, description, created_at, created_by
	FROM process
	ORDER BY name ASC;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all processes: %w", err)
	}
	defer rows.Close()

	var processes []Process

	for rows.Next() {
		var p Process
		err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.CreatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan process row: %w", err)
		}
		processes = append(processes, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over process rows: %w", err)
	}
	return processes, nil
}

// addStep insterts a single step into an existing process
func AddStep(db *sql.DB, s Step) error {
	query := `
	INSERT INTO step (process_id, name, description, required, critical, step_order)
	VALUES (?,?,?,?,?,?);
	`

	_, err := db.Exec(query, s.ProcessID, s.Name, s.Description, s.Required, s.Critical, s.Order)
	if err != nil {
		return fmt.Errorf("failed to insert step: %w", err)
	}
	return nil
}

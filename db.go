package main

import (
	"database/sql"
	"fmt" // Imported for formatted error messages

	_ "modernc.org/sqlite"
)

func initDB() (*sql.DB, error) {
	DB, err := sql.Open("sqlite", "bassurance.sqlite?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Renamed to 'queries' to prevent shadowing the 'sql' package
	queries := []string{
		// 1. The Foundation
		`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        name TEXT NOT NULL,
        role TEXT NOT NULL CHECK(role IN('admin', 'operator')));`,

		// 2. The Blueprints
		`CREATE TABLE IF NOT EXISTS process (
  		id INTEGER PRIMARY KEY,
   		name TEXT NOT NULL,
    	description TEXT NOT NULL,
  		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  		created_by INTEGER NOT NULL,
        FOREIGN KEY(created_by) REFERENCES users(id));`,

		`CREATE TABLE IF NOT EXISTS step (
        id INTEGER PRIMARY KEY,
        process_id INTEGER NOT NULL,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        required INTEGER NOT NULL CHECK(required IN(0, 1)),
        critical INTEGER NOT NULL CHECK(critical IN(0, 1)),
        step_order INTEGER NOT NULL,
        FOREIGN KEY(process_id) REFERENCES process(id));`,

		// 3. The Active Floor
		`CREATE TABLE IF NOT EXISTS work_orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        process_id INTEGER NOT NULL,
        identifier TEXT NOT NULL,
        status TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        created_by INTEGER NOT NULL,
        FOREIGN KEY(process_id) REFERENCES process(id),
        FOREIGN KEY(created_by) REFERENCES users(id));`,

		// 4. The Paper Trail
		`CREATE TABLE IF NOT EXISTS completion(
        id INTEGER PRIMARY KEY,
        process_id INTEGER NOT NULL,
        step_id INTEGER NOT NULL,
        work_order_id INTEGER NOT NULL,
        completed_by INTEGER NOT NULL,
        completed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        notes TEXT,
        FOREIGN KEY(process_id) REFERENCES process(id),
        FOREIGN KEY(step_id) REFERENCES step(id),
        FOREIGN KEY(work_order_id) REFERENCES work_orders(id),
        FOREIGN KEY(completed_by) REFERENCES users(id));`,

		`CREATE TABLE IF NOT EXISTS qaphoto(
        id INTEGER PRIMARY KEY,
        completion_id INTEGER NOT NULL,
        file_path TEXT NOT NULL,
        uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        uploaded_by INTEGER NOT NULL,
        FOREIGN KEY(completion_id) REFERENCES completion(id),
        FOREIGN KEY(uploaded_by) REFERENCES users(id));`,

		`CREATE TABLE step_field (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        step_id INTEGER NOT NULL,
        prompt TEXT NOT NULL,
        field_type TEXT NOT NULL,
        target_val REAL,
        tolerance REAL,
        field_order INTERGER NOT NULL,
        FOREIGN KEY(step_id) REFERENCES step(id)
        );`,
	}

	// The loop that actually executes the SQL strings
	for _, q := range queries {
		if _, err := DB.Exec(q); err != nil {
			return nil, fmt.Errorf("failed to execute query %q: %w", q, err)
		}
	}

	return DB, nil
}

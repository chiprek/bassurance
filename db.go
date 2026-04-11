//so i have a predicament here is one of my sql table and have forgotten how to do multi line strings
// "CREATE TABLE IF NOT EXISTS step (
//       id INTEGER PRIMARY KEY,
//       process_id INTEGER NOT NULL,
//       name TEXT NOT NULL,
//       description TEXT NOT NULL,
//       required INTEGER NOT NULL,
//       critical INTEGER NOT NULL,
//       step_order INTEGER NOT NULL,
//       FOREIGN KEY(process_id) REFERENCES process(id));"

package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func initDB() (*sql.DB, error) {
	DB, err := sql.Open("sqlite", "bassurance.sqlite?_pragma=foreign_keys(1)")
	sql := []string{
		`CREATE TABLE IF NOT EXISTS users (
  		id INT PRIMARY KEY,
   		name TEXT NOT NULL,
    	description TEXT NOT NULL,
  		created_at TEXT NOT NULL,
  		created_by TEXT NOT NULL);`,

		`CREATE TABLE IF NOT EXISTS step (
        id INTEGER PRIMARY KEY,
        process_id INTEGER NOT NULL,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        required INTEGER NOT NULL,
        critical INTEGER NOT NULL,
        step_order INTEGER NOT NULL,
        FOREIGN KEY(process_id) REFERENCES process(id));`,

        `CREATE TABLE IF NOT EXISTS completion(
        id INTEGER PRIMARY KEY,
        process_id INTEGER NOT NULL,
        step_id INTEGER NOT NULL,
        completed_by INTEGER NOT NULL,
        completed_at TEXT NOT NULL,
        notes TEXT NOT NULL,
        FOREIGN KEY(process_id) REFRENCES process(id),
        FOREIGN KEY(completed_by) REFRENCES user(id));`,

        `CREATE TABLE IF NOT EXISTS user (
        id INTEGER PRIMARY KEY,
        username TEXT UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        name TEXT NOT NULL,
        role TEXT NOT NULL CHECK(role IN('admin', 'operator')));`,

        `CREATE TABLE IF NOT EXISTS workorder()`,

        `CREATE TABLE IF NOT EXISTS qaphoto(
        id INTEGER PRIMARY KEY,
        completion_id INTEGER NOT NULL,
        `
	}
	if err != nil {

	}

	return DB, nil
}

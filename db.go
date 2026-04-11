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
)

func initDB() (*sql.DB, error) {
	DB, err := sql.Open("sqlite", "bassurance.sqlite")
	sql := []string{
		`create table users (
  		id INT not null,
   		name TEXT not null,
    	description TEXT not null,
  		created_at TEXT not null,
  		created_by TEXT not null,
  		primary key (id))`,

		`CREATE TABLE IF NOT EXISTS step (
        id INTEGER PRIMARY KEY,
        process_id INTEGER NOT NULL,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        required INTEGER NOT NULL,
        critical INTEGER NOT NULL,
        step_order INTEGER NOT NULL,
        FOREIGN KEY(process_id) REFERENCES process(id));`,
	}
	if err != nil {

	}

	return DB, nil
}

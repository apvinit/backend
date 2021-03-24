package main

import (
	"database/sql"
	"log"
)

func initDB(db *sql.DB) {

	// posts
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts(
			id INTEGER PRIMARY KEY,
			short_link TEXT,
			image_link TEXT,
			type TEXT,
			title TEXT,
			name TEXT,
			info TEXT,
			created_date TEXT,
			updated_date TEXT,
			organisation TEXT,
			total_vacancy TEXT,
			age_limit_as_on TEXT,
			draft BOOLEAN,
			trash BOOLEAN DEFAULT false
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Important dates
	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS dates(
			id INTEGER PRIMARY KEY,
			date TEXT,
			title TEXT,
			post_id INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Links
	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS links(
			id INTEGER PRIMARY KEY,
			title TEXT,
			url TEXT,
			post_id INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Fees
	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS fees(
			id INTEGER PRIMARY KEY,
			title TEXT,
			body TEXT,
			post_id INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Items
	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS items(
			id INTEGER PRIMARY KEY,
			title TEXT,
			body TEXT,
			post_id INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Vacancies
	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS vacancies(
			id INTEGER PRIMARY KEY,
			category TEXT,
			name TEXT,
			gen TEXT,
			obc TEXT,
			bca TEXT,
			bcb TEXT,
			ews TEXT,
			sc TEXT,
			st TEXT,
			ph TEXT,
			total TEXT,
			age_limit TEXT,
			eligibility TEXT,
			post_id INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// posts search
	_, err = db.Exec(`
	  CREATE VIRTUAL TABLE IF NOT EXISTS posts_search USING FTS5(
			id, title, name, info, organisation)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

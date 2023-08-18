package handler

import "database/sql"

type Deps struct {
	DB *sql.DB
}

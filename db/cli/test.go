package main

import "gitlab.com/otis-team/backend/db/migrations/tables"

func main () {
	migrations.CreateUsersTableIfNotExists()
}
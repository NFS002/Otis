package main

import "gitlab.com/otis_team/backend/db/migrations/tables"

func main () {
	migrations.CreateUsersTableIfNotExists()
}

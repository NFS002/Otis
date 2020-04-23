package main

import (
	migrations "gitlab.com/otis_team/backend/db/migrations/tables"
)

func main () {
	migrations.CreateUsersTableIfNotExists()
}

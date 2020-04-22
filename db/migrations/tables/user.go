package migrations

import (
	//"gitlab.com/otis_team/backend/db/client"
	"io/ioutil"
	"log"
	"strings"
)

// CreateUsersTableIfNotExists function creates the 'Users' table in the DB
func CreateUsersTableIfNotExists() error  {
	const PATH = "/Users/noah/Otis/backend/dtypes/user/schema/users.sql"
	//cli := client.RDSClient{}
	//cliErr := cli.Init()
	var cliErr error
	if cliErr != nil {
		log.Fatal(cliErr)
		return cliErr
	}
	allsql, sqlErr := ioutil.ReadFile(PATH)
	if sqlErr != nil {
		log.Fatal(sqlErr)
		return sqlErr
	}

	statements := strings.Split(string(allsql),";\n\n")
	for i, s := range statements {
		log.Printf("Executing statement %d: '%.20s...' ...", i + 1, s)
		//err := cli.DB.Exec(s)
		//if err != nil {
		//	log.Printf("Error executing statement: %v", err)
		//}
	}
	return nil
}

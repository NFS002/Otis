package main

import (
	"fmt"
	"gitlab.com/otis-team/backend/db/client"
)

func main() {
	cli := client.RDSClient{}
	cli.Init()
	fmt.Println(cli)
	cli.DB.Close()
}

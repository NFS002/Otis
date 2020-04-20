package main

import (
	"errors"
	"flag"
	"gitlab.com/otis-team/backend/db/client"
	"gitlab.com/otis-team/backend/db/migrations/tables"
	"log"
	"strings"
)

const (
	allTables = "users,transactions,merchants,plans,wallets,outlets," +
		"tags,merchant_tags,accounts,groups,members,friends" +
		"countries,admins,contributions,universities"
)

// CreateTable : Utility function as part of a CLI program to automate and execute various db migrations
func CreateTable(t string) {
	cli := client.RDSClient{}
	cli.DB.
		t = strings.ToLower( t )
	log.Printf("Creating table: %s",t)
	var err error
	switch t {
	case "users":
		err = migrations.CreateUsersTableIfNotExists()
	case "groups":
		err = migrations.CreateGroupsTableIfNotExists()
	case "admins":
		err = migrations.CreateAdminsTableIfNotExists()
	case "transactions":
		err = migrations.CreateTransactionsTableIfNotExists()
	case "merchants":
		err = migrations.CreateMerchantsTableIfNotExists()
	case "friends":
		err = migrations.CreateFriendsTableIfNotExists()
	case "members":
		err = migrations.CreateMembersTableIfNotExists()
	case "contributions":
		err = migrations.CreateContributionsTableIfNotExists()
	case "wallets":
		err = migrations.CreateWalletsTableIfNotExists()
	case "accounts":
		err = migrations.CreateAccountsTableIfNotExists()
	case "countries":
		err = migrations.CreateCountriesTableIfNotExists()
	case "outlets":
		err = migrations.CreateOutletsTableIfNotExists()
	case "tags":
		err = migrations.CreateTagsTableIfNotExists()
	case "merchant_tags":
		err = migrations.CreateMerchantTagsTableIfNotExists()
	case "plans":
		err = migrations.CreatePlansTableIfNotExists()
	case "universities":
		err = migrations.CreateUniversitiesTableIfNotExists()
	case "ALL":
		err = errors.New("the 'ALL' argument is only valid when specified by itself")
	default:
		err = errors.New("Table name " + t + " could not be resolved")
	}

	if err == nil {
		log.Printf("Table: %s created succesfully",t)
	} else {
		log.Printf("Error: %v",err)
	}
}

func main() {
	inp := *flag.String("create","","Comma separated list of the names of the tables to create. 'ALL' creates all tables")
	if inp == "" {
		log.Println("No tables specified.")
	} else {
		if inp == "ALL" {
			inp = allTables
		}
		tables := strings.Split(inp,",")
		for _, t := range tables {
			CreateTable(t)
		}
		cli := client.RDSClient{}
		cli.Init()
		log.Println(cli.DB.Raw("\\list"))
		cli.DB.Close()
	}
}

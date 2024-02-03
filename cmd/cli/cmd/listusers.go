/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hculpan/kabbase/pkg/dbbadger"
	"github.com/hculpan/kabbase/pkg/entities"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// listusersCmd represents the listusers command
var listusersCmd = &cobra.Command{
	Use:   "listusers",
	Short: "List authorized users",
	Long:  `List authorized users`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("too many arguments given")
		}
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		dbPath := os.Getenv("DB_PATH")
		if len(dbPath) == 0 {
			dbPath = "./data"
		}

		db, err := dbbadger.OpenDB(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		kvMap, err := dbbadger.GetKeyValuesWithPrefix(db, "user_")
		if err != nil {
			return err
		}

		for k, v := range kvMap {
			user := entities.User{}
			err := json.Unmarshal(v, &user)
			if err != nil {
				return err
			}
			fmt.Printf("%s: Username: %s, Last Login: %s \n", k, user.Username, user.LastLogin.Format(time.RFC822))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listusersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listusersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listusersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

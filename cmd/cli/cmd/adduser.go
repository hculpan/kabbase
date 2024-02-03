/*
Copyright © 2024 Harry Culpan <harry@culpan.org>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/hculpan/kabbase/pkg/dbbadger"
	"github.com/hculpan/kabbase/pkg/entities"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// adduserCmd represents the adduser command
var adduserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "Add an authorized user",
	Long:  `Add an authorized user. This user should have a username and a secret key.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("must specify username and passkey")
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

		_, err = dbbadger.FetchKey(db, "user_"+args[0])
		if err != nil && !errors.Is(err, badger.ErrKeyNotFound) {
			log.Fatal(err)
		} else if err != nil && errors.Is(err, badger.ErrKeyNotFound) {
			user := entities.NewUser(args[0], args[1])
			jsonData, err := json.Marshal(user)
			if err != nil {
				return err
			}
			if err := dbbadger.SetKey(db, "user_"+args[0], []byte(jsonData)); err != nil {
				return err
			}
			fmt.Printf("User %q successfully added\n", args[0])
		} else {
			fmt.Printf("User %q already registered\n", args[0])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(adduserCmd)
	adduserCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		return errors.New("adduser [username] [key]")
	})

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adduserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adduserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

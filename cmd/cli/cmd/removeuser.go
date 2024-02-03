/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/hculpan/kabbase/pkg/dbbadger"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// removeuserCmd represents the removeuser command
var removeuserCmd = &cobra.Command{
	Use:   "removeuser",
	Short: "Remove an authorized user",
	Long:  `Remove an authorized user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("must specify username")
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
		if err != nil && errors.Is(err, badger.ErrKeyNotFound) {
			fmt.Printf("User %q not found\n", args[0])
		} else {
			err := dbbadger.DeleteKey(db, "user_"+args[0])
			if err != nil {
				return err
			}
			fmt.Printf("User %q successfully removed\n", args[0])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

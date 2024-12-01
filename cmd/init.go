/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"batkor/restapp/kernel"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cobra"
	"strings"
)

func dropTable(bundleName string) {
	queryStr := fmt.Sprintf("DROP TABLE %s", bundleName)
	_, err := kernel.Database().Query(queryStr)

	if err != nil {
		panic(err)
	}
}

func exitTable(table string) bool {
	var result bool
	query := fmt.Sprintf("SELECT 1 FROM pg_tables WHERE schemaname = '%s' and tablename='%s'",
		kernel.GetSettings().Database.Schema,
		table,
	)
	err := kernel.Database().QueryRow(query).Scan(&result)

	if errors.Is(err, sql.ErrNoRows) {
		return result
	}

	if err != nil {
		panic(err)
	}

	return result
}

func createBundleTable(bundleName string) {
	bundle, ok := kernel.GetSettings().Bundles[bundleName]

	if !ok {
		log.Fatalf("Not found bundle \"%s\".", bundleName)
	}

	queryCreateTableStr := fmt.Sprintf("CREATE TABLE %s.%s (\n", kernel.GetSettings().Database.Schema, bundleName)
	var queryColumnsList []string
	var queryIndexColumnsList []string

	for fieldName, fieldObj := range bundle.Fields {
		if fieldObj.OwnTable {
			continue
		}

		queryColumnsList = append(queryColumnsList, fieldObj.QueryTableColumn(fieldName))

		if fieldObj.Index {
			queryIndexColumnsList = append(queryIndexColumnsList, fieldName)
		}
	}

	queryCreateTableStr += strings.Join(queryColumnsList, ",\n")
	queryCreateTableStr += ")"
	fmt.Println(queryCreateTableStr)
	fmt.Println(bundle)
	_, err := kernel.Database().Query(queryCreateTableStr)

	if err != nil {
		panic(err)
	}

	if len(queryIndexColumnsList) > 0 {
		_, err := kernel.
			Database().
			Query(fmt.Sprintf("CREATE INDEX idx_%s ON %s.%s (%s)",
				bundleName,
				kernel.GetSettings().Database.Schema,
				bundleName,
				strings.Join(queryIndexColumnsList, ","),
			))

		if err != nil {
			panic(err)
		}
	}
}

// initCmd represents the create command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize new project",
	Long:  `Create tables needs for projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")

		for bundle, _ := range kernel.GetSettings().Bundles {
			tableExist := exitTable(bundle)

			if tableExist {
				if force {
					dropTable(bundle)
				} else {
					log.Fatalf("Already exists table for bundle \"%s\". Add flag -f for force recreate table.", bundle)
				}
			}

			createBundleTable(bundle)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolP("force", "f", false, "Force recreate tables")
}

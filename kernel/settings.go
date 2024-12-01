package kernel

import (
	"gopkg.in/yaml.v3"
	"os"
)

var instance *Settings

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Schema   string
}

type field struct {
	Type     string
	OwnTable bool
	Primary  bool
	Null     bool
	Unique   bool
	Index    bool
}

// QueryTableColumn Returns query string for create table column.
func (field field) QueryTableColumn(name string) string {
	columnStr := "\t" + name + "\t" + field.Type

	if field.Primary {
		columnStr += " PRIMARY KEY"
	} else {
		if field.Unique {
			columnStr += " UNIQUE KEY"
		}

		if !field.Null {
			columnStr += " NOT NULL"
		}
	}

	return columnStr
}

type bundles struct {
	Fields map[string]field
}

type Settings struct {
	Address  string
	Database database
	Bundles  map[string]bundles
}

func GetSettings() *Settings {
	if instance != nil {
		return instance
	}

	file, err := os.Open("settings.yml")

	if err != nil {
		panic(err)
	}

	rawSettings := yaml.NewDecoder(file)
	err = rawSettings.Decode(&instance)

	if err != nil {
		panic(err)
	}

	if instance.Database.Schema == "" {
		instance.Database.Schema = "public"
	}

	return instance
}

package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB = nil

func Connect() {
	if DB == nil {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/blog")
		if err != nil {
			panic(err)
		}
		DB = db
	} else {
		if err := DB.Ping(); err != nil {
			panic(err)
		}
	}
}

func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			panic(err)
		}
		DB = nil
	}
}

func Init() {
	tables := []any{User{}, Blog{}, BlogImages{}, BlogLikes{}, BlogComments{}, Error{}}

	Connect()
	defer Close()

	for _, table := range tables {
		createTable(table)
	}
}

func createTable(model any) {
	t := reflect.TypeOf(model)
	isModelPtr := reflect.TypeOf(model).Kind() == reflect.Ptr

	structName := ""
	if isModelPtr {
		structName = t.Elem().Name()
	} else {
		structName = t.Name()
	}

	firstChar := string(structName[0])
	lowerCaseFirstChar := strings.ToLower(firstChar)
	structName = strings.Replace(structName, firstChar, lowerCaseFirstChar, 1)

	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (", structName)

	var fields reflect.Value

	if isModelPtr {
		fields = reflect.ValueOf(model).Elem()
	} else {
		fields = reflect.ValueOf(model)
	}
	numFields := fields.NumField()

	for i := 0; i < numFields; i++ {
		name := fields.Type().Field(i).Name
		field, found := t.FieldByName(name)

		if !found {
			continue
		}

		query := field.Tag.Get("sql")
		json := field.Tag.Get("json")

		if i == numFields-1 {
			sql += fmt.Sprintf("%v %v)", json, query)
		} else {
			sql += fmt.Sprintf("%v %v,", json, query)
		}
	}

	fmt.Println(sql)
	if _, err := DB.Exec(sql); err != nil {
		panic(err)
	}
}

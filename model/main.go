package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
	return db
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func Init() {
	tables := []any{User{}, Blog{}, BlogImages{}, BlogLikes{}, BlogComments{}, Error{}}

	db := Connect()
	defer Close(db)

	for _, table := range tables {
		createTable(db, table)
	}
}

func createTable(db *sql.DB, model any) {
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
	numFiled := fields.NumField()

	for i := 0; i < numFiled; i++ {
		name := fields.Type().Field(i).Name
		field, found := t.FieldByName(name)

		if !found {
			continue
		}

		sql := field.Tag.Get("sql")
		json := field.Tag.Get("json")

		if i == numFiled-1 {
			sql += fmt.Sprintf("%v %v)", json, sql)
		} else {
			sql += fmt.Sprintf("%v %v,", json, sql)
		}
	}

	if _, err := db.Exec(sql); err != nil {
		panic(err);
	}
}

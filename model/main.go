package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

var DB *sql.DB = nil

func Connect() {
	if DB == nil {
		db, err := sql.Open("postgres", "")
		if err != nil {
			panic(err)
		}
		DB = db
	}
}

func Close()  {
	if DB != nil {
		if err := DB.Close(); err != nil {
			panic(err)
		}
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

	if _, err := DB.Exec(sql); err != nil {
		panic(err);
	}
}

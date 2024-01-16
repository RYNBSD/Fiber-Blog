package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

var DB *sql.DB = nil

func Connect() {
	if DB == nil {
		db, err := sql.Open("postgres", "postgres:password@tcp(127.0.0.1:5432)/blog")
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

	wg := sync.WaitGroup{}

	for _, table := range tables {
		wg.Add(1)
		go func(table any) {
			defer wg.Done()
			createTable(table)
		}(table)
	}
	wg.Wait()
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
		panic(err)
	}
}

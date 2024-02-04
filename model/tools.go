package model

import (
	"blog/types"
	"blog/util"
	"database/sql"
	"os"
	"path"
	"reflect"
	"sync"
)

func createImages(blogId string, images ...string) {
	if images != nil {
		wg := sync.WaitGroup{}

		const sql = `DELETE FROM "blogImages" WHERE "blogId"=$1`
		if _, err := DB.Exec(sql, blogId); err != nil {
			panic(err)
		}

		for _, image := range images {
			wg.Add(1)

			go func(image string) {
				defer wg.Done()
				const sql = `INSERT INTO "blogImages" ("image", "blogId") VALUES ($1, $2)`

				if _, err := DB.Exec(sql, image, blogId); err != nil {
					panic(err)
				}
			}(image)
		}
		wg.Wait()
	}
}

func deleteImages(images ...string) {
	if images != nil {
		wg := sync.WaitGroup{}

		for _, image := range images {
			wg.Add(1)

			go func(image string) {
				defer wg.Done()
				imagePath := path.Join(util.PublicDir(), image)

				if err := os.Remove(imagePath); err != nil {
					panic(err)
				}
			}(image)
		}
		wg.Wait()
	}
}

func scanUnknownColumns(rows *sql.Rows, elems *[]types.Map) {
	columns, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
	}

	for rows.Next() {

		values := make([]any, len(columns))
		elem := types.Map{}

		for i, column := range columns {
			elem[column.Name()] = reflect.New(column.ScanType()).Interface()
			values[i] = elem[column.Name()]
		}

		if err := rows.Scan(values...); err != nil {
			panic(err)
		}

		*elems = append(*elems, elem)
	}
}

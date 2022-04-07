package lib

import (
	"database/sql"
	"log"
)

type GenericRowset struct {
	cols, colPtrs []interface{}
	colNames      []string
	rows          *sql.Rows
}

func InitRowset(r *sql.Rows) (*GenericRowset, error) {
	var err error
	result := GenericRowset{
		rows: r,
	}

	// reading columns names
	result.colNames, err = r.Columns()

	if err != nil {
		return nil, err
	}
	result.cols = make([]interface{}, len(result.colNames))
	result.colPtrs = make([]interface{}, len(result.colNames))

	for i := 0; i < len(result.colNames); i++ {
		result.colPtrs[i] = &result.cols[i]
	}

	return &result, nil
}

func (r *GenericRowset) ReadRowAsString() map[string]interface{} {
	var myMap = make(map[string]interface{})
	var err error

	err = r.rows.Scan(r.colPtrs...)
	if err != nil {
		log.Fatal(err)
	}
	for i, col := range r.cols {
		s, ok := col.([]byte)
		if ok {
			myMap[r.colNames[i]] = string(s)
		} else {
			myMap[r.colNames[i]] = col
		}

	}

	// // Do something with the map
	// for key, val := range myMap {
	// 	fmt.Println("Key:", key, "Value Type:", reflect.TypeOf(val))
	// }

	return myMap

}

func (r *GenericRowset) ReadRow() map[string]interface{} {
	var myMap = make(map[string]interface{})
	var err error

	err = r.rows.Scan(r.colPtrs...)
	if err != nil {
		log.Fatal(err)
	}
	for i, col := range r.cols {
		myMap[r.colNames[i]] = col
	}

	// // Do something with the map
	// for key, val := range myMap {
	// 	fmt.Println("Key:", key, "Value Type:", reflect.TypeOf(val))
	// }

	return myMap

}

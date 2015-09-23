package DBFuncs

import (
    "database/sql"
    _"github.com/lib/pq"
    "fmt"
    "strconv"
)


type Tuple struct {
    Col string
    Value interface{}
}

type Rows struct {
    Tuples []Tuple 
}

var db *sql.DB

func OpenDBConnection() {
    tempdb, err := sql.Open("postgres", "user=postgres password=lol dbname=servermonitor")
	if err != nil {
		fmt.Println("Unable to conenct to the db! ", err)
	}
    db = tempdb
}



func InsertIntoTable(tableName string, rows Rows) {
    query := "INSERT INTO " + tableName + "("
    valuesString := "VALUES("
    values := []interface{}{}
    //building our query
    for i,v := range rows.Tuples { //foreach
        query += v.Col + ",";
        valuesString +="$" + strconv.Itoa((i+1)) + ","
        values = append(values, v.Value)
    }
    fmt.Println("valueString " + valuesString)
    query = replaceAtIndex(query, ')', len(query)-1)
    valuesString = replaceAtIndex(valuesString, ')', len(valuesString)-1)
    query += " " +valuesString;
    fmt.Println(query)
   
    stmt, err := db.Prepare(query) 
    defer stmt.Close()
    _, err = stmt.Query(values...)
    if err != nil {
        fmt.Println("everything is dead")		
	}

} 

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    fmt.Println("String out " + string(out))
    return string(out)
}



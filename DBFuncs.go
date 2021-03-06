package DBFuncs

import (
    "database/sql"
    _"github.com/lib/pq"
    "fmt"
   // "strconv"
)



//these structs are not inused!
//dont delete, may need them later
type Values struct {
    A []Type
}

type Type struct {
    Value interface{}
}

//how to make a query

/*
    var values []interface{}
    values = append(values, 601)
    r := db.Query("INSERT INTO information(cpu_temp) values($1)", values ) 
    var col string
    for r.Next() { 
        r.Scan(&col)
        fmt.Println(col)
    }
*/



var db *sql.DB

func OpenDBConnection() {
    var err error
    db, err = sql.Open("postgres", "user=postgres password=lol dbname=servermonitor sslmode=disable")
	check(err)
    
}

func check(err error) {
    if err != nil {
        fmt.Println(err)		
	}
}

//////////
// A Generic function to query to database
// 
// param values - the values contains the value we are going to query
//////////   
func Query(query string, values []interface{}) (*sql.Rows) {
    var rows *sql.Rows
    var err error
    var stmt *sql.Stmt
    if(values == nil) { // no stmt
          rows, err = db.Query(query)
          check(err)      
    }else {
        stmt, err = db.Prepare(query)
        check(err)
        
        rows, err = stmt.Query(values...)
        check(err)
        defer stmt.Close()
    } 
    check(err)
    return rows
    //important to call DeferRows,when u are done, to avoid runtime panic
}
//after all queries, you have to call this after you are done with Rows
func DeferRows(rows *sql.Rows) {
   defer rows.Close()
}


//help function for replacing a char in a specific index
func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}



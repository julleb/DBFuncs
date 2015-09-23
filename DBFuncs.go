package DBFuncs

import (
    "database/sql"
    _"github.com/lib/pq"
    "fmt"
    "strconv"
)


type SQLRows struct {
   Rows sql.Rows 

}

//struct for holding column name and its value
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


//////////
// A Generic function to insert into a database table
// param tableName - name of the table
// param rows - A struct which includes the column names and values for each columnn
//////////   
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
        fmt.Println("everything is dead " , err)		
	}

} 

func SelectFromTable(tableName string, rows Rows) {
    query := "SELECT ";
      
    
}

func SelectAllFromTable(tableName string) (*sql.Rows) {
    query := "SELECT * FROM " + tableName
    rows, err := db.Query(query)
    if err != nil {
		fmt.Println(err)
	}
	
    return rows
    //it is very important to call DeferRows, when u are done, to avoid runtime panic
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



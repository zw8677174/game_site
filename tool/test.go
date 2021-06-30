package main

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "zw:Zongwen123!@#@tcp(49.232.87.103:3306)/game_site?charset=utf8")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }
    rows, err := db.Query("select id,username,password from auth where id>?", 0)
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()
    var id int
    var username string
    var password string
    for rows.Next() {
        err := rows.Scan(&id, &username,&password)
        if err != nil {
            panic(err.Error())
        }
        log.Println(id, username, password)
    }
    err = rows.Err()
    if err != nil {
        panic(err.Error())
    }
}


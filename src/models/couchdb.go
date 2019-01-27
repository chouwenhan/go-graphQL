package models

import (
    "fmt"
    "time"
    "github.com/rhinoman/couchdb-go"

    "github.com/rs/xid"
)

func Authentication() (*couchdb.Connection, couchdb.BasicAuth) {
    var timeout = time.Duration(500 * time.Millisecond)
    conn, err := couchdb.NewConnection("192.168.0.105", 5984, timeout)
    if err != nil {
        panic(err)
    }
    Bauth := couchdb.BasicAuth{Username: "admin", Password: "admin" }
    return conn, Bauth
}

func CreateDatabase(DbName string) (error){
    var timeout = time.Duration(500 * time.Millisecond)
    conn, err := couchdb.NewConnection("192.168.0.105", 5984, timeout)
    if err != nil {
        panic(err)
    }
    Bauth := couchdb.BasicAuth{Username: "admin", Password: "admin" }
    var auth couchdb.Auth = &Bauth
    err = conn.CreateDB(DbName , auth)
    fmt.Printf("CreateDB successful")
    return err
}

func ConnDB(DbName string) *couchdb.Database{
    conn, Bauth := Authentication()
    db := conn.SelectDB(DbName, &Bauth)
    return db
}

func CreateDocument(db *couchdb.Database, doc interface{}) (string){
    theId := xid.New().String()
    _, err := db.Save(doc, theId, "")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Create Document successful")
    return theId
}

func ReadDocument(db *couchdb.Database, id string) (*Article, error){
    a := Article{}
    _, err := db.Read(id, &a, nil)
    if err != nil {
        panic(err)
    }
    return &a, err
}


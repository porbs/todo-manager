package database

import (
	"gopkg.in/mgo.v2"
)

/*
Database session
*/
var Session *mgo.Session

/*
Notes model connection
*/
var Notes *mgo.Collection

/*
Init database
*/
func Init(uri, dbName string) error {
	session, err := mgo.Dial(uri)

	if err != nil {
		return err
	}

	session.SetMode(mgo.Strong, true)

	Session = session
	Notes = session.DB(dbName).C("notes")

	return nil
}
package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	databaseName string = "broccolibot"
	ipAddress    string = "127.0.0.1"
)

/*
  ========================================
  Basics (comment)
  ========================================
*/

func initMongoDB(collectionName string) (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial(ipAddress)
	if err != nil {
		log.Println(err)
	}

	return session.DB(databaseName).C(collectionName), session
}
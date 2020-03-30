package database_client

import (
	"fmt"
	"github.com/gocql/gocql")

func init()  {
	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err !=nil{
		panic(err)
	}
	fmt.Print("Connection created succesfull")
	defer session.Close()
}
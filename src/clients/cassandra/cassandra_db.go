package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to Cassandra cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	fmt.Println("ik ga nu de sessie opvragen")
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Er ging iets mis bij de sessie ophalen", err)
		return nil, err
	}
	return session, nil
}

package singleton

import (
	"fmt"
	"sync"
)

type Database struct {
	connection string
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
	once.Do(func() {
		fmt.Println("Creating singleton instance...")
		instance = &Database{
			connection: "MySQL Connection",
		}
	})
	return instance
}

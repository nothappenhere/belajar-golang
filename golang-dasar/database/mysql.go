package database

var connection string

func init() {
	connection = "MySQL"
}

func GetDB() string{
	return connection
}
package database_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/database"
	_ "github.com/go-sql-driver/mysql" //database driver
	"testing"
)

func TestDatabase(t *testing.T) {
	var tests = struct{
		queries []string
	}{
		queries: []string{
			"CREATE SCHEMA IF NOT EXISTS cs53test",
			"USE cs53test",
			"CREATE TABLE IF NOT EXISTS test(id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, title TINYTEXT NOT NULL, number INT NOT NULL, created TIMESTAMP DEFAULT CURRENT_TIMESTAMP);",
			"INSERT INTO test (title, number) VALUES (\"Hello World\", 42)",
			"INSERT INTO test (title, number) VALUES (\"Blaze It\", 420)",
			"SELECT * FROM test",
			"SELECT number FROM test WHERE title=\"Hello World\"",
			"SELECT title FROM test WHERE number=420",
			"DROP TABLE test",
			"DROP SCHEMA IF EXISTS cs53test",
		},
	}

	database.Configure(&database.MySQLInfo{
		Hostname: "localhost",
		Port: 3306,
		Username: "root",
		Password: "",
		Database: "",
		ParseTime: true,
	})

	database.Open()
	defer database.Close()

	for _, query := range tests.queries {
		_, err := database.Get().Exec(query)
		if err != nil {
			t.Logf("error: %v\nquery: %s\n", err, query)
			t.Fail()
		}
	}
}

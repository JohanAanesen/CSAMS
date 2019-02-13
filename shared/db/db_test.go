package db_test

import (
	_ "github.com/go-sql-driver/mysql" //database driver
)

/* todo remove this or add tables to database.sql
func TestDatabase(t *testing.T) {
	var tests = struct {
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

	db.ConfigureDB(&db.MySQLInfo{
		Hostname:  "localhost",
		Port:      3306,
		Username:  "root",
		Password:  "",
		Database:  "",
		ParseTime: true,
	})

	for _, query := range tests.queries {
		_, err := db.GetDB().Exec(query)
		if err != nil {
			t.Logf("error: %v\nquery: %s\n", err, query)
			t.Fail()
		}
	}
}*/

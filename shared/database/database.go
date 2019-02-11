package database

// Database struct
type Database struct {
	Info MySQLInfo
}

type MySQLInfo struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func Configure(db *MySQLInfo) {

}

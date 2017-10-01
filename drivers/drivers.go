package drivers

// Normalized drivers
var (
	Sqlite = "sqlite3"
	Mysql  = "mysql"
	Pgsql  = "postgres"
)

// Drivers is a normalized list of drivers.
var Drivers = map[string]string{
	"*sqlite3.SQLiteDriver": Sqlite,
	"*mysql.MySQLDriver":    Mysql,
	"*pq.Driver":            Pgsql,
}

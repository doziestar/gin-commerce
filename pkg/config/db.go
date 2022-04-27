package config


type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func DbConfiguration() string {
	// dbname := viper.GetString("database.dbname")
	// username := viper.GetString("database.username")
	// password := viper.GetString("database.password")
	// host := viper.GetString("database.host")
	// port := viper.GetString("database.port")
	// sslMode := viper.GetString("database.sslmode")

	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	// 	host, username, password, dbname, port, sslMode,
	// )
	dns := "mamun:123@192.168.10.160:5432/test_pg_go?charset=utf8"
	return dns
}

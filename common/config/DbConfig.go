package config

type DbConfig struct {
	engine             string
	host               string
	database           string
	port               int
	user               string
	password           string
	maxIdleConnections int
	maxOpenConnections int
	charset            string
}

func (d *DbConfig) Charset() string {
	return d.charset
}

func (d *DbConfig) SetCharset(charset string) {
	d.charset = charset
}

var dbConfig = &DbConfig{}

func GetDbConfig() *DbConfig {
	return dbConfig
}

func (d *DbConfig) MaxIdleConnections() int {
	return d.maxIdleConnections
}

func (d *DbConfig) SetMaxIdleConnections(maxIdleConnections int) {
	d.maxIdleConnections = maxIdleConnections
}

func (d *DbConfig) MaxOpenConnections() int {
	return d.maxOpenConnections
}

func (d *DbConfig) SetMaxOpenConnections(maxOpenConnections int) {
	d.maxOpenConnections = maxOpenConnections
}

func (d *DbConfig) Password() string {
	return d.password
}

func (d *DbConfig) SetPassword(password string) {
	d.password = password
}

func (d *DbConfig) User() string {
	return d.user
}

func (d *DbConfig) SetUser(user string) {
	d.user = user
}

func (d *DbConfig) Port() int {
	return d.port
}

func (d *DbConfig) SetPort(port int) {
	d.port = port
}

func (d *DbConfig) Database() string {
	return d.database
}

func (d *DbConfig) SetDatabase(database string) {
	d.database = database
}

func (d *DbConfig) Host() string {
	return d.host
}

func (d *DbConfig) SetHost(host string) {
	d.host = host
}

func (d *DbConfig) Engine() string {
	return d.engine
}

func (d *DbConfig) SetEngine(engine string) {
	d.engine = engine
}

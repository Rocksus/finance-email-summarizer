package config

type Config struct {
	App      *AppConfig                 `yaml:"app"`
	Log      *LogConfig                 `yaml:"log"`
	Database map[string]*DatabaseConfig `yaml:"database"`
}

type LogConfig struct {
	LogLevel  string `yaml:"log_level"`
	LogFormat string `yaml:"log_format"`
}

type AppConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Name            string `yaml:"name"`
	Driver          string `yaml:"driver"`
	Host            string `yaml:"-"`
	Credentials     string `yaml:"-"`
	Port            int    `yaml:"-"`
	MaxOpenConn     int    `yaml:"max_open_conn"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime_secs"`
	ConnMaxIdleTime int    `yaml:"conn_max_idle_time_secs"`
	SSLRootCert     string `yaml:"-"`
	SSLCert         string `yaml:"-"`
	SSLKey          string `yaml:"-"`
}

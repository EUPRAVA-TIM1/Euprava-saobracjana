package config

import (
	"os"
)

type Config struct {
	Port            string
	Host            string
	MysqlPort       string
	MySqlHost       string
	MySqlRootPass   string
	SSOIssuer       string
	SsoServiceHost  string
	SsoServicePort  string
	SudServiceHost  string
	SudServicePort  string
	MupServiceHost  string
	MupServicePort  string
	FileServicePort string
	FileServiceHost string
}

const (
	portKey              = "SAOBRACAJNA_SERVICE_PORT"
	hostKey              = "SAOBRACAJNA_SERVICE_HOST"
	defaultHost          = "e_uprava_saobracajna"
	defaultPort          = "8002"
	mySqlPort            = "SAOBRACAJNA_MYSQL_PORT"
	defaultMySqlPort     = "3306"
	mySqlHost            = "SAOBRACAJNA_MYSQL_HOST"
	defaultMySqlHost     = "saobracajna_mysql"
	mySqlRootPass        = "SAOBRACAJNA_MYSQL_ROOT_PASSWORD"
	defaultMySqlRootPass = "root"
	SSOIssuerKey         = "SAOBRACAJNA_ISSUER_KEY"
	DefaultSSOIssuer     = "saobracajna614393"
	SsoServicePortKey    = "SSO_SERVICE_PORT"
	SsoServiceHostKey    = "SSO_SERVICE_HOST"
	MupServiceHostKey    = "MUP_HOST"
	MupServicePortKey    = "MUP_PORT"
	SudServicePortKey    = "SUD_SERVICE_PORT"
	SudServiceHostKey    = "SUD_SERVICE_HOST"
	FileServicePortKey   = "FILE_SERVICE_PORT"
	FileServiceHostKey   = "FILE_SERVICE_HOST"
)

func NewConfig() (c Config) {
	if port, set := os.LookupEnv(portKey); set && port != "" {
		c.Port = port
	} else {
		c.Port = defaultPort
	}

	if host, set := os.LookupEnv(hostKey); set && host != "" {
		c.Host = host
	} else {
		c.Host = defaultHost
	}

	if port, set := os.LookupEnv(mySqlPort); set && port != "" {
		c.MysqlPort = port
	} else {
		c.MysqlPort = defaultMySqlPort
	}

	if host, set := os.LookupEnv(mySqlHost); set && host != "" {
		c.MySqlHost = host
	} else {
		c.MySqlHost = defaultMySqlHost
	}

	if pass, set := os.LookupEnv(mySqlRootPass); set && pass != "" {
		c.MySqlRootPass = pass
	} else {
		c.MySqlRootPass = defaultMySqlRootPass
	}

	if key, set := os.LookupEnv(SSOIssuerKey); set && key != "" {
		c.SSOIssuer = key
	} else {
		c.SSOIssuer = DefaultSSOIssuer
	}

	if port, set := os.LookupEnv(SsoServicePortKey); set && port != "" {
		c.SsoServicePort = port
	} else {
		panic("No ssoService port provided!")
	}
	if host, set := os.LookupEnv(SsoServiceHostKey); set && host != "" {
		c.SsoServiceHost = host
	} else {
		panic("No ssoService host provided!")
	}

	if port, set := os.LookupEnv(MupServicePortKey); set && port != "" {
		c.MupServicePort = port
	} else {
		panic("No mupService port provided!")
	}
	if host, set := os.LookupEnv(MupServiceHostKey); set && host != "" {
		c.MupServiceHost = host
	} else {
		panic("No mupService host provided!")
	}

	if port, set := os.LookupEnv(SudServicePortKey); set && port != "" {
		c.SudServicePort = port
	} else {
		panic("No sudService port provided!")
	}
	if host, set := os.LookupEnv(SudServiceHostKey); set && host != "" {
		c.SudServiceHost = host
	} else {
		panic("No sudService host provided!")
	}
	if host, set := os.LookupEnv(FileServiceHostKey); set && host != "" {
		c.FileServiceHost = host
	} else {
		panic("No fileService host provided!")
	}
	if port, set := os.LookupEnv(FileServicePortKey); set && port != "" {
		c.FileServicePort = port
	} else {
		panic("No fileService port provided!")
	}
	return
}

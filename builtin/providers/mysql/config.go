package mysql

import (
	"database/sql"
	"fmt"
	// MySQL driver
	_ "github.com/ziutek/mymysql/godrv"

	"github.com/hashicorp/go-multierror"
)

// Config - provider config
type Config struct {
	Endpoint string
	Username string
	Password string
}

// Client struct holding connection string
type Client struct {
	username string
	connStr  string
}

//NewClient returns new client config
func (c *Config) NewClient() (*Client, error) {
	// Connection String
	var connStr string
	var errs *multierror.Error

	connStr = fmt.Sprintf("tcp:%s*/%s/%s", c.Endpoint, c.Username, c.Password)
	if c.Endpoint != "" {
		if c.Endpoint[0] == '/' {
			connStr = fmt.Sprintf("unix:%s*/%s/%s", c.Endpoint, c.Username, c.Password)
		}
	}

	client := Client{
		username: c.Username,
		connStr:  connStr,
	}

	return &client, errs.ErrorOrNil()
}

//Connect provides database connection
func (c *Client) Connect() (*sql.DB, error) {
	conn, err := sql.Open("mymysql", c.connStr)

	if err != nil {
		return nil, fmt.Errorf("Error connecting to MySQL Server: %s", err)
	}

	return conn, nil
}

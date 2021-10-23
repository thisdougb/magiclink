// +build dev test

//
// The build tag means this file is included when the prod tag is used.
//
// eg: go run api/server.gp
//
// https://golang.org/pkg/go/build/

package config

import (
	"time"
)

const (
	API_PORT = "8080"

	DB_HOST = "127.0.0.1"
	DB_PORT = "6379"

	MAGICLINK_ID_LENGTH       = 64
	MAGICLINK_EXPIRES_MINUTES = 15

	SESSION_ID_LENGTH          = 64
	SESSION_ID_EXPIRES_MINUTES = 60

	HttpSessionTTL  = time.Hour * time.Duration(24) * 7
	HttpSessionName = "magiclink:sessionid"
)

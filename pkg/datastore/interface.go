package datastore

import (
	"github.com/thisdougb/magiclink/pkg/usecase/auth"
	"github.com/thisdougb/magiclink/pkg/usecase/send"
)

// DatastoreInterface methods are implemented by any concrete datastore
type DatastoreInterface interface {
	Connect() bool
	Disconnect()

	send.Repository
	auth.Repository
}

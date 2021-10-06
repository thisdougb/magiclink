package datastore

import (
	"github.com/thisdougb/magiclink/pkg/usecase/requestlink"
)

// DatastoreInterface methods are implemented by any concrete datastore
type DatastoreInterface interface {
	Connect() bool
	Disconnect()

	requestlink.Repository
}

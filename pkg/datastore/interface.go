package datastore

import (
	"github.com/thisdougb/magiclink/pkg/usecase/enablething"
)

// DatastoreInterface methods are implemented by any concrete datastore
type DatastoreInterface interface {
	Connect() bool
	Disconnect()

	enablething.Repository
}

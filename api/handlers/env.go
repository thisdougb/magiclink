package handlers

import (
	"github.com/thisdougb/magiclink/pkg/usecase/send"
	"log"
)

/*
   The Env struct allows us to pass the datastore pointer around,
   it also allows us to inject mocks in usecase packages.
*/

type Env struct {
	Logger      *log.Logger
	SendService *send.Service
}

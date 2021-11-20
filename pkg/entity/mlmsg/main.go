package mlmsg

import ()

type MagicLinkMsg struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewMagicLinkMsg() *MagicLinkMsg {
	return &MagicLinkMsg{}
}

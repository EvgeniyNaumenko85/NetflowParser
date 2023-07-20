package models

import (
	"errors"
	"net"
)

type NetFlowRecord struct {
	Source      net.IP
	Destination net.IP
	AccountID   []byte
	TClass      []byte
}

var ErrNoFlags = errors.New("error: должен быть использован хотябы один флаг")

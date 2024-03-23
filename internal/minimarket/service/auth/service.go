package auth

import "golang.org/x/mod/sumdb/storage"

type name interface {
}

type Service struct {
	storage storage.Storage
}

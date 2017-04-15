package grpc

import (
	"google.golang.org/grpc/naming"
)

type resolver struct {
	w naming.Watcher
}

func NewResolver(addrs ...string) naming.Resolver {
	return &resolver{
		w: NewWatcher(addrs...),
	}
}

func (r *resolver) Resolve(target string) (naming.Watcher, error) {
	return r.w, nil
}

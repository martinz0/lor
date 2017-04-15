package grpc

import (
	"errors"

	"google.golang.org/grpc/naming"
)

type watcher struct {
	called    bool
	initAddrs []*naming.Update
}

func NewWatcher(addrs ...string) naming.Watcher {
	initAddrs := make([]*naming.Update, len(addrs))
	for idx, addr := range addrs {
		initAddrs[idx] = &naming.Update{
			Op:   naming.Add,
			Addr: addr,
		}
	}
	return &watcher{
		called:    false,
		initAddrs: initAddrs,
	}
}

func (w *watcher) Next() ([]*naming.Update, error) {
	if w.called {
		return nil, errors.New("not need watch now")
	}
	w.called = true
	return w.initAddrs, nil
}

func (w *watcher) Close() {}

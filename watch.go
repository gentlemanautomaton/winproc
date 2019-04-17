// +build windows

package winproc

import (
	"context"
	"time"
)

// Watch polls the process tree on an interval until an error is encountered or the
// context is cancelled.
//
// The provided action is called each time the tree is polled. If action returns
// an error watch will stop polling and return it.
//
// This function is experimental and may be changed in future revisions.
func Watch(ctx context.Context, action TreeAction, interval time.Duration, options ...CollectionOption) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			tree, err := Tree(options...)
			if err != nil {
				return err
			}
			if err := action(tree); err != nil {
				return err
			}
		}
	}
}

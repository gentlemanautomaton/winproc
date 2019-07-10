// +build windows

package winproc

import (
	"context"
	"time"
)

// Watch polls the process tree on an interval until an error is encountered
// or the context is cancelled. It sends differences in the process list on
// the returned channel.
//
// This function is experimental and may be changed in future revisions.
func Watch(ctx context.Context, interval time.Duration, chanSize int, options ...CollectionOption) <-chan ChangeSet {
	ch := make(chan ChangeSet, chanSize)

	go func() {
		defer close(ch)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		known := make(map[UniqueID]Process)

		for {
			select {
			case <-ctx.Done():
				ch <- ChangeSet{Err: ctx.Err(), Time: time.Now()}
				return
			case <-ticker.C:
				list, err := List(options...)
				if err != nil {
					ch <- ChangeSet{Err: err, Time: time.Now()}
					return
				}

				var (
					current = make(map[UniqueID]struct{})
					changes []Change
				)
				for i := range list {
					id := list[i].UniqueID()
					current[id] = struct{}{}
					if _, exists := known[id]; exists {
						continue
					}
					known[id] = list[i]
					changes = append(changes, Change{Process: list[i]})
				}

				for id := range known {
					if _, found := current[id]; !found {
						changes = append(changes, Change{Process: known[id], Removed: true})
						delete(known, id)
					}
				}

				if len(changes) > 0 {
					ch <- ChangeSet{
						Changes: changes,
						Time:    time.Now(),
					}
				}
			}
		}
	}()

	return ch
}

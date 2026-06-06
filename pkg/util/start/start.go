package start

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Starter is an interface for types that can sync and start controllers.
type Starter interface {
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

// All syncs and then starts all starters with the given threadiness.
func All(ctx context.Context, threadiness int, starters ...Starter) error {
	if err := Sync(ctx, starters...); err != nil {
		return err
	}
	return Start(ctx, threadiness, starters...)
}

// Sync concurrently syncs all starters.
func Sync(ctx context.Context, starters ...Starter) error {
	eg, _ := errgroup.WithContext(ctx)
	for _, starter := range starters {
		func(starter Starter) {
			eg.Go(func() error {
				return starter.Sync(ctx)
			})
		}(starter)
	}
	return eg.Wait()
}

// Start sequentially starts all starters.
func Start(ctx context.Context, threadiness int, starters ...Starter) error {
	for _, starter := range starters {
		if err := starter.Start(ctx, threadiness); err != nil {
			return err
		}
	}
	return nil
}

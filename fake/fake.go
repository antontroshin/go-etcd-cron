/*
Copyright (c) 2024 Diagrid Inc.
Licensed under the MIT License.
*/

package fake

import (
	"context"

	"github.com/diagridio/go-etcd-cron/api"
)

// Fake is a fake cron instance used for testing.
type Fake struct {
	runFn  func(ctx context.Context) error
	addFn  func(ctx context.Context, name string, job *api.Job) error
	getFn  func(ctx context.Context, name string) (*api.Job, error)
	delFn  func(ctx context.Context, name string) error
	delPFn func(ctx context.Context, prefixes ...string) error
}

func New() *Fake {
	return &Fake{
		runFn: func(ctx context.Context) error {
			<-ctx.Done()
			return ctx.Err()
		},
		addFn: func(context.Context, string, *api.Job) error {
			return nil
		},
		getFn: func(context.Context, string) (*api.Job, error) {
			return nil, nil
		},
		delFn: func(context.Context, string) error {
			return nil
		},
		delPFn: func(context.Context, ...string) error {
			return nil
		},
	}
}

func (f *Fake) WithRun(fn func(context.Context) error) *Fake {
	f.runFn = fn
	return f
}

func (f *Fake) WithAdd(fn func(context.Context, string, *api.Job) error) *Fake {
	f.addFn = fn
	return f
}

func (f *Fake) WithGet(fn func(context.Context, string) (*api.Job, error)) *Fake {
	f.getFn = fn
	return f
}

func (f *Fake) WithDelete(fn func(context.Context, string) error) *Fake {
	f.delFn = fn
	return f
}

func (f *Fake) WithDeletePrefixes(fn func(context.Context, ...string) error) *Fake {
	f.delPFn = fn
	return f
}

func (f *Fake) Run(ctx context.Context) error {
	return f.runFn(ctx)
}

func (f *Fake) Add(ctx context.Context, name string, job *api.Job) error {
	return f.addFn(ctx, name, job)
}

func (f *Fake) Get(ctx context.Context, name string) (*api.Job, error) {
	return f.getFn(ctx, name)
}

func (f *Fake) Delete(ctx context.Context, name string) error {
	return f.delFn(ctx, name)
}

func (f *Fake) DeletePrefixes(ctx context.Context, prefixes ...string) error {
	return f.delPFn(ctx, prefixes...)
}
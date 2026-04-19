package r_tree

import "r-tree/internal/split"

const (
	defaultMaxEntries = 16
	defaultMinEntries = 4
)

type TreeOptions struct {
	maxEntries int
	minEntries int
	splitter   split.Splitter
}

type Option func(o *TreeOptions)

func WithMaxEntries(maxEntries int) Option {
	return func(o *TreeOptions) {
		o.maxEntries = maxEntries
	}
}

func WithMinEntries(minEntries int) Option {
	return func(o *TreeOptions) {
		o.minEntries = minEntries
	}
}

func WithSplitter(splitter split.Splitter) Option {
	return func(o *TreeOptions) {
		o.splitter = splitter
	}
}

func defaultOptions() TreeOptions {
	return TreeOptions{
		maxEntries: defaultMaxEntries,
		minEntries: defaultMinEntries,
		splitter:   nil,
	}
}

func applyOptions(opts []Option) TreeOptions {
	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

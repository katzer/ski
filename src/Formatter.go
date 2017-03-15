package main

import "io"

// IFormatter interface for formatters
type IFormatter interface {
	format(planets []Planet, opts *Opts, writer io.Writer)
	init()
}

//Formatter a factory for creating types of IFormatter implementations
type Formatter struct {
}

func (formatter *Formatter) getFormatter(opts *Opts) IFormatter {
	// Ugly
	if !opts.Pretty && len(opts.Template) == 0 {
		return nil
	}
	if !opts.Pretty && len(opts.Template) > 0 {
		realDeal := TableFormatter{}
		proxy := TFAdapter{real: &realDeal}
		return proxy
	}
	// Pretty
	if len(opts.Template) > 0 {
		realDeal := PrettyTableFormatter{}
		proxy := PTFAdapter{real: &realDeal}
		return proxy
	}
	realDeal := PrettyFormatter{}
	proxy := PFAdapter{real: &realDeal}
	return proxy
}

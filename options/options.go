/*
	This is the options module. It will update the Options Data base table if there were any changes the user made
*/

package options

import (
	Layout "github.com/alphastigma101/Coconuts-At-Wars/layout"
)

// GameMode represents the dimension type (2D or 3D)
type gameMode int
type dndMode int
type coopMode int

type Options struct {
	GameMode gameMode
	DndMode  dndMode
	CoopMode coopMode
}

// CreateRenderer factory function to create appropriate renderer
func UpdateOptions(opts *Options, table *Layout.Table) (*Options, interface{}) {
	if opts.GameMode == 1 {
		// Need to update the Options Table
		return opts, table
	} else {
		newTable, newOpts := table.Options.Init(*table, *opts)
		// Update the underlying variable
		tempTable := newTable.(Layout.Table)
		*table = tempTable
		tempOpts := newOpts.(Options)
		*opts = tempOpts
		return opts, *table
	}
}

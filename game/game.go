package game

import (
	dnd "input_handler/Dnd"

	"github.com/alphastigma101/Coconuts-At-Wars/campaign"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
	/*"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"*/)

// Layout of the game. It will copied to the database assigned with a unique player id
type Game struct {
	Options  options.Options
	Title    options.TitleScreen
	campaign campaign.Campaign
	dnd      dnd.Dnd
}

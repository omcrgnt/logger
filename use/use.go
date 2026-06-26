// Package use registers logger system defaults in unique.Global.
//
// Import for side effects at the app composition root (main or a meta use package):
//
//	import _ "github.com/omcrgnt/logger/use"
package use

import (
	"github.com/omcrgnt/logger"
	"github.com/omcrgnt/res/unique"
)

func init() {
	unique.MustAddReplaceable(logger.DefaultLog())
	unique.MustAddReplaceable(logger.DefaultStdout())
}

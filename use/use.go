// Package use registers logger system defaults in res.Default.
//
// Import for side effects at the app composition root (main or a meta use package):
//
//	import _ "github.com/omcrgnt/logger/use"
package use

import (
	"github.com/omcrgnt/logger"
	"github.com/omcrgnt/res"
)

func init() {
	_ = res.AddWithTags(logger.OutputStdout{}, res.TagReplaceable)
	_ = res.AddWithTags(logger.DefaultLog(), res.TagReplaceable)
}

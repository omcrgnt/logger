package logger

import (
	"github.com/omcrgnt/res"
)

func init() {
	out := stdoutOutput{}
	_ = res.AddBuiltin(out)
	if err := applyHandler(defaultLevelValue, defaultFormatValue, out); err != nil {
		panic(err)
	}
}

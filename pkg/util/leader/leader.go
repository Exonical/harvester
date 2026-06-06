package leader

import (
	wranglerleader "github.com/rancher/wrangler/v3/pkg/leader"
)

type Callback = wranglerleader.Callback

var RunOrDie = wranglerleader.RunOrDie

package abci

import (
	kcfg "kchain/types/cfg"
	klog "kchain/utils/log"
)

var (
	logger = klog.GetLogWithKeyVals("module", "abci")
	cfg = kcfg.GetConfig()
)

package pay

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

func init() {
	xlog.Info("GoPay Version: ", gopay.Version)
}

package async

import (
	"context"

	"github.com/sirupsen/logrus"
)

func AsyncCtx(ctx context.Context, logger *logrus.Logger, fn func(c context.Context)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(err)
			}
		}()
		fn(ctx)
	}()
}

package basictypes

import "github.com/AnimusPEXUS/utils/logger"

type PackagerI interface {
	Run(log *logger.Logger) error
}

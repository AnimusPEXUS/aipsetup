package basictypes

import "github.com/AnimusPEXUS/utils/logger"

type PrePackagerI interface {
	Run(log *logger.Logger) error
}

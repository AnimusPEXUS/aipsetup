package aipsetup

//
// import "github.com/AnimusPEXUS/utils/logger"
//
// // Currently main goal of AipSetup struct is to manage global Logger
// // and pass it to subsystems.
// // It is used by AipSetup commands, but not mening necessety to be used by
// // 3rd party software.
//
// type AipSetup struct {
// 	log  *logger.Logger
// 	root string
// }
//
// func NewAipSetup(root string, log *logger.Logger) *AipSetup {
// 	self := new(AipSetup)
// 	self.log = log
// 	self.root = root
// 	return self
// }
//
// func (self *AipSetup) NewSystem() *System {
// 	ret := NewSystem(self.root, self.log)
// 	return ret
// }
//
// func (self *AipSetup) NewSystemWithRoot(root string) *System {
// 	ret := NewSystem(root, self.log)
// 	return ret
// }
//
// func (self *AipSetup) NewBuildingSiteCtl(
// 	sys *System,
// 	path string,
// ) (*BuildingSiteCtl, error) {
//
// 	ret, err := NewBuildingSiteCtl(sys, path, self.log)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return ret, nil
// }

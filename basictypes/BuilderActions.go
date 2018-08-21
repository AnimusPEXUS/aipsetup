package basictypes

import (
	"errors"

	"github.com/AnimusPEXUS/utils/logger"
)

type BuilderActionCallable func(log *logger.Logger) error

type BuilderAction struct {
	Name     string
	Callable BuilderActionCallable
}

type BuilderActions []*BuilderAction

func NewBuilderActions() BuilderActions {
	return make(BuilderActions, 0)
}

func (self BuilderActions) Get(name string) (*BuilderAction, bool) {
	for _, i := range self {
		if i.Name == name {
			return i, true
		}
	}
	return nil, false
}

func (self BuilderActions) Replace(name string, action *BuilderAction) error {
	for k := range self {
		v := self[k]
		if v.Name == name {
			self[k] = action
			return nil
		}
	}

	return errors.New("not found")
}

func (self BuilderActions) Remove(name string) BuilderActions {

	ret := self

	for i := len(ret) - 1; i != -1; i-- {
		iv := ret[i]
		if iv.Name == name {
			ret = append(ret[:i], ret[i+1:]...)
		}

	}

	return ret
}

func (self BuilderActions) Move(index, new_index int) BuilderActions {
	ret := self

	z := ret[index]

	if new_index > index {
		new_index -= 1
	}

	ret = append(ret[:index], ret[index+1:]...)
	ret = append(ret[:new_index], append(BuilderActions{z}, ret[new_index:]...)...)

	return ret
}

func (self BuilderActions) Index(name string) int {

	for k := range self {
		v := self[k]
		if v.Name == name {
			return k
		}
	}

	return -1
}

func (self BuilderActions) MoveAfter(index int, name string) error {
	new_index := self.Index(name)
	if new_index == -1 {
		return errors.New("name not found")
	}
	self.Move(index, new_index+1)
	return nil
}

func (self BuilderActions) MoveBefore(index int, name string) error {
	new_index := self.Index(name)
	if new_index == -1 {
		return errors.New("name not found")
	}
	self.Move(index, new_index)
	return nil
}

func (self BuilderActions) MoveNamedAfter(index, name string) error {
	index_index := self.Index(index)
	if index_index == -1 {
		return errors.New("index index not found")
	}
	return self.MoveAfter(index_index, name)
}

func (self BuilderActions) MoveNamedBefore(index, name string) error {
	index_index := self.Index(index)
	if index_index == -1 {
		return errors.New("index index not found")
	}
	return self.MoveBefore(index_index, name)
}

func (self BuilderActions) AddActionsBefore(value BuilderActions, index int) BuilderActions {
	ret := self
	ret = append(ret[:index], append(append(BuilderActions{}, value...), ret[index:]...)...)
	return ret
}

func (self BuilderActions) AddActionsAfter(value BuilderActions, index int) BuilderActions {
	ret := self
	ret = append(ret[:index+1], append(append(BuilderActions{}, value...), ret[index+1:]...)...)
	return ret
}

func (self BuilderActions) AddActionsBeforeName(value BuilderActions, name string) (BuilderActions, error) {

	index := self.Index(name)
	if index == -1 {
		return nil, errors.New("index not found")
	}

	return self.AddActionsBefore(value, index), nil
}

func (self BuilderActions) AddActionsAfterName(value BuilderActions, name string) (BuilderActions, error) {

	index := self.Index(name)
	if index == -1 {
		return nil, errors.New("index not found")
	}

	return self.AddActionsAfter(value, index), nil
}

func (self BuilderActions) Append(value BuilderActions) (BuilderActions, error) {
	return append(self, value...), nil
}

func (self BuilderActions) ActionList() []string {
	ret := make([]string, 0)
	for _, i := range self {
		ret = append(ret, i.Name)
	}
	return ret
}

// Short Functions

func (self BuilderActions) ReplaceShort(name string, callable BuilderActionCallable) error {

	ret := self.Replace(
		name,
		&BuilderAction{
			Name:     name,
			Callable: callable,
		},
	)

	return ret
}

func NewBuilderActionsLenOne(name string, callable BuilderActionCallable) BuilderActions {
	ret := BuilderActions{
		&BuilderAction{
			Name:     name,
			Callable: callable,
		},
	}
	return ret
}

func (self BuilderActions) AddActionAfterNameShort(
	name string,
	new_name string, callable BuilderActionCallable,
) (BuilderActions, error) {

	ret, err := self.AddActionsAfterName(
		NewBuilderActionsLenOne(
			new_name,
			callable,
		),
		name,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self BuilderActions) AddActionBeforeNameShort(
	name string,
	new_name string, callable BuilderActionCallable,
) (BuilderActions, error) {

	ret, err := self.AddActionsBeforeName(
		NewBuilderActionsLenOne(
			new_name,
			callable,
		),
		name,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

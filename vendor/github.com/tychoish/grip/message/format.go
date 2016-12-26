package message

import (
	"fmt"

	"github.com/tychoish/grip/level"
)

type formatMessenger struct {
	base    string
	args    []interface{}
	Base    `bson:"metadata" json:"metadata" yaml:"metadata"`
	Message string `bson:"message" json:"message" yaml:"message"`
}

// NewFormattedMessage takes arguments as fmt.Sprintf(), and returns
// an object that only runs the format operation as part of the
// Resolve() method.
func NewFormattedMessage(p level.Priority, base string, args ...interface{}) Composer {
	m := &formatMessenger{
		base: base,
		args: args,
	}
	m.SetPriority(p)

	return m
}

func NewFormatted(base string, args ...interface{}) Composer {
	return &formatMessenger{
		base: base,
		args: args,
	}
}

func (f *formatMessenger) Resolve() string {
	if f.Message == "" {
		f.Message = fmt.Sprintf(f.base, f.args...)
	}

	return f.Message
}

func (f *formatMessenger) Loggable() bool {
	return f.base != ""
}

func (f *formatMessenger) Raw() interface{} {
	_ = f.Collect()
	_ = f.Resolve()

	return f
}
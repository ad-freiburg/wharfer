package wrap

import (
	"strings"
)

type StringSliceFlag []string

func (s *StringSliceFlag) String() string {
	return strings.Join(*s, ",")
}

func (s *StringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

type WrappedCommand interface {
	InitFlags()
	ParseToArgs(rawArgs []string) []string
}

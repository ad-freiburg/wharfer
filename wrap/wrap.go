package wrap

import (
	"fmt"
	"os"
	"os/user"
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

func PrependUsername(s string) string {
	user, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve username")
		os.Exit(3)
	}

	if !strings.HasPrefix(s, user.Username+"_") {
		return user.Username + "_" + s
	}
	return s
}

func IsHexOnly(s string) bool {
	for _, c := range s {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return false
		}
	}
	return true
}

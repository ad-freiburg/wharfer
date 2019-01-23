package wrap

import (
	"encoding/hex"
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
	_, err := hex.DecodeString(s)
	return err == nil
}

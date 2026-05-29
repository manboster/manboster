package gatekeeper

import (
	"errors"
	"fmt"
)

var ErrHachimiSuspicious = errors.New("suspicious")
var ErrHachimiSafe = errors.New("hachimi allowed")
var ErrHachimiDenied = fmt.Errorf("hachimi thinks it's suspicious and user denied it")

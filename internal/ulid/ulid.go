package ulid

import (
	"fmt"
	"github.com/oklog/ulid/v2"
)

type Ulid struct{}

func NewUlid() Ulid {
	return Ulid{}
}

func (u *Ulid) String() string {
	return fmt.Sprintf("%s", ToString())
}

func (u *Ulid) Json() string {
	return fmt.Sprintf("{\"ulid\":\"%s\"}", ToString())
}

func ToString() string {
	return ulid.Make().String()
}

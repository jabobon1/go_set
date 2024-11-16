package set

import "golang.org/x/exp/constraints"

type AllowedTypes interface {
	constraints.Float | constraints.Integer | constraints.Signed | constraints.Unsigned | string
}

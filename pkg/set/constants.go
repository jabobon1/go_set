package set

import "golang.org/x/exp/constraints"

// AllowedTypes interface with allowed types for Set struct
type AllowedTypes interface {
	constraints.Float | constraints.Integer | constraints.Signed | constraints.Unsigned | string
}

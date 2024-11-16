package set

type ISet[T AllowedTypes] interface {
	Add(value T)
	Remove(value T)
	In(value T) bool
	All() []T
	Clear() []T
	Intersect(other Set[T]) Set[T]
	Diff(other Set[T]) Set[T]
	Union(other Set[T]) Set[T]
	Len() int
}

package generics

type HasID interface {
	GetID() string
}

func FindByID[T HasID](items []T, id string) T {
	for _, it := range items {
		if it.GetID() == id {
			return it
		}
	}
	var zero T
	return zero
}

package lib

// TODO Determine common lifecycle or handler mechanisms that any Engine
// implementation must provide.
type Engine interface {
	Start()
}

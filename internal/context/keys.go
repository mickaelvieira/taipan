package context

type key int

// https://blog.golang.org/context#TOC_3.2.
const (
	KeyPrincipalID key = iota
	// ...
)

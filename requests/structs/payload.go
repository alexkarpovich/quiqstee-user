package structs

type Base struct {
	Errors map[string]string
}

type Payload interface {
	Validate() bool
} 


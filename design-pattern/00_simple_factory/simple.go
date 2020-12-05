package simplefactory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	switch t {
	case 1:
		return &hiApi{}
	case 2:
		return &helloApi{}
	}
	return nil
}

type hiApi struct {
}

func (h *hiApi) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloApi struct {
}

func (h *helloApi) Say(name string) string {
	return fmt.Sprintf("hello, %v", name)
}

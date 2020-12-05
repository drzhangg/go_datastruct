package facade

type API interface {
	Test() string
}

type apiImple struct {
}

type aModuleAPI interface {
	TestA() string
}

type aModuleImpl struct {
}

func (*aModuleImpl) TestA() string {
	return "a module string"
}

func NewAModuleAPI() aModuleAPI {
	return nil
}

type bModuleAPI interface {
	TestB() string
}

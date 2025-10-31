package exceptions

type ErrorCode string

type ErrorProps struct {
	Code    ErrorCode
	Message string
	//TODO: Verificar para adicionar parametros dinamicos
}

type PropFinder interface {
	GetProps(ErrorCode) ErrorProps
	IsValid(ErrorCode) bool
}

var finder PropFinder

func RegisterPropFinder(f PropFinder) {
	finder = f
}

func (s ErrorCode) GetProps() ErrorProps {
	if finder == nil {
		return ErrorProps{Code: s, Message: "errors.finder.not.registered"}
	}
	return finder.GetProps(s)
}

func (s ErrorCode) IsValid() bool {
	if finder == nil {
		return false
	}
	return finder.IsValid(s)
}

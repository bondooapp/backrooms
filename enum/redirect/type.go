package redirect

type Type string

const (
	LOGIN_PAGE Type = "login_page"
)

func (enum *Type) String() string {
	return string(*enum)
}

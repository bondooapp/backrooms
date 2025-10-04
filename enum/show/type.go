package show

type Type string

const (
	ERROR_MESSAGE Type = "error_message"
	NOTIFICATION  Type = "notification"
	REDIRECT      Type = "redirect"
	SILENT        Type = "silent"
	WARN_MESSAGE  Type = "warn_message"
)

func (enum *Type) String() string {
	return string(*enum)
}

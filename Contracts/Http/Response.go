package Http

type Response interface {
	Prepare(Request) Response
	SetContent(interface{}) Response
	SetStatusCode(int) Response
	Header(string, string, ...bool) Response
	Status() int
	Content() []byte
	ContentString() string
	SendHeaders() Response
	SendContent() Response
	Send() Response
	IsInvalid() bool
	IsInformational() bool
	IsSuccessful() bool
	IsRedirection() bool
	IsClientError() bool
	IsServerError() bool
	IsOk() bool
	IsForbidden() bool
	IsNotFound() bool
	IsRedirect(...string) bool
	IsEmpty() bool
}

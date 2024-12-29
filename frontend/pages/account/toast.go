package frontend

type ToastType string

func (t ToastType) String() string {
	return string(t)
}

const (
	SuccessToast ToastType = "success"
	ErrorToast   ToastType = "error"
	InfoToast    ToastType = "info"
)

package contract

import "net/http"

const KernelKey = "gser:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}

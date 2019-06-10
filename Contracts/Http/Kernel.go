package Http

import (
	"github.com/larisgo/framework/Foundation"
	"github.com/larisgo/framework/Http"
	"github.com/valyala/fasthttp"
)

type Kernel interface {

	/**
	 * Bootstrap the application for HTTP requests.
	 *
	 * @return void
	 */
	Bootstrap()

	/**
	 * Handle an incoming HTTP request.
	 *
	 * @param  Request  request
	 * @return Response
	 */
	Handle(*fasthttp.RequestCtx)

	/**
	 * Perform any final actions for the request lifecycle.
	 *
	 * @param  Request  request
	 * @param  Response  response
	 * @return void
	 */
	Terminate(*Http.Request, *Http.Response)

	/**
	 * Get the Laravel application instance.
	 *
	 * @return Application
	 */
	GetApplication() *Foundation.Application
}

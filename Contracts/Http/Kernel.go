package Http

import (
	"github.com/larisgo/framework/Foundation"
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
	Handle(request interface{})

	/**
	 * Perform any final actions for the request lifecycle.
	 *
	 * @param  Request  request
	 * @param  Response  response
	 * @return void
	 */
	Terminate(request interface{}, response interface{})

	/**
	 * Get the Laravel application instance.
	 *
	 * @return Application
	 */
	GetApplication() *Foundation.Application
}

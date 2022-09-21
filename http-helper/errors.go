package httphelper

import "errors"

var (
	ErrBadRequest                   = errors.New("400") // RFC 9110, 15.5.1
	ErrUnauthorized                 = errors.New("401") // RFC 9110, 15.5.2
	ErrPaymentRequired              = errors.New("402") // RFC 9110, 15.5.3
	ErrForbidden                    = errors.New("403") // RFC 9110, 15.5.4
	ErrNotFound                     = errors.New("404") // RFC 9110, 15.5.5
	ErrMethodNotAllowed             = errors.New("405") // RFC 9110, 15.5.6
	ErrNotAcceptable                = errors.New("406") // RFC 9110, 15.5.7
	ErrProxyAuthRequired            = errors.New("407") // RFC 9110, 15.5.8
	ErrRequestTimeout               = errors.New("408") // RFC 9110, 15.5.9
	ErrConflict                     = errors.New("409") // RFC 9110, 15.5.10
	ErrGone                         = errors.New("410") // RFC 9110, 15.5.11
	ErrLengthRequired               = errors.New("411") // RFC 9110, 15.5.12
	ErrPreconditionFailed           = errors.New("412") // RFC 9110, 15.5.13
	ErrRequestEntityTooLarge        = errors.New("413") // RFC 9110, 15.5.14
	ErrRequestURITooLong            = errors.New("414") // RFC 9110, 15.5.15
	ErrUnsupportedMediaType         = errors.New("415") // RFC 9110, 15.5.16
	ErrRequestedRangeNotSatisfiable = errors.New("416") // RFC 9110, 15.5.17
	ErrExpectationFailed            = errors.New("417") // RFC 9110, 15.5.18
	ErrTeapot                       = errors.New("418") // RFC 9110, 15.5.19 (Unused)
	ErrMisdirectedRequest           = errors.New("421") // RFC 9110, 15.5.20
	ErrUnprocessableEntity          = errors.New("422") // RFC 9110, 15.5.21
	ErrLocked                       = errors.New("423") // RFC 4918, 11.3
	ErrFailedDependency             = errors.New("424") // RFC 4918, 11.4
	ErrTooEarly                     = errors.New("425") // RFC 8470, 5.2.
	ErrUpgradeRequired              = errors.New("426") // RFC 9110, 15.5.22
	ErrPreconditionRequired         = errors.New("428") // RFC 6585, 3
	ErrTooManyRequests              = errors.New("429") // RFC 6585, 4
	ErrRequestHeaderFieldsTooLarge  = errors.New("431") // RFC 6585, 5
	ErrUnavailableForLegalReasons   = errors.New("451") // RFC 7725, 3

	ErrInternalServerError           = errors.New("500") // RFC 9110, 15.6.1
	ErrNotImplemented                = errors.New("501") // RFC 9110, 15.6.2
	ErrBadGateway                    = errors.New("502") // RFC 9110, 15.6.3
	ErrServiceUnavailable            = errors.New("503") // RFC 9110, 15.6.4
	ErrGatewayTimeout                = errors.New("504") // RFC 9110, 15.6.5
	ErrHTTPVersionNotSupported       = errors.New("505") // RFC 9110, 15.6.6
	ErrVariantAlsoNegotiates         = errors.New("506") // RFC 2295, 8.1
	ErrInsufficientStorage           = errors.New("507") // RFC 4918, 11.5
	ErrLoopDetected                  = errors.New("508") // RFC 5842, 7.2
	ErrNotExtended                   = errors.New("510") // RFC 2774, 7
	ErrNetworkAuthenticationRequired = errors.New("511") // RFC 6585, 6
)

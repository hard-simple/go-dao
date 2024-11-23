package dao

// Pagination keeps page by page processing information. This entity uses for requests as well as for responses.
// In other words, it could be used to define the pagination properties whenever you send a request to the
// target storage as well as using it from response for metrics/observation/further identical operation requests.
//
// Usually, the workflow of requests/responses looks like next:
//
// 1. Send empty Pagination instance in request.
//
// 2. Receive response and take Pagination.NextToken for the next request.
//
// 3. Handle response and start over step 1 if you want to continue.
type Pagination struct {

	// It's a page offset to start from. It's optional. The implementation should take care
	// of the default definition.
	// It has a lower priority than PrevToken in requests.
	Offset *uint

	// Size of the page. Min 1. It's optional. If it isn't defined for
	// the request then the default value of an implementation should be taken.
	Size *uint

	// Next page token. It's optional. It may contain value for requests only in case the service layer
	// is aware about the token format or service layer has taken it from previous DAO request.
	// Usually, service layer doesn't know anything about this format and only share them between
	// each request and response.
	// NextToken has more priority than PrevToken in requests.
	//
	// If you got this from response then the service layer may use it for the next similar paginated operation request.
	NextToken *[]byte

	// Previous page token. It's optional.
	// PrevToken has more priority than Offset in requests.
	// PrevToken has a lower priority than NextToken in requests.
	// If it's defined in a request and there is no value for NextToken then an implementation
	// should use next logic: parse previous token value + Size.
	// Usually, it uses for observation purpose in case of response.
	PrevToken *[]byte

	// Overall entity size. It's optional. Some implementations can calculate it during the main process
	// and return it in a response.
	// It doesn't have any value in request.
	Total *uint64

	// It determines whether there is a data after this page or not. It's optional.
	// Some implementations may use it to improve user experience and efficiently handle further queries/requests.
	// If it isn't specified then it should be treated as unknown.
	// If true then there is at least one item in the next page otherwise 0.
	// Implementations should ignore this field while seeing it in request. Only response may bring value
	// for a response consumer.
	HasNext *bool
}

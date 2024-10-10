package service

var (
	InternalCode = map[int]string{
		100200: "Success",
		100400: "Invalid Request",
		100401: "Unauthorized Request",
		100403: "Forbidden Request",
		100404: "Resource Not Found",
		100500: "Internal Server Error",
		100503: "Service Unavailable",
	}
)

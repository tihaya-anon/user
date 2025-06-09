package enum

var MSG = struct {
	SUCCESS string

	CUSTOMER_ERROR   string
	VALIDATION_ERROR string
	MISSING_TOKEN    string
	INVALID_TOKEN    string

	SYSTEM_ERROR string

	THIRD_PARTY_ERROR string
}{
	SUCCESS: "Success",

	CUSTOMER_ERROR:   "Customer Error",
	VALIDATION_ERROR: "Validation Error",
	MISSING_TOKEN:    "Must Provide Token",
	INVALID_TOKEN:    "Invalid Token",

	SYSTEM_ERROR: "System Error",

	THIRD_PARTY_ERROR: "Third Party Error",
}

package gerror

func StatusText(errorCode uint32) string {
	switch errorCode {

	//Client side
	case ErrorBindData:
		return "Failed to bind data"
	case ErrorValidData:
		return "Failed to valid data"

	//Serevr side
	case ErrorConnect:
		return "Failed to conenct database"
	case ErrorRetriveData:
		return "Failed to retrive data"
	case ErrorSaveData:
		return "Failed to save data"
	case ErrorNotFound:
		return "Data not found"
	case ErrorInternal:
		return "Server error"
	case ErrorBadRequest:
		return "Bad request"
	case ErrorPermission:
		return "Permission denied"

	}
	return "Unknown error"
}

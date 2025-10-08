package client

// FgaRequiredParamError Provides access to the body, error and model on returned errors.
type FgaRequiredParamError struct {
	error string
	param string
}

// Error returns non-empty string if there was an error.
func (e FgaRequiredParamError) Error() string {
	if e.error == "" {
		return "Required parameter " + e.Param() + " was not provided"
	}
	return e.error
}

// Param returns the name of the missing parameter
func (e FgaRequiredParamError) Param() string {
	return e.param
}

// FgaInvalidError Provides access to the body, error and model on returned errors.
type FgaInvalidError struct {
	error       string
	param       string
	description string
}

// Error returns non-empty string if there was an error.
func (e FgaInvalidError) Error() string {
	if e.error == "" {
		return "Parameter " + e.Param() + " is not a valid " + e.description
	}
	return e.error
}

// Param returns the name of the invalid parameter
func (e FgaInvalidError) Param() string {
	return e.param
}

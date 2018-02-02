package errdefs

type causer interface {
	Cause() error
}

func getImplementer(err error) error {
	switch e := err.(type) {
	case
		ErrNotFound,
		ErrInvalidParameter,
		ErrConflict,
		ErrUnauthorized,
		ErrUnavailable,
		ErrForbidden,
		ErrSystem,
		ErrNotModified,
		ErrNotImplemented,
		ErrUnknown:
		return e
	case causer:
		return getImplementer(e.Cause())
	default:
		return err
	}
}

// ErrNotFound signals that the requested object doesn't exist
type ErrNotFound interface {
	NotFound()
}

// IsNotFound returns if the passed in error is an ErrNotFound
func IsNotFound(err error) bool {
	_, ok := getImplementer(err).(ErrNotFound)
	return ok
}

// ErrInvalidParameter signals that the user input is invalid
type ErrInvalidParameter interface {
	InvalidParameter()
}

// IsInvalidParameter returns if the passed in error is an ErrInvalidParameter
func IsInvalidParameter(err error) bool {
	_, ok := getImplementer(err).(ErrInvalidParameter)
	return ok
}

// ErrConflict signals that some internal state conflicts with the requested action and can't be performed.
// A change in state should be able to clear this error.
type ErrConflict interface {
	Conflict()
}

// IsConflict returns if the passed in error is an ErrConflict
func IsConflict(err error) bool {
	_, ok := getImplementer(err).(ErrConflict)
	return ok
}

// ErrUnauthorized is used to signify that the user is not authorized to perform a specific action
type ErrUnauthorized interface {
	Unauthorized()
}

// IsUnauthorized returns if the the passed in error is an ErrUnauthorized
func IsUnauthorized(err error) bool {
	_, ok := getImplementer(err).(ErrUnauthorized)
	return ok
}

// ErrUnavailable signals that the requested action/subsystem is not available.
type ErrUnavailable interface {
	Unavailable()
}

// IsUnavailable returns if the passed in error is an ErrUnavailable
func IsUnavailable(err error) bool {
	_, ok := getImplementer(err).(ErrUnavailable)
	return ok
}

// ErrForbidden signals that the requested action cannot be performed under any circumstances.
// When a ErrForbidden is returned, the caller should never retry the action.
type ErrForbidden interface {
	Forbidden()
}

// IsForbidden returns if the passed in error is an ErrForbidden
func IsForbidden(err error) bool {
	_, ok := getImplementer(err).(ErrForbidden)
	return ok
}

// ErrSystem signals that some internal error occurred.
// An example of this would be a failed mount request.
type ErrSystem interface {
	ErrSystem()
}

// IsSystem returns if the passed in error is an ErrSystem
func IsSystem(err error) bool {
	_, ok := getImplementer(err).(ErrSystem)
	return ok
}

// ErrNotModified signals that an action can't be performed because it's already in the desired state
type ErrNotModified interface {
	NotModified()
}

// IsNotModified returns if the passed in error is a NotModified error
func IsNotModified(err error) bool {
	_, ok := getImplementer(err).(ErrNotModified)
	return ok
}

// ErrNotImplemented signals that the requested action/feature is not implemented on the system as configured.
type ErrNotImplemented interface {
	NotImplemented()
}

// IsNotImplemented returns if the passed in error is an ErrNotImplemented
func IsNotImplemented(err error) bool {
	_, ok := getImplementer(err).(ErrNotImplemented)
	return ok
}

// ErrUnknown signals that the kind of error that occurred is not known.
type ErrUnknown interface {
	Unknown()
}

// IsUnknown returns if the passed in error is an ErrUnknown
func IsUnknown(err error) bool {
	_, ok := getImplementer(err).(ErrUnknown)
	return ok
}

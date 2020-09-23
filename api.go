package locky

// Client provides the API for Locky.
type Client interface {
	// Lock attempts to acquire a key. You can use LockOptions to configure the
	// kind of lock to acquire. The default lock is a simple non-expiring mutex
	// lock.
	Lock(key string) error

	// Unlock releases a key.
	Unlock(key string) error
}

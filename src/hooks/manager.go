package hooks

type Interface interface {
	Getenv(key string)
	// Stat(file string)
	// Open(file string)
	// Chdir(dir string)
}

/*
var manager atomic.Value
// SetManager sets the test logger implementation for the current process.
// It must be called only once, at process startup.
func SetManager(impl Interface) {
	if manager.Load() != nil {
		panic("testlog: SetLogger must be called only once")
	}
	manager.Store(&impl)
}
// Logger returns the current test logger implementation.
// It returns nil if there is no logger.
func Manager() Interface {
	impl := manager.Load()
	if impl == nil {
		return nil
	}
	return *impl.(*Interface)
}
// Getenv calls Logger().Getenv, if a logger has been set.
func Getenv(name string) {
	if man := Manager(); man != nil {
		man.Getenv(name)
	}
}
*/

var impl *Interface

// SetManager sets the test logger implementation for the current process.
// It must be called only once, at process startup.
func SetManager(_impl Interface) {
	if impl != nil {
		panic("testlog: SetLogger must be called only once")
	}
	impl = &_impl
}

// Logger returns the current test logger implementation.
// It returns nil if there is no logger.
func Manager() Interface {
	if impl == nil {
		return nil
	}
	return *impl
}

// Getenv calls Logger().Getenv, if a logger has been set.
func Getenv(name string) {
	if m := Manager(); m != nil {
		m.Getenv(name)
	}
}

package dsr

type DSR struct {
	locker *lock
	hasher Hasher
	config *Config
}

// New provides initialization of the dsr app
func New(conf *Config) (*DSR, error) {
	if conf == nil {
		conf = DefaultConfig()
	}
	dsr := &DSR{
		config: conf,
	}
	return dsr, nil
}

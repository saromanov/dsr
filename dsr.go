package dsr

type DSR struct {
	locker *lock
}

// New provides initialization of the dsr app
func New(conf *Config) (*DSR, error) {
	if conf == nil {
		conf = DefaultConfig()
	}
	dsr := &DSR{}
	return dsr, nil
}

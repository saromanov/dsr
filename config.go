package dsr

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/hashicorp/memberlist"
)

const (
	DefaultLogLevel = "DEBUG"
)

type EvictionPolicy string

const (
	DefaultLRUSamples int            = 5
	LRUEviction       EvictionPolicy = "LRU"
)

type DMapConfig struct {
	MaxIdleDuration time.Duration
	TTLDuration     time.Duration
	MaxKeys         int
	LRUSamples      int
	EvictionPolicy  EvictionPolicy
}

type CacheConfig struct {
	MaxIdleDuration time.Duration
	TTLDuration     time.Duration
	MaxKeys         int
	LRUSamples      int
}
type Config struct {
	LogLevel         string
	Name             string
	KeepAlivePeriod  time.Duration
	DialTimeout      time.Duration
	Peers            []string
	PartitionCount   uint64
	BackupCount      int
	BackupMode       int
	LoadFactor       float64
	Hasher           Hasher
	KeyFile          string
	LogOutput        io.Writer
	Logger           *log.Logger
	MemberlistConfig *memberlist.Config
}

func NewMemberlistConfig(env string) (*memberlist.Config, error) {
	switch {
	case env == "local":
		return memberlist.DefaultLocalConfig(), nil
	case env == "lan":
		return memberlist.DefaultLANConfig(), nil
	case env == "wan":
		return memberlist.DefaultWANConfig(), nil
	}
	return nil, fmt.Errorf("unknown env: %s", env)
}

func DefaultConfig() *Config{
	return &Config{}
}

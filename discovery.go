package dsr

import (
	"github.com/hashicorp/memberlist"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type discovery struct {
	list  *memberlist.Memberlist
	entry *logrus.Entry
}

// New provides initialization of the discovery app
func New(conf *Config) (*discovery, error) {
	list, err := memberlist.Create(conf.MemberlistConfig)
	if err != nil {
		return nil, xerrors.Errorf("unable to init memberlist config: %v", err)
	}

	return &discovery{
		list: list,
	}, nil
}

func (d *discovery) Join(peer string) error {
	nodes, err := d.list.Join([]string{peer})
	if err != nil {
		return xerrors.Errorf("unable to join node: %v", err)
	}

	if nodes == 0 {
		d.entry.Warn("unable to join nodes")
	}
	return nil
}

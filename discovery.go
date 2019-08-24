package dsr

import (
	"fmt"

	"github.com/hashicorp/memberlist"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type discovery struct {
	list  *memberlist.Memberlist
	log *logrus.Entry
}

// NewDiscovery provides initialization of the discovery app
func NewDiscovery(conf *Config) (*discovery, error) {
	list, err := memberlist.Create(conf.MemberlistConfig)
	if err != nil {
		return nil, xerrors.Errorf("unable to init memberlist config: %v", err)
	}

	return &discovery{
		list: list,
	}, nil
}

// Join provides joining of the peer
func (d *discovery) Join(peer string) error {
	nodes, err := d.list.Join([]string{peer})
	if err != nil {
		return xerrors.Errorf("unable to join node: %v", err)
	}

	if nodes == 0 {
		d.log.Warn("unable to join nodes")
	}
	return nil
}

// GetMembers returns list of the members
func (d *discovery) GetMembers() {
	members := d.list.Members()
	for _, m := range members {
		fmt.Println(m)
	}
}

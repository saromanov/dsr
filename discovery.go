package dsr

import (
	"github.com/hashicorp/memberlist"
	"golang.org/x/xerrors"
)

type discovery struct {
	list *memberlist.Memberlist
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
	_, err := d.list.Join([]string{peer})
	if err != nil {
		return xerrors.Errorf("unable to join node: %v", err)
	}

	return nil
}

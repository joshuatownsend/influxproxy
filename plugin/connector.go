package plugin

import (
	"errors"
)

type Connector struct {
	e Exposer
}

func NewConnector(e Exposer) (*Connector, error) {
	if e == nil {
		return nil, errors.New("No exposer provided")
	} else {
		c := &Connector{
			e: e,
		}
		return c, nil
	}
}

func (c *Connector) Ping(in []*interface{}, pong *bool) error {
	*pong = c.e.Ping()
	return nil
}

func (c *Connector) Describe(in []*interface{}, description *string) error {
	*description = c.e.Describe()
	return nil
}

func (c *Connector) Run(in []*interface{}, description *string) error {
	*description = c.e.Run(in)
	return nil
}

type Exposer interface {
	Ping() bool
	Describe() string
	Run(in []*interface{}) string
}

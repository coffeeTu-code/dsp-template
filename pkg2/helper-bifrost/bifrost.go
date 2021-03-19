package bifrost

import (
	"errors"

	"dsp-template/pkg2/helper-bifrost/container"
)

type Bifrost struct {
	DataStreamers map[string]Streamer
	logger        *BiLogger
}

func NewBifrost() *Bifrost {
	return &Bifrost{
		DataStreamers: make(map[string]Streamer),
	}
}

func (l *Bifrost) Get(name string, key container.MapKey) (interface{}, error) {
	s, ok := l.DataStreamers[name]
	if !ok {
		return nil, errors.New("not found streamer[" + name + "]")
	}
	c := s.GetContainer()
	if c == nil {
		return nil, errors.New("contain is nil, streamer[" + name + "]")
	}
	return c.Get(key)
}

func (l *Bifrost) Register(name string, streamer Streamer) error {
	if _, ok := l.DataStreamers[name]; ok {
		return errors.New("streamer[" + name + "] has already exist")
	}
	l.DataStreamers[name] = streamer
	return nil
}

func (l *Bifrost) GetStreamer(name string) (Streamer, error) {
	s, ok := l.DataStreamers[name]
	if !ok {
		return nil, errors.New("not found streamer[" + name + "]")
	}
	return s, nil
}

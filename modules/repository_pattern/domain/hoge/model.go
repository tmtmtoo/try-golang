package hoge

import (
	"errors"
	"repository_pattern/domain/primitives"
)

type HogeVisitor interface {
	VisitUnprocessedHoge(hoge *UnprocessedHoge) error
	VisitCanceledHoge(hoge *CanceledHoge) error
	VisitProcessedHoge(hoge *ProcessedHoge) error
}

type Hoge interface {
	Accept(v HogeVisitor) error
	Cancel(reason primitives.Text) (Hoge, error)
	Process(piyoValue primitives.Text) (Hoge, error)
}

type UnprocessedHoge struct {
	Id    primitives.Id
	Value primitives.Text
}

func NewUnprocessedHoge(id []byte, value string) (*UnprocessedHoge, error) {
	parsedId, err := primitives.ParseIdBytes(id)

	if err != nil {
		return nil, err
	}

	hoge := &UnprocessedHoge{
		Id:    parsedId,
		Value: primitives.NewText(value),
	}

	return hoge, nil
}

func (h *UnprocessedHoge) Accept(v HogeVisitor) error {
	return v.VisitUnprocessedHoge(h)
}

func (h *UnprocessedHoge) Cancel(reason primitives.Text) (Hoge, error) {
	return &CanceledHoge{
		Id:     h.Id,
		Reason: reason,
	}, nil
}

func (h *UnprocessedHoge) Process(piyoValue primitives.Text) (Hoge, error) {
	return &ProcessedHoge{
		Id:   h.Id,
		Piyo: &Piyo{Id: primitives.GenerateId(), Value: piyoValue},
	}, nil
}

type CanceledHoge struct {
	Id     primitives.Id
	Reason primitives.Text
}

func NewCanceledHoge(id []byte, reason string) (*CanceledHoge, error) {
	parsedId, err := primitives.ParseIdBytes(id)

	if err != nil {
		return nil, err
	}

	hoge := &CanceledHoge{
		Id:     parsedId,
		Reason: primitives.NewText(reason),
	}

	return hoge, nil
}

func (h *CanceledHoge) Accept(v HogeVisitor) error {
	return v.VisitCanceledHoge(h)
}

func (h *CanceledHoge) Cancel(reason primitives.Text) (Hoge, error) {
	return h, nil
}

func (h *CanceledHoge) Process(piyoValue primitives.Text) (Hoge, error) {
	return nil, errors.New("already canceled")
}

type Piyo struct {
	Id    primitives.Id
	Value primitives.Text
}

func NewPiyo(id []byte, value string) (*Piyo, error) {
	parsedId, err := primitives.ParseIdBytes(id)

	if err != nil {
		return nil, err
	}

	piyo := &Piyo{
		Id:    parsedId,
		Value: primitives.NewText(value),
	}

	return piyo, nil
}

type ProcessedHoge struct {
	Id   primitives.Id
	Piyo *Piyo
}

func NewProcessedHoge(id []byte, piyo *Piyo) (*ProcessedHoge, error) {
	parsedId, err := primitives.ParseIdBytes(id)

	if err != nil {
		return nil, err
	}

	hoge := &ProcessedHoge{
		Id:   parsedId,
		Piyo: piyo,
	}

	return hoge, nil
}

func (h *ProcessedHoge) Accept(v HogeVisitor) error {
	return v.VisitProcessedHoge(h)
}

func (h *ProcessedHoge) Cancel(reason primitives.Text) (Hoge, error) {
	return nil, errors.New("already processed")
}

func (h *ProcessedHoge) Process(piyoValue primitives.Text) (Hoge, error) {
	return h, nil
}

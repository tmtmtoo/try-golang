package primitives

import "github.com/google/uuid"

type Id [16]byte

func NewId(bytes [16]byte) Id {
	return Id(bytes)
}

func GenerateId() Id {
	return Id(uuid.New())
}

func (id Id) String() string {
	return uuid.UUID(id).String()
}

func ParseIdString(s string) (Id, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return Id{}, err
	}
	return Id(u), nil
}

func (id Id) Bytes() [16]byte {
	return id
}

func ParseIdBytes(b []byte) (Id, error) {
	u, err := uuid.FromBytes(b)
	if err != nil {
		return Id{}, err
	}
	return Id(u), nil
}

type Text string

func NewText(s string) Text {
	return Text(s)
}

func (t Text) String() string {
	return string(t)
}

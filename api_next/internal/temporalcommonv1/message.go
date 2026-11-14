package common

// These are dummy messages, never used

type Payloads struct {
	Payloads []*Payload
}

type Payload struct {
	Metadata map[string][]byte
	Data     []byte
}

func (*Payload) Reset()         { panic("NOT IMPLEMENTED") }
func (*Payload) ProtoMessage()  { panic("NOT IMPLEMENTED") }
func (*Payload) String() string { panic("NOT IMPLEMENTED") }

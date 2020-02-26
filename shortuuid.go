package shortuuid

import (
	"strings"

	"github.com/gofrs/uuid"
)

// DefaultEncoder is the default encoder uses when generating new UUIDs, and is
// based on Base57.
var DefaultEncoder = &base57{newAlphabet(DefaultAlphabet)}

var (
	NameSpaceDNS  = uuid.Must(uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceURL  = uuid.Must(uuid.FromString("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceOID  = uuid.Must(uuid.FromString("6ba7b812-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceX500 = uuid.Must(uuid.FromString("6ba7b814-9dad-11d1-80b4-00c04fd430c8"))
)

// Encoder is an interface for encoding/decoding UUIDs to strings.
type Encoder interface {
	Encode(uuid.UUID) string
	Decode(string) (uuid.UUID, error)
}

// New returns a new UUIDv4, encoded with base57.
func New() string {
	return DefaultEncoder.Encode(uuid.Must(uuid.NewV4()))
}

// NewWithEncoder returns a new UUIDv4, encoded with enc.
func NewWithEncoder(enc Encoder) string {
	return enc.Encode(uuid.Must(uuid.NewV4()))
}

// NewWithNamespace returns a new UUIDv5 (or v4 if name is empty), encoded with base57.
func NewWithNamespace(name string) string {
	var u uuid.UUID

	switch {
	case name == "":
		u = uuid.Must(uuid.NewV4())
	case strings.HasPrefix(name, "http"):
		u = uuid.NewV5(NameSpaceURL, name)
	default:
		u = uuid.NewV5(NameSpaceDNS, name)
	}

	return DefaultEncoder.Encode(u)
}

// NewWithAlphabet returns a new UUIDv4, encoded with base57 using the
// alternative alphabet abc.
func NewWithAlphabet(abc string) string {
	enc := base57{newAlphabet(abc)}
	return enc.Encode(uuid.Must(uuid.NewV4()))
}

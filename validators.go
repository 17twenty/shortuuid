package shortuuid

import (
	"regexp"

	"github.com/gofrs/uuid"
)

var (
	v4SchemaRegEx = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
)

// IsUUIDv4 returns true if the provided uuid matches the v5 schema
func IsUUIDv4(u uuid.UUID) bool {
	return v4SchemaRegEx.MatchString(u.String())
}

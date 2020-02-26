package shortuuid

import (
	"testing"

	"github.com/gofrs/uuid"
)

var newTest = []struct {
	uuid     string
	expected bool
}{
	{"51e5fe10-5325-4a32-bce8-7ebe9708c453", true},
	{"50ee7e88-c29e-4e38-b94b-6937b5a4c8c ", false},
	{"aeb11973-3be9-47fc-8059-ecdac9ab1887", true},
	{"abcdefgh-ijkl-mnop-qrst-uvwxyzabcdef", false},
	{"688831FF-D2D6-4DA9-8676-AC4C9436BC0E", true},
}

func TestValidations(t *testing.T) {
	for _, test := range newTest {
		uid, err := uuid.FromString(test.uuid)
		if err != nil && test.expected {
			t.Error("Failed with", err, "but should've passed")
		} else if err != nil {
			continue
		}
		if IsUUIDv4(uid) != test.expected {
			t.Errorf("Expected uid %s to result in %v but got %v", uid, test.expected, !test.expected)
		}
	}
}

func TestOne(t *testing.T) {

	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)

	}
	if IsUUIDv4(uid) == false {
		t.Errorf("Expected uid %s to be valid", uid.String())
	}

}

package fireness

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentify(t *testing.T) {
	localhost, err := Localhost(
		WithLocalhostExclude("127.0.0.1"),
		WithLocalhostExcludeIPv6(true),
	)

	assert.NoError(t, err)
	log.Printf("localhost: %+v\n", localhost)
}

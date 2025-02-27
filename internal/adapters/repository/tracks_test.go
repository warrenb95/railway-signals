package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/warrenb95/railway-signals/internal/domain"
)

func TestCreateTrack(t *testing.T) {
	tests := map[string]struct {
		req *domain.Track

		errorContains string
	}{
		"successfully create track in the store": {
			req: &domain.Track{
				ID:     1,
				Source: "source",
				Target: "target",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := testDB.CreateTrack(context.Background(), test.req)
			if test.errorContains != "" {
				require.ErrorContains(t, err, test.errorContains, "create error contains")
				return
			}
			require.NoError(t, err, "creating new track")

			track, err := testDB.GetTrack(context.Background(), test.req.ID)
			require.NoError(t, err, "getting track")

			assert.Equal(t, test.req, track, "track")
		})
	}
}

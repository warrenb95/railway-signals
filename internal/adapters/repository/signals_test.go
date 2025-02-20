package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/warrenb95/railway-signals/internal/domain"
)

func TestCreateSignals(t *testing.T) {
	tests := map[string]struct {
		req *domain.Signal

		errorContains string
	}{
		"successfully create signal in the store": {
			req: &domain.Signal{
				ID:   1,
				Name: "signal",
				ELR:  "asdf",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := testDB.CreateSignal(context.Background(), test.req)
			if test.errorContains != "" {
				require.ErrorContains(t, err, test.errorContains, "create error contains")
				return
			}
			require.NoError(t, err, "creating new signal")

			track, err := testDB.GetSignal(context.Background(), test.req.ID)
			require.NoError(t, err, "getting signal")

			assert.Equal(t, test.req, track, "signal")
		})
	}
}

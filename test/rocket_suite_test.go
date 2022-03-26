// +acceptance

package test

import (
	"context"
	"testing"

	"github.com/goethesum/gRPC/internal/rocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("adds a new rocket successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					ID:   "ae7b0bf0-fe75-4176-b20a-3ca1371f3226",
					Name: "V1",
					Type: "Faclon Heavy",
				},
			},
		)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "ae7b0bf0-fe75-4176-b20a-3ca1371f3226", resp.Rocket.Id)
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}

package common

import (
	"card_service/test/utils/container"
	"context"
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	"go.uber.org/dig"
)

func SetupTesting(t *testing.T) (*container.Components, context.Context, gomega.Gomega) {
	g := GetGomega(t)

	c, err := container.BuildContainer()
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), fmt.Sprintf("unable to build container: %v", dig.RootCause(err)))

	ctx := context.Background()

	return c, ctx, g
}

package tests

import (
	cardservice "card_service/generated/proto"
	cardservicechecks "card_service/test/utils/assertions/cardservice"
	"card_service/test/utils/common"
	"card_service/test/utils/container"
	"card_service/test/utils/reports"
	"context"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/severity"
	"github.com/onsi/gomega"
)

// Basic valid transaction request
func newValidTransactionReq() *cardservice.CreateTransactionRequest {
	return &cardservice.CreateTransactionRequest{
		FromCardId: 1,
		ToCardId:   2,
		Amount:     100.50,
	}
}

// Helper to create a transaction and verify basic fields
func CreateTestTransaction(t *testing.T, c *container.Components, ctx context.Context, g gomega.Gomega) *cardservice.Transaction {
	t.Helper()
	req := newValidTransactionReq()

	created, err := c.CardService.CreateTransaction(ctx, req)
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), "CreateTransaction should succeed with valid request")
	g.Expect(created).ShouldNot(gomega.BeNil())

	// Basic field checks
	g.Expect(created.FromCardId).To(gomega.Equal(req.FromCardId))
	g.Expect(created.ToCardId).To(gomega.Equal(req.ToCardId))
	g.Expect(created.Amount).To(gomega.Equal(req.Amount))

	return created
}

// Create transaction
func TestCreateTransaction(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.TransactionFeature, reports.TransactionSuite,
		allure.Severity(severity.Critical),
		allure.Tags(reports.TransactionTag),
		allure.Name("Create transaction"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)
			CreateTestTransaction(t, c, ctx, g)
		}),
	)
}

// Get transaction
func TestGetTransaction(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.TransactionFeature, reports.TransactionSuite,
		allure.Severity(severity.Critical),
		allure.Tags(reports.TransactionTag),
		allure.Name("Get transaction"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)
			created := CreateTestTransaction(t, c, ctx, g)

			resp, err := c.CardService.GetTransaction(ctx, &cardservice.GetTransactionRequest{
				FromCardId: created.FromCardId,
			})
			g.Expect(err).ShouldNot(gomega.HaveOccurred(), "GetTransaction should not return error")
			g.Expect(resp).ShouldNot(gomega.BeNil())
			g.Expect(resp.Transactions).ShouldNot(gomega.BeEmpty())

			found := false
			for _, tx := range resp.Transactions {
				if tx.Id == created.Id {
					cardservicechecks.CheckTransaction(g, tx, created)
					found = true
					break
				}
			}
			g.Expect(found).To(gomega.BeTrue(), "Created transaction must be found in GetTransaction response")
		}),
	)
}

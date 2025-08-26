package tests

import (
	cardservice "card_service/generated/proto"
	cardservicechecks "card_service/test/utils/assertions/cardservice"
	solutions "card_service/test/utils/assertions/common"
	"card_service/test/utils/common"
	"card_service/test/utils/container"

	"card_service/test/utils/reports"
	"context"
	"testing"
	"time"

	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/severity"
	"github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateTestCard(t *testing.T, c *container.Components, ctx context.Context, g gomega.Gomega) *cardservice.Card {
	t.Helper()
	req := newValidCreateReq()

	created, _, err := c.CardService.CreateCard(ctx, req)
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), "CreateCard should succeed with a valid request")
	g.Expect(created).ShouldNot(gomega.BeNil())

	// Basic field checks
	g.Expect(created.UserId).To(gomega.Equal(req.UserId))
	g.Expect(created.Operator.Id).To(gomega.Equal(req.OperatorId))
	g.Expect(created.CardNumber).To(gomega.Equal(req.CardNumber))
	g.Expect(created.IssueDate).To(gomega.Equal(req.IssueDate))
	g.Expect(created.ExpiryDate).To(gomega.Equal(req.ExpiryDate))
	return created
}

func newValidCreateReq() *cardservice.CreateCardRequest {
	return &cardservice.CreateCardRequest{
		UserId:     1,
		OperatorId: 1,
		CardNumber: solutions.GenerateRandomCardNumber(),
		IssueDate:  time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
		ExpiryDate: time.Now().AddDate(1, 0, 0).Format("2006-01-02"),
		Balance:    100,
	}
}

func TestCreateCard(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.CardServiceFeature, reports.CardSuite,
		allure.Severity(severity.Critical),
		allure.Tags(reports.CardTag),
		allure.Name("Create card"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)
			CreateTestCard(t, c, ctx, g)
		}),
	)
}

func TestGetCard(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.CardServiceFeature, reports.CardSuite,
		allure.Severity(severity.Critical),
		allure.Tags(reports.CardTag),
		allure.Name("Get card"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)
			created := CreateTestCard(t, c, ctx, g)

			response, err := c.CardService.GetCard(ctx, &cardservice.GetCardRequest{Id: created.Id})
			g.Expect(err).ShouldNot(gomega.HaveOccurred(), "GetCard should not return error")
			cardservicechecks.CheckCard(g, response, created)
		}),
	)
}

func TestDeleteCard(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.CardServiceFeature, reports.CardSuite,
		allure.Severity(severity.Critical),
		allure.Tags(reports.CardTag),
		allure.Name("Delete card"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)
			created := CreateTestCard(t, c, ctx, g)

			// Delete the card
			resp, err := c.CardService.DeleteCard(ctx, &cardservice.DeleteCardRequest{Id: created.Id})
			g.Expect(err).ShouldNot(gomega.HaveOccurred(), "DeleteCard should not return error")
			g.Expect(resp).ShouldNot(gomega.BeNil(), "DeleteCard response should not be nil")

			if resp != nil {
				g.Expect(resp.Success).To(gomega.BeTrue(), "DeleteCard success should be true")
			}

			_, err = c.CardService.GetCard(ctx, &cardservice.GetCardRequest{Id: created.Id})
			g.Expect(err).Should(gomega.HaveOccurred(), "Expected error for deleted card")
			st, _ := status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.NotFound), "Expected NotFound after delete")
		}),
	)
}

func TestCreateCardNegative(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.CardServiceFeature, reports.CardSuite,
		allure.Severity(severity.Normal),
		allure.Tags(reports.CardTag),
		allure.Name("Create card negative cases with error check"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)

			// Empty UserId
			_, _, err := c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     0,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 1,
				IssueDate:  "2025-01-01",
				ExpiryDate: "2026-01-01",
				Balance:    100,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ := status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("UserId must be provided"))

			// Uncorrect IssueDate
			_, _, err = c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 1,
				IssueDate:  "01-01-2025",
				ExpiryDate: "2026-01-01",
				Balance:    100,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("Invalid issue date format"))

			// ExpiryDate earlier than IssueDate
			_, _, err = c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 1,
				IssueDate:  "2026-01-01",
				ExpiryDate: "2025-01-01",
				Balance:    100,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("Expiry date must be after issue date"))

			// ExpiryDate == IssueDate
			_, _, err = c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 1,
				IssueDate:  "2025-01-01",
				ExpiryDate: "2025-01-01",
				Balance:    100,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("Expiry date must be after issue date"))

			// Empty CardNumber
			_, _, err = c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: "",
				OperatorId: 1,
				IssueDate:  "2025-01-01",
				ExpiryDate: "2026-01-01",
				Balance:    100,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("Invalid card number"))

			// CardNumber < 16 numbers
			_, _, err = c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: "12345",
				OperatorId: 1,
				IssueDate:  "2025-01-01",
				ExpiryDate: "2026-01-01",
				Balance:    100,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("Invalid card number"))

			// Empty OperatorId
			_, _, err = c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 0,
				IssueDate:  "2025-01-01",
				ExpiryDate: "2026-01-01",
				Balance:    100,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("OperatorId must be provided"))
		}),
	)
}

func TestCreateCardPositiveCases(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.CardServiceFeature, reports.CardSuite,
		allure.Severity(severity.Normal),
		allure.Tags(reports.CardTag),
		allure.Name("Create card positive cases"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)

			// Balance with decimals
			card1, _, err := c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 1,
				IssueDate:  "2025-01-01",
				ExpiryDate: "2026-01-01",
				Balance:    123.45,
			})
			g.Expect(err).ShouldNot(gomega.HaveOccurred())
			g.Expect(card1.Balance).To(gomega.Equal(float64(123.45)))

			// Check Balance precision
			card2, _, err := c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     2,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 2,
				IssueDate:  "2025-02-01",
				ExpiryDate: "2026-02-01",
				Balance:    500,
			})
			g.Expect(err).ShouldNot(gomega.HaveOccurred())
			g.Expect(card2.Operator).ShouldNot(gomega.BeNil())
			g.Expect(card2.Operator.Id).To(gomega.Equal(int32(2)))
		}),
	)
}

func TestGetCardNegative(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.CardServiceFeature, reports.CardSuite,
		allure.Severity(severity.Normal),
		allure.Tags(reports.CardTag),
		allure.Name("Get card negative cases"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)

			// Id = 0
			_, err := c.CardService.GetCard(ctx, &cardservice.GetCardRequest{Id: 0})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ := status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("Id must be provided"))

			// Unexisting Id
			_, err = c.CardService.GetCard(ctx, &cardservice.GetCardRequest{Id: 999999})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.NotFound))
			g.Expect(err.Error()).To(gomega.ContainSubstring("card with id 999999 not found"))
		}),
	)
}

func TestDeleteCardNegative(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.CardServiceFeature, reports.CardSuite,
		allure.Severity(severity.Normal),
		allure.Tags(reports.CardTag),
		allure.Name("Delete card negative cases"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)

			// Id = 0
			_, err := c.CardService.DeleteCard(ctx, &cardservice.DeleteCardRequest{Id: 0})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ := status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("Id must be provided"))

			// Unexisting Id
			_, err = c.CardService.DeleteCard(ctx, &cardservice.DeleteCardRequest{Id: 999999})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.NotFound))
			g.Expect(err.Error()).To(gomega.ContainSubstring("card with id 999999 not found"))
		}),
	)
}

func TestCreateCardNegativeBalance(t *testing.T) {
	t.Parallel()
	allure.Test(t, reports.CardServiceFeature, reports.CardSuite,
		allure.Severity(severity.Normal),
		allure.Tags(reports.CardTag),
		allure.Name("Create card with invalid balance"),
		allure.Action(func() {
			c, ctx, g := common.SetupTesting(t)

			// Balance < 0
			_, _, err := c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 1,
				IssueDate:  "2025-01-01",
				ExpiryDate: "2026-01-01",
				Balance:    -100,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ := status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
			g.Expect(err.Error()).To(gomega.ContainSubstring("Balance cannot be negative"))

			// Balance too large
			_, _, err = c.CardService.CreateCard(ctx, &cardservice.CreateCardRequest{
				UserId:     1,
				CardNumber: solutions.GenerateRandomCardNumber(),
				OperatorId: 1,
				IssueDate:  "2025-01-01",
				ExpiryDate: "2026-01-01",
				Balance:    1e15,
			})
			g.Expect(err).Should(gomega.HaveOccurred())
			st, _ = status.FromError(err)
			g.Expect(st.Code()).To(gomega.Equal(codes.InvalidArgument))
		}),
	)
}

package cardservicechecks

import (
	cardservice "card_service/generated/proto"
	solutions "card_service/test/utils/assertions/common"

	"github.com/dailymotion/allure-go"
	"github.com/onsi/gomega"
)

func CheckCard(g gomega.Gomega, actualCard, expectedCard *cardservice.Card) {
	allure.Step(allure.Description("Checking card"), allure.Action(func() {
		solutions.AssertToEqual(g, actualCard.UserId, expectedCard.UserId, "Card UserId")
		solutions.AssertToEqual(g, actualCard.CardNumber, expectedCard.CardNumber, "Card CardNumber")
		solutions.AssertToEqual(g, actualCard.Operator.Code, expectedCard.Operator.Code, "Card OperatorCode")
		solutions.AssertToEqual(g, actualCard.Operator.Name, expectedCard.Operator.Name, "Card OperatorName")
		solutions.AssertToEqual(g, actualCard.IssueDate, expectedCard.IssueDate, "Card IssueDate")
		solutions.AssertToEqual(g, actualCard.ExpiryDate, expectedCard.ExpiryDate, "Card ExpiryDate")
		solutions.AssertToEqual(g, actualCard.IsActive, expectedCard.IsActive, "Card IsActive")
		solutions.AssertToEqual(g, actualCard.Balance, expectedCard.Balance, "Card Balance")
	}))
}

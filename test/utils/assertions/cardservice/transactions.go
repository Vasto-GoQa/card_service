package cardservicechecks

import (
	cardservice "card_service/generated/proto"
	solutions "card_service/test/utils/assertions/common"

	"github.com/dailymotion/allure-go"
	"github.com/onsi/gomega"
)

func CheckTransaction(g gomega.Gomega, actual, expected *cardservice.Transaction) {
	allure.Step(allure.Description("Checking transaction"), allure.Action(func() {
		solutions.AssertToEqual(g, actual.Id, expected.Id, "Transaction Id")
		solutions.AssertToEqual(g, actual.FromCardId, expected.FromCardId, "Transaction FromCardId")
		solutions.AssertToEqual(g, actual.ToCardId, expected.ToCardId, "Transaction ToCardId")
		solutions.AssertToEqual(g, actual.Amount, expected.Amount, "Transaction Amount")
	}))
}

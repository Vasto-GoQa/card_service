package common

import (
	"runtime/debug"
	"testing"
	"time"

	"github.com/dailymotion/allure-go"
	"github.com/onsi/gomega"
	"github.com/pkg/errors"
)

func GetGomega(t *testing.T) gomega.Gomega {
	g := gomega.NewWithT(t)
	g.SetDefaultEventuallyTimeout(time.Second * 20)
	g.SetDefaultEventuallyPollingInterval(time.Second * 3)
	g.ConfigureWithFailHandler(func(message string, callerSkip ...int) {
		g.THelper()
		allure.Fail(errors.New(message))
		t.Fatalf("\n%s %s", message, debug.Stack())
	})
	g.THelper = t.Helper
	return g
}

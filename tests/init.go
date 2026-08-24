package tests

import (
	"log"
	"reflect"
	"testing"

	mock "github.com/sanchaimac/go-amnet/tests/mocks"
)

type Mock struct {
	FundConnext mock.MockFundConnext
}

type Calls [][]interface{}
type AssertCalled struct {
	FnName string
	IType  reflect.Type
	Calls  Calls
}

type Expected struct {
	Error     error
	Called    []AssertCalled
	NotCalled []AssertCalled
}

type Testcase map[string]struct {
	Mock     *Mock
	Input    interface{}
	Expected interface{}
}

type SetupOptions struct {
	Mock *Mock
}

type TestTools struct {
	Teardown func(tb testing.TB)
	FC       *mock.FundConnext
}

func SetupSuite(tb testing.TB, opts SetupOptions) TestTools {
	log.Println("setup suite")

	// Setup testsuite base data

	return TestTools{
		Teardown: func(tb testing.TB) {
			log.Println("teardown suite")
		},
	}
}

func SetupTest(tb testing.TB, opts SetupOptions) TestTools {
	log.Println("setup test")
	// Setup testcase
	m := opts.Mock

	return TestTools{
		Teardown: func(tb testing.TB) {
			log.Println("teardown test")
		},
		FC: mock.NewFundConnext(m.FundConnext),
	}
}

// Assertion
// func ExpectedCalls(t *testing.T, called []AssertCalled) {
// 	// Asset called.
// 	for _, v := range called {
// 		for _, call := range v.Calls {
// 			v.IType.(*mock.FundConnext).Mock.AssertCalled(t, v.FnName, call...)
// 		}
// 	}
// }

// func ExpectedNotCalls(t *testing.T, c *container_mock.TestContainer, notCalled []AssertCalled) {
// 	// Assert not called.
// 	c.RequireInvoke(func(
// 		or repository.IOrderRepository,
// 	) {
// 		for _, v := range notCalled {
// 			switch v.IType {
// 			case reflect.TypeOf(or):
// 				for _, call := range v.Calls {
// 					or.(*repository_mock.OrderRepository).Mock.AssertNotCalled(t, v.FnName, call...)
// 				}
// 			default:
// 				t.Fatalf("assertCalled no matched type: %s", v.IType.String())
// 			}
// 		}
// 	})
// }

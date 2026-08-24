package tests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sanchaimac/go-amnet"
	mock "github.com/sanchaimac/go-amnet/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRedemption(t *testing.T) {
	testTools := SetupSuite(t, SetupOptions{})
	defer testTools.Teardown(t)

	type expected struct {
		Expected
	}
	suite := Testcase{
		"001.Create LTF Redemption E274": {
			Mock: &Mock{
				FundConnext: mock.MockFundConnext{
					APICall: &mock.ExpectedAPICall{
						Error: func() error {
							errResp := `{"errMsg":{"code":"E274","message":"This fund does not support advance order. Please change effective date."}}`
							reader := io.NopCloser(strings.NewReader(errResp))
							respBody, err := ioutil.ReadAll(reader)
							if err != nil {
								t.Fatal(err)
							}
							var errMsg fundconnext.FCError
							if err := json.Unmarshal(respBody, &errMsg); err != nil {
								t.Fatal(err)
							}
							return &errMsg
						}(),
					},
				},
			},
			Input: fundconnext.RedemptionOrder{},
			Expected: expected{
				Expected: Expected{
					Error: fundconnext.UnSupportAdvanceOrder,
				},
			},
		},
	}

	for name, tc := range suite {
		t.Run(name, func(t *testing.T) {
			testTools := SetupTest(t, SetupOptions{Mock: tc.Mock})
			defer testTools.Teardown(t)
			input := tc.Input.(fundconnext.RedemptionOrder)
			_, err := testTools.FC.CreateRedemption(input)
			expected := tc.Expected.(expected)
			if expected.Error != nil {
				require.Error(t, err)
				assert.EqualError(t, err, expected.Error.Error())
			}
		})
	}
}

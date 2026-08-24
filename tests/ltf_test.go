package tests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/codefin-stack/go-fundconnext"
	mock "github.com/codefin-stack/go-fundconnext/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLTFRedeemableUnitInquiry(t *testing.T) {
	testTools := SetupSuite(t, SetupOptions{})
	defer testTools.Teardown(t)
	type input struct {
		unitholderId string
		fundCode     string
	}
	type expected struct {
		Expected
	}
	suite := Testcase{
		"3. LTF unit not found should return error not found": {
			Mock: &Mock{
				FundConnext: mock.MockFundConnext{
					APICall: &mock.ExpectedAPICall{
						Error: func() error {
							errResp := `{"errMsg":{"code":"E339","message":"LTF Balance Not Found"}}`
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
			Input: input{
				unitholderId: "user1",
				fundCode:     "ASP-GLTF-T",
			},
			Expected: expected{
				Expected: Expected{
					Error: fundconnext.LtfBalanceNotFound,
				},
			},
		},
	}

	for name, tc := range suite {
		t.Run(name, func(t *testing.T) {
			testTools := SetupTest(t, SetupOptions{Mock: tc.Mock})
			defer testTools.Teardown(t)

			input := tc.Input.(input)
			_, err := testTools.FC.LTFRedeemableUnitInquiry(input.unitholderId, input.fundCode)
			expected := tc.Expected.(expected)
			if expected.Error != nil {
				require.Error(t, err)
				assert.EqualError(t, err, expected.Error.Error())
			}
		})
	}
}

package tests

import (
	"testing"
)

func TestAccount(t *testing.T) {
	fc, err := NewFundConnext()
	if err != nil {
		t.Error(err)
	}
	profile, err := fc.RetrieveIndividualCustomerProfileAndAccount("1100701324225", "") // 1100701324225
	if err != nil {
		t.Error(err)
	}
	t.Log(profile.CddScore)
}

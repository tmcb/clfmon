package routines

import (
	"testing"
)

func TestAlertLogic(t *testing.T) {
	if !(checkAlert(false, 0, 1) == true) {
		t.Error()
	}
	if !(checkAlert(false, 1, 1) == false) {
		t.Error()
	}
	if !(checkAlert(false, 2, 1) == false) {
		t.Error()
	}
	if !(checkAlert(true, 0, 1) == true) {
		t.Error()
	}
	if !(checkAlert(true, 1, 1) == true) {
		t.Error()
	}
	if !(checkAlert(true, 2, 1) == false) {
		t.Error()
	}
}

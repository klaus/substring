package substring

import "testing"

var expr string = "The black cat climbed the green tree"

func TestColor(t *testing.T) {

	expect := "black"
	got := Substr(expr, 4, 5)
	if got != expect {
		t.Errorf("%q is not %q", got, expect)
	}

}

func TestMiddle(t *testing.T) {
	got := Substr(expr, 4, -11)

	expect := "black cat climbed the"
	if got != expect {
		t.Errorf("%q is not %q", got, expect)
	}

}
func TestEnd(t *testing.T) {
	got := Substr(expr, 14)

	expect := "climbed the green tree"
	if got != expect {
		t.Errorf("%q is not %q", got, expect)
	}

}

func TestTail(t *testing.T) {
	got := Substr(expr, -4)

	expect := "tree"
	if got != expect {
		t.Errorf("%q is not %q", got, expect)
	}

}

func TestZ(t *testing.T) {
	got := Substr(expr, -4, 2)

	expect := "tr"
	if got != expect {
		t.Errorf("%q is not %q", got, expect)
	}

}

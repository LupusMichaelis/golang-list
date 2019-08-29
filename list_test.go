package main

import (
	"testing"
)

func TestList(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{ []string{}, "List []" },
		{ []string{"first"}, "List ['first']" },
		{ []string{"first", "second"}, "List ['second','first']" },
		{ []string{"first", "second", "third"}, "List ['third','second','first']" },
	}

	for _, esac := range cases {
		var list = List{nil}
		list.AddMany(esac.in...)

		var got = list.String()

		if got != esac.want {
			t.Errorf("Add elements of (%q) == %q, want %q", esac.in, got, esac.want)
		}
	}

}

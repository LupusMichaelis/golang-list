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

func TestReverse(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{ []string{}, "List []" },
		{ []string{"first"}, "List ['first']" },
		{ []string{"first", "second"}, "List ['first','second']" },
		{ []string{"first", "second", "third"}, "List ['first','second','third']" },
	}

	for _, esac := range cases {
		var list = List{nil}
		list.AddMany(esac.in...)
		list.Reverse()

		var got = list.String()

		if got != esac.want {
			t.Errorf("Reverse(%q) == %q, want %q", esac.in, got, esac.want)
		}
	}

}

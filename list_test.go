package main

import (
	"testing"
)

func TestList(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{}, "List []"},
		{[]string{"first"}, "List ['first']"},
		{[]string{"first", "second"}, "List ['second','first']"},
		{[]string{"first", "second", "third"}, "List ['third','second','first']"},
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
		in   []string
		want string
	}{
		{[]string{}, "List []"},
		{[]string{"first"}, "List ['first']"},
		{[]string{"first", "second"}, "List ['first','second']"},
		{[]string{"first", "second", "third"}, "List ['first','second','third']"},
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

func TestReverseAndPop(t *testing.T) {
	cases := []struct {
		in     []string
		want   string
		popped string
		err    error
	}{
		{[]string{}, "List []", "", CherryNotPopped{}},
		{[]string{"cherry"}, "List []", "cherry", nil},
		{[]string{"cherry", "plum"}, "List ['plum']", "cherry", nil},
		{[]string{"cherry", "apricot", "pear"}, "List ['apricot','pear']", "cherry", nil},
	}

	for _, esac := range cases {
		var list = List{nil}
		list.AddMany(esac.in...)
		list.Reverse()

		gotPopped, err := list.Pop()
		if err != esac.err {
			t.Errorf("Pop(%q) got %q, want %q", esac.in, err, esac.err)
		}

		if gotPopped != esac.popped {
			t.Errorf("Pop(%q) got %q, want %q", esac.in, gotPopped, esac.popped)
		}

		var got = list.String()
		if got != esac.want {
			t.Errorf("PopAndReverse(%q) == %q, want %q", esac.in, got, esac.want)
		}
	}

}

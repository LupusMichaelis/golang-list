package main

import (
	"testing"
)

func TestList(t *testing.T) {
	cases := []struct {
		candidate []string
		expected  string
	}{
		{[]string{}, "List []"},
		{[]string{"first"}, "List ['first']"},
		{[]string{"first", "second"}, "List ['second','first']"},
		{[]string{"first", "second", "third"}, "List ['third','second','first']"},
	}

	for _, esac := range cases {
		var list = List{nil}
		list.AddMany(esac.candidate...)

		var got = list.String()

		if got != esac.expected {
			t.Errorf("Add elements of (%q) == %q, expected %q", esac.candidate, got, esac.expected)
		}
	}

}

func TestReverse(t *testing.T) {
	cases := []struct {
		candidate []string
		expected  string
	}{
		{[]string{}, "List []"},
		{[]string{"first"}, "List ['first']"},
		{[]string{"first", "second"}, "List ['first','second']"},
		{[]string{"first", "second", "third"}, "List ['first','second','third']"},
	}

	for _, esac := range cases {
		var list = List{nil}
		list.AddMany(esac.candidate...)
		list.Reverse()

		var got = list.String()

		if got != esac.expected {
			t.Errorf("Reverse(%q) == %q, expected %q", esac.candidate, got, esac.expected)
		}
	}

}

func TestReverseAndPop(t *testing.T) {
	cases := []struct {
		candidate       []string
		expected        string
		expectedPayload string
		expectedErr     error
	}{
		{[]string{}, "List []", "", CherryNotPopped{}},
		{[]string{"cherry"}, "List []", "cherry", nil},
		{[]string{"cherry", "plum"}, "List ['plum']", "cherry", nil},
		{[]string{"cherry", "apricot", "pear"}, "List ['apricot','pear']", "cherry", nil},
	}

	for _, esac := range cases {
		var list = List{nil}
		list.AddMany(esac.candidate...)
		list.Reverse()

		gotPopped, err := list.Pop()
		if err != esac.expectedErr {
			t.Errorf("Pop(%q) got %q, expected %q", esac.candidate, err, esac.expectedErr)
		}

		if gotPopped != esac.expectedPayload {
			t.Errorf("Pop(%q) got %q, expected %q", esac.candidate, gotPopped, esac.expectedPayload)
		}

		var got = list.String()
		if got != esac.expected {
			t.Errorf("PopAndReverse(%q) == %q, expected %q", esac.candidate, got, esac.expected)
		}
	}

}

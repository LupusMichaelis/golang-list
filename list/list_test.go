package list_test

import (
	behave "github.com/onsi/ginkgo"
	assert "github.com/onsi/gomega"
	listModule "lupusmic.org/golang-list/list"
	"testing"
)

func TestList(t *testing.T) {
	assert.RegisterFailHandler(behave.Fail)
	behave.RunSpecs(t, "Calculator Suite")
}

var _ = behave.Describe("List", func() {
	behave.Context("List building", func() {
		behave.It(
			`
			must print 'List' followed by a space, a bracket surrounded list of element
			surrounded with quotes and separated by commas
			`, func() {
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
					var candidate = listModule.Build(esac.candidate...)

					assert.Expect(candidate.String()).Should(assert.Equal(esac.expected))
				}
			})
	})

	behave.Context("List reversing", func() {
		behave.It(
			`
			must print a reversed list
			`, func() {
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
					var candidate = listModule.Build(esac.candidate...)
					candidate.Reverse()

					assert.Expect(candidate.String()).Should(assert.Equal(esac.expected))
				}
			})
	})

	behave.Context("Reverse list and pop", func() {
		behave.It(
			`
			must print a reversed list minus the first inserted element
			`, func() {
				cases := []struct {
					candidate       []string
					expected        string
					expectedPayload string
					expectedErr     error
				}{
					{[]string{}, "List []", "", listModule.CherryNotPopped{}},
					{[]string{"cherry"}, "List []", "cherry", nil},
					{[]string{"cherry", "plum"}, "List ['plum']", "cherry", nil},
					{[]string{"cherry", "apricot", "pear"}, "List ['apricot','pear']", "cherry", nil},
				}

				for _, esac := range cases {

					var candidate = listModule.List{}
					candidate.AddMany(esac.candidate...)
					candidate.Reverse()

					gotPopped, err := candidate.Pop()
					if nil != esac.expectedErr {
						assert.Expect(err).Should(assert.Equal(esac.expectedErr))
					} else {
						assert.Expect(err).Should(assert.BeNil())
					}

					assert.Expect(gotPopped).Should(assert.Equal(esac.expectedPayload))
					assert.Expect(candidate.String()).Should(assert.Equal(esac.expected))
				}
			})
	})
})

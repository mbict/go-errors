package errors

import (
	"errors"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

var (
	errTest1 = errors.New("err 1")
	errTest2 = errors.New("err 2")
)

type ErrorSuite struct{}

var _ = Suite(&ErrorSuite{})

func (es *ErrorSuite) TestToString(c *C) {
	errs := Errors{errTest1}
	c.Assert(errs.Error(), Matches, "err 1")

	errs = Errors{errTest1, errTest2}
	c.Assert(errs.Error(), Matches, "err 1, err 2")
}

func (es *ErrorSuite) TestNoErrorsToEmptyString(c *C) {
	errs := Errors{}

	c.Assert(errs, ErrorMatches, "")
}

func (es *ErrorSuite) TestAdd(c *C) {
	errs := Errors{}

	errs.Add(errTest1)
	errs.Add(errTest2)

	c.Assert(errs, HasLen, 2)
	c.Assert(errs, DeepEquals, Errors{errTest1, errTest2})
}

func (es *ErrorSuite) TestHasString(c *C) {
	errs := Errors{errTest2}

	c.Assert(errs.HasString(errTest1.Error()), Equals, false)
	c.Assert(errs.HasString(errTest2.Error()), Equals, true)
}

func (es *ErrorSuite) TestHasError(c *C) {
	errs := Errors{errTest2}

	c.Assert(errs.HasError(errTest1), Equals, false)
	c.Assert(errs.HasError(errTest2), Equals, true)
}

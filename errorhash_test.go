package errors

import (
	. "gopkg.in/check.v1"
)

type ErrorHashSuite struct{}

var es = Suite(&ErrorHashSuite{})

func (es *ErrorHashSuite) TestToString(c *C) {
	errs := NewErrorHash()
	errs["A"] = Errors{errTest1, errTest2}
	errs["B"] = Errors{errTest2}

	c.Assert(errs.Error(), Equals, "A:[err 1, err 2], B:[err 2]")
}

func (es *ErrorHashSuite) TestNoErrorsToEmptyString(c *C) {
	errs := NewErrorHash()

	c.Assert(errs, ErrorMatches, "")
}

func (es *ErrorHashSuite) TestAdd(c *C) {
	errs := NewErrorHash()

	errs.Add("A", errTest1)
	errs.Add("B", errTest1)
	errs.Add("A", errTest2)

	c.Assert(errs, HasLen, 2)
	c.Assert(errs["A"], HasLen, 2)
	c.Assert(errs["A"], DeepEquals, Errors{errTest1, errTest2})
	c.Assert(errs["B"], HasLen, 1)
	c.Assert(errs["B"], DeepEquals, Errors{errTest1})
}

func (es *ErrorHashSuite) TestHas(c *C) {
	errs := NewErrorHash()
	errs["A"] = Errors{errTest2}

	c.Assert(errs.Has("B"), Equals, false)
	c.Assert(errs.Has("A"), Equals, true)
}

func (es *ErrorHashSuite) TestHasString(c *C) {
	errs := NewErrorHash()
	errs["A"] = Errors{errTest2}

	c.Assert(errs.HasString("B", errTest2.Error()), Equals, false)
	c.Assert(errs.HasString("A", errTest1.Error()), Equals, false)
	c.Assert(errs.HasString("A", errTest2.Error()), Equals, true)
}

func (es *ErrorHashSuite) TestHasError(c *C) {
	errs := NewErrorHash()
	errs["A"] = Errors{errTest2}

	c.Assert(errs.HasError("B", errTest2), Equals, false)
	c.Assert(errs.HasError("A", errTest1), Equals, false)
	c.Assert(errs.HasError("A", errTest2), Equals, true)
}

package slicewithminimalaverage

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSliceWithMinimalAverage(t *testing.T)  {
	Convey("Given I have an array", t, func() {
		A := []int{4,2,2,5,1,5,8}
		Convey("When calculate slice with minimal average", func() {
			res := SliceWithMinimalAverage(A)
			Convey("Result should be 1", func() {
				So(res, ShouldEqual, 1)
			})
		})
	})
	Convey("Given I have 2 element array", t, func() {
		A := []int{2,2}
		Convey("When calculate slice with minimal average", func() {
			res := SliceWithMinimalAverage(A)
			Convey("Result should be 1", func() {
				So(res, ShouldEqual, 0)
			})
		})
	})
	Convey("Given I have 3 element array", t, func() {
		A := []int{2,1,1}
		Convey("When calculate slice with minimal average", func() {
			res := SliceWithMinimalAverage(A)
			Convey("Result should be 1", func() {
				So(res, ShouldEqual, 1)
			})
		})
	})

}
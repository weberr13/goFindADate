package errorable

import (
	"errors"
	log "github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var anErr = errors.New("an Error")

func GetOne(string) (int, error) {
	return 1, nil
}
func GetErr(string) (int, error) {
	return 1, anErr
}
func ErrWhenA(s string) (int, error) {
	if s == "a" {
		return 1, anErr
	}
	return 1, nil
}
func TestErrorable(t *testing.T) {

	defer log.Flush()

	Convey("Make single IntErr", t, func() {
		i := NewIntErr(GetOne, "1")
		So(i.I, ShouldEqual, 1)
		So(i.E, ShouldBeNil)
	})
	Convey("Make single IntErr, fails", t, func() {
		i := NewIntErr(GetErr, "1")
		So(i.E, ShouldEqual, anErr)
	})
	Convey("Make IntErrs", t, func() {
		is := NewIntErrs(GetOne, "1", "1", "1")
		So(is.Len(), ShouldEqual, 3)
		for i, v := range *is {
			So(v.I, ShouldEqual, 1)
			So(v.E, ShouldBeNil)
			So(is.Get(i), ShouldEqual, 1)
		}
		So(is.GetFirstErr(), ShouldBeNil)
	})
	Convey("Make IntErrs, all fail", t, func() {
		is := NewIntErrs(ErrWhenA, "a", "a", "a")
		for _, v := range *is {
			So(v.E, ShouldEqual, anErr)
		}
		So(is.GetFirstErr(), ShouldEqual, anErr)
	})
	Convey("Make IntErrs, one fail", t, func() {
		is := NewIntErrs(ErrWhenA, "1", "a", "1")
		So(is.GetFirstErr(), ShouldEqual, anErr)
	})
}

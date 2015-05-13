package gofindadate

import (
	log "github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestYearDate(t *testing.T) {

	defer log.Flush()

	Convey("regex compiled", t, func() {
		So(initErr, ShouldBeNil)
	})
	Convey("test happy slice path", t, func() {

		year, err := NewYearDate(GetDateWithYYYYMMDD, "2999-10-31", "2999", "10", "31")
		So(err, ShouldBeNil)
		So(year, ShouldNotBeNil)
		So(year.year, ShouldEqual, 2999)
		So(year.month, ShouldEqual, 10)
		So(year.day, ShouldEqual, 31)

	})
	Convey("Extract the oldest from an empty slice", t, func() {
		emptySlice := make([]string, 0)
		So(GetOldestOfSlice(emptySlice), ShouldEqual, "")
	})
	Convey("Stract the oldest from a single slice", t, func() {
		singleSlice := []string{"something"}
		So(GetOldestOfSlice(singleSlice), ShouldEqual, "something")
	})
	Convey("Stract the oldest from a single slice without dates", t, func() {
		singleSlice := []string{"something", "somethingelse"}
		So(GetOldestOfSlice(singleSlice), ShouldEqual, "something")
	})
	Convey("Stract the oldest from a single slice without only one date", t, func() {
		singleSlice := []string{"something", "2004-01-01"}
		So(GetOldestOfSlice(singleSlice), ShouldEqual, "something")
	})
	Convey("Stract the oldest from an ordered slice", t, func() {
		singleSlice := []string{"2004-01-01", "2005-01-01", "2005-01-02"}
		So(GetOldestOfSlice(singleSlice), ShouldEqual, "2004-01-01")
	})
	Convey("Stract the oldest from an reverse ordered slice", t, func() {
		singleSlice := []string{"2005-01-02", "2005-01-01", "2004-01-01"}
		So(GetOldestOfSlice(singleSlice), ShouldEqual, "2004-01-01")
	})
	Convey("Stract the oldest from an unordered slice", t, func() {
		singleSlice := []string{"2005-01-01", "2005-01-02", "2004-01-01"}
		So(GetOldestOfSlice(singleSlice), ShouldEqual, "2004-01-01")
	})
	Convey("Stract the oldest from an corrupted slice", t, func() {
		singleSlice := []string{"2005-01-01", "-01-02", "2004-01-01"}
		So(GetOldestOfSlice(singleSlice), ShouldEqual, "-01-02")
	})
	Convey("Stract the oldest from an corrupted slice, already in order", t, func() {
		singleSlice := []string{"-01-02", "2004-01-01"}
		So(GetOldestOfSlice(singleSlice), ShouldEqual, "-01-02")
	})
	Convey("test happy string path", t, func() {

		year, err := NewYearDate(GetDateWithYYYYMMDD, "stuff2999-10-31morestuff")
		So(err, ShouldBeNil)
		So(year, ShouldNotBeNil)
		So(year.year, ShouldEqual, 2999)
		So(year.month, ShouldEqual, 10)
		So(year.day, ShouldEqual, 31)

	})
	Convey("Year Less", t, func() {

		year_less, _ := NewYearDate(GetDateWithYYYYMMDD, "stuff2999-10-31morestuff")
		year_great, _ := NewYearDate(GetDateWithYYYYMMDD, "stuff3000-10-31morestuff")
		So(year_less.Less(year_great), ShouldBeTrue)

	})
	Convey("Month Less", t, func() {

		year_less, _ := NewYearDate(GetDateWithYYYYMMDD, "stuff2999-10-31morestuff")
		year_great, _ := NewYearDate(GetDateWithYYYYMMDD, "stuff2999-11-31morestuff")
		So(year_less.Less(year_great), ShouldBeTrue)

	})
	Convey("Day Less", t, func() {

		year_less, _ := NewYearDate(GetDateWithYYYYMMDD, "stuff2999-10-30morestuff")
		year_great, _ := NewYearDate(GetDateWithYYYYMMDD, "stuff2999-10-31morestuff")
		So(year_less.Less(year_great), ShouldBeTrue)

	})
	// This will cause badly named indexes to be deleted FIRST
	Convey("Nil less", t, func() {
		year_less, _ := NewYearDate(GetDateWithYYYYMMDD, "stuff2999-10-30morestuff")
		var nullcheck *YearDate
		nullcheck = nil
		less, ok := year_less.InvalidIsLess(nil)
		So(less, ShouldBeFalse)
		So(ok, ShouldBeFalse)
		less, ok = nullcheck.InvalidIsLess(year_less)
		So(less, ShouldBeTrue)
		So(ok, ShouldBeFalse)
		less, ok = year_less.InvalidIsLess(year_less)
		So(less, ShouldBeFalse)
		So(ok, ShouldBeTrue)
	})
	Convey("test no date string", t, func() {

		year, err := NewYearDate(GetDateWithYYYYMMDD, "noyear-10-31morestuff")
		So(err, ShouldEqual, InvalidDateString)
		So(year, ShouldBeNil)

	})
	Convey("test empty args", t, func() {

		year, err := NewYearDate(GetDateWithYYYYMMDD)
		So(err, ShouldEqual, NoInput)
		So(year, ShouldBeNil)

	})
	Convey("test too many args", t, func() {

		year, err := NewYearDate(GetDateWithYYYYMMDD, "1", "2", "3")
		So(err, ShouldEqual, InvalidDateLength)
		So(year, ShouldBeNil)

	})
	Convey("test bad year", t, func() {

		year, err := NewYearDate(GetDateWithYYYYMMDD, "", "a", "2", "3")
		So(err, ShouldNotBeNil)
		So(year, ShouldBeNil)

	})
	Convey("test bad month", t, func() {

		year, err := NewYearDate(GetDateWithYYYYMMDD, "", "1", "a", "3")
		So(err, ShouldNotBeNil)
		So(year, ShouldBeNil)

	})
	Convey("test bad day", t, func() {

		year, err := NewYearDate(GetDateWithYYYYMMDD, "", "1", "2", "a")
		So(err, ShouldNotBeNil)
		So(year, ShouldBeNil)

	})
}

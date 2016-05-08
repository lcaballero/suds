package solve

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)


func TestTile(t *testing.T) {

	Convey("A new tile shouldn't have a value", t, func() {
		b := NewTile(0)
		So(b.HasValue(), ShouldBeFalse)
	})

	Convey("A new tile should have 0 excludes", t, func() {
		b := NewTile(0)
		So(b.Excluded, ShouldNotBeNil)
		So(len(b.Excluded), ShouldEqual, 0)
	})
}

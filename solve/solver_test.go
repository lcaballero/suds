package solve

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBoard(t *testing.T) {

	Convey("Board-2 should have < 81 plays left", t, func() {
		b,err := board2()
		So(err, ShouldBeNil)
		So(b.PlaysLeft(), ShouldBeLessThan, ROWS*COLS)
	})

	Convey("A new empty board should have 81 plays left", t, func() {
		b,err := board1()
		So(err, ShouldBeNil)
		So(b.PlaysLeft(), ShouldEqual, ROWS*COLS)
	})

	Convey("Board 1 should correctly craete a board", t, func() {
		b,err := board1()
		So(err, ShouldBeNil)
		So(len(b.Tiles), ShouldEqual, 9)
		for _,row := range b.Tiles {
			So(len(row), ShouldEqual, 9)
		}
	})

	Convey("A new empty board should have 81 plays left", t, func() {
		b := NewEmptyBoard()
		So(b.PlaysLeft(), ShouldEqual, ROWS*COLS)
	})

	Convey("Newly created board should have 9x9 grid", t, func() {
		b := NewEmptyBoard()
		So(len(b.Tiles), ShouldEqual, 9)
		for _,row := range b.Tiles {
			So(len(row), ShouldEqual, 9)
		}
	})
}


package mock

import (
	"database/sql"
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSelectAll(t *testing.T) {
	convey.Convey("TestSelectAll before mock", t, func() {
		path := strings.Join([]string{"root:pass1234@tcp(", "localhost:3306", ")/", "ktable"}, "")
		db, err := sql.Open("mysql", path)
		convey.So(err, convey.ShouldBeNil)
		defer db.Close()

		rows := SelectAll(db)
		t.Log(rows)
		convey.So(rows, convey.ShouldNotBeNil)
	})

	convey.Convey("TestSelectAll after mock", t, func() {
		db, mock, err := sqlmock.New()
		defer db.Close()

		convey.So(err, convey.ShouldBeNil)

		//mock.ExpectBegin()
		mock.ExpectQuery("select").WillReturnRows(sqlmock.NewRows([]string{"col1", "col2"}))
		rows := SelectAll(db)
		t.Log(rows)
		convey.So(rows, convey.ShouldNotBeNil)
		//mock.ExpectCommit()
	})
}

func TestUpdateAll(t *testing.T) {
	convey.Convey("before mock", t, func() {
		path := strings.Join([]string{"root:pass1234@tcp(", "localhost:3306", ")/", "ktable"}, "")
		db, err := sql.Open("mysql", path)
		convey.So(err, convey.ShouldBeNil)
		// 没有被mock前获取一遍看看
		ret := UpdateAll(db)
		t.Log(ret)

		convey.So(ret, convey.ShouldNotBeNil)
	})

	convey.Convey("after mock", t, func() {
		db, mock, err := sqlmock.New()
		convey.So(err, convey.ShouldBeNil)
		defer db.Close()

		//mock.ExpectBegin()
		mock.ExpectExec("update ktable_cluster").WillReturnResult(sqlmock.NewResult(1, 100))
		ret := UpdateAll(db)
		t.Log(ret)
		cnt, err := ret.RowsAffected()
		convey.So(err, convey.ShouldBeNil)
		convey.So(cnt, convey.ShouldEqual, 100)
		//mock.ExpectCommit()
	})
}

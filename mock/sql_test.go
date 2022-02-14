package mock

import (
	"database/sql"
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSelectAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	convey.Convey("not err", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})

	mock.ExpectBegin()
	mock.ExpectQuery("select").WillReturnRows(sqlmock.NewRows([]string{"col1", "col2"}))
	SelectAll(db)
	mock.ExpectCommit()
}

func TestUpdateAll(t *testing.T) {
	path := strings.Join([]string{"root:pass1234@tcp(", "localhost:3306", ")/", "ktable"}, "")
	db, err := sql.Open("mysql", path)
	convey.Convey("should be nil", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
	UpdateAll(db)

	db, mock, err := sqlmock.New()
	defer db.Close()

	convey.Convey("not err", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})

	mock.ExpectBegin()
	mock.ExpectExec("update ktable_cluster").WillReturnResult(sqlmock.NewResult(1, 1))
	UpdateAll(db)
	mock.ExpectCommit()
}

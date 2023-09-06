package mock_test

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"testing"
	"weicai.zhao.io/mock"
)

func doSql(db *gorm.DB, id int64) (int64, error) {
	var item = struct {
		ID int64 `json:"id" gorm:"column:id;"`
	}{}
	err := db.Model(&item).Table("t_testing").Where("id = ?", id).First(&item).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// SELECT `id` FROM `t_testing` WHERE id = 10 ORDER BY `t_testing`.`id` LIMIT 1
// SELECT * FROM `t_testing` WHERE id = ? ORDER BY `t_testing`.`id` LIMIT 1
func TestNewGormMock(t *testing.T) {
	var (
		id      int64 = 10
		manager       = mock.NewGormMock()
	)

	convey.Convey("find one record", t, func() {
		rows := manager.Mock().NewRows([]string{"id"}).AddRow(id)
		manager.Mock().
			ExpectQuery("SELECT * FROM `t_testing` WHERE id = ? ORDER BY `t_testing`.`id` LIMIT 1").
			WithArgs(id).
			WillReturnRows(rows)

		_id, err := doSql(manager.DB().Debug(), id)

		convey.So(err, convey.ShouldBeNil)
		convey.So(id, convey.ShouldEqual, _id)

		if err := manager.Mock().ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	convey.Convey("record not found", t, func() {
		rows := manager.Mock().NewRows([]string{"id"}).RowError(1, gorm.ErrRecordNotFound)
		manager.Mock().
			ExpectQuery("SELECT * FROM `t_testing` WHERE id = ? ORDER BY `t_testing`.`id` LIMIT 1").
			WithArgs(id).
			WillReturnRows(rows)

		_id, err := doSql(manager.DB().Debug(), id)

		fmt.Println(err)

		convey.So(err, convey.ShouldBeError)
		convey.So(_id, convey.ShouldEqual, 0)

		if err := manager.Mock().ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

}

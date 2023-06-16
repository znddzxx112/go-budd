package server

import "github.com/DATA-DOG/go-sqlmock"

func (dsm *defaultServerMock) userMock() {
	dsm.UserLoginMock()
}

func (dsm *defaultServerMock) UserLoginMock() {

	rows0 := sqlmock.NewRows([]string{"id", "account", "true_name", "mobile", "status"})
	dsm.sqlMock.ExpectQuery("SELECT * FROM `user`  WHERE (mobile = ?) AND (status = ?) ORDER BY `user`.`id` ASC LIMIT 1").WithArgs("18800011121", 1).WillReturnRows(rows0)

	rows1 := sqlmock.NewRows([]string{"id", "account", "true_name", "mobile", "status"}).
		AddRow(10001, "budd", "gobudd", "18800011122", 1)
	dsm.sqlMock.ExpectQuery("SELECT * FROM `user`  WHERE (mobile = ?) AND (status = ?) ORDER BY `user`.`id` ASC LIMIT 1").WithArgs("18800011122", 1).WillReturnRows(rows1)

}

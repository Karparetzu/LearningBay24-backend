package course

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"regexp"
	"testing"
)

func TestSqlBoilerWithMock(t *testing.T) {
	// Mock DB instance by sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

	// Inject mock instance into boil.
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)

	// Create mock data with specific columns
	mockShipperRows := sqlmock.NewRows([]string{"id", "column_name"}).AddRow(1, 123)

	// Use regexp to avoid comparing query failed
	mockQueryFindShipper := regexp.QuoteMeta("SELECT * FROM `objects` WHERE (column_name=?) AND (`objects`.deleted_at is null) LIMIT 1;")

	// Mock a data
	mock.ExpectQuery(mockQueryFindShipper).WithArgs(123).
		WillReturnRows(mockShipperRows).
		RowsWillBeClosed()

	// Mock an error
	mock.ExpectQuery(mockQueryFindShipper).WithArgs(124).
		WillReturnError(sql.ErrNoRows).
		RowsWillBeClosed()
}

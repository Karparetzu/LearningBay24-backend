package course

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"learningbay24.de/backend/models"
	"regexp"
	"testing"
)

var c = &models.Course{
	ID:        1,
	Name:      "Testname",
	EnrollKey: "12345",
	ForumID:   1,
}

func TestGetCourse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

	oldDB := boil.GetDB()
	ctrl := &PublicController{db}

	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)

	rows := sqlmock.NewRows([]string{"ID", "Name", "Description", "EnrollKey", "ForumID"}).AddRow(c.ID, c.Name, null.String{}, c.EnrollKey, c.ForumID)
	query := regexp.QuoteMeta("select * from `course` where `id`=? and `deleted_at` is null")
	mock.ExpectQuery(query).WithArgs(c.ID).WillReturnRows(rows).RowsWillBeClosed()

	course, err := ctrl.GetCourse(c.ID)
	assert.NotNil(t, course)
	assert.NoError(t, err)
}

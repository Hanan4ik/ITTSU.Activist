package access

import (
	"activist/constants"
	"os"
	"slices"
	"strconv"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestInsertRead(t *testing.T) {
	sqlDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("CANNOT OPEN DB")
	}

	db := NewAcessDB(sqlDB)
	err = db.InitDB()
	if err != nil {
		t.Fatalf("CANNOT INIT DB")
	}

	for i := range uint(10) {
		t.Run("test"+strconv.FormatInt(int64(i), 10), func(t *testing.T) {
			right := Right{}
			right.UserID = i
			right.Rights = i
			err := db.AddRight(right)
			if err != nil {
				t.Error("ERROR INSERTING RIGHT")
			}
			readRight, err := db.GetUserRight(i)
			condition := readRight.UserID == i && readRight.Rights == i
			if !condition || err != nil {
				t.Error("ERROR READING RIGHT")
			}
			t.Log(t.Name() + " passed succesfully")
		})
	}

	os.Remove("test.db")

}

func TestCompareFunctions(t *testing.T) {
	tests := []struct {
		name   string
		userID uint
		value  uint
		result []bool
	}{
		{
			name:   "viewer",
			value:  0,
			userID: 0,
			result: []bool{false, false, false},
		},
		{
			name:   "student",
			value:  constants.STUDENT,
			userID: 1,
			result: []bool{true, false, false},
		},
		{
			name:   "organisator",
			value:  constants.ORGANISATOR,
			userID: 2,
			result: []bool{true, true, false},
		},
		{
			name:   "admin",
			value:  constants.ADMIN,
			userID: 3,
			result: []bool{true, true, true},
		},
	}
	sqlDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("CANNOT OPEN DB")
	}

	db := NewAcessDB(sqlDB)
	err = db.InitDB()
	if err != nil {
		t.Fatalf("CANNOT INIT DB")
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			uid := tt.userID
			insertRight := Right{UserID: uid, Rights: tt.value}
			db.AddRight(insertRight)
			isStudent, err := db.IsStudent(uid)
			if err != nil {
				t.Error(err.Error())
			}
			isOrgranisator, err := db.IsOrganisator(uid)
			if err != nil {
				t.Error(err.Error())
			}
			isAdmin, err := db.IsAdmin(uid)
			if err != nil {
				t.Error(err.Error())
			}
			results := []bool{isStudent, isOrgranisator, isAdmin}
			t.Logf("%v", results)
			res := slices.Equal(results, tt.result)
			if !res {
				t.Fatalf("ERROR IN RESULT IN TEST %s", tt.name)
			}
			t.Logf("SUCCES IN TEST %s", tt.name)
		})
	}

	os.Remove("test.db")
}

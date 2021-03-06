package domain

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"red/errs"
	"reflect"
	"regexp"
	"testing"
	"time"
)

var mock sqlmock.Sqlmock

func setup() *gorm.DB {
	//sqlmock
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New()
	if nil != err {
		log.Fatalf("Init sqlmock failed, err %v", err)
	}
	//gorm、sqlmock
	client, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if nil != err {
		log.Fatalf("Init DB with sqlmock failed, err %v", err)
	}
	return client
}

func TestAccountRepositoryDb_Save_Success(t *testing.T) {
	// Arrange
	client := setup()
	db := NewAccountRepositoryDb(client)

	mock.ExpectBegin()

	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `accounts`"+
			" (`customer_id`,`opening_date`,`account_type`,`amount`,`status`,`created_at`,`updated_at`,`deleted_at`,`account_id`) "+
			"VALUES (?,?,?,?,?,?,?,?,?)")).
		WithArgs(2, "2012-10-18 10:10:10", "saving", 6000.00, "1", time.Now(), time.Now(), nil, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	account := Account{
		AccountId:   1,
		CustomerId:  2,
		OpeningDate: "2012-10-18 10:10:10",
		AccountType: "saving",
		Amount:      6000,
		Status:      "1",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	_, err := db.Save(account)
	// Assert
	if err != nil {
		t.Error("test failed while save account into database")
	}
}

func TestAccountRepositoryDb_Save_Failed(t *testing.T) {
	// Arrange
	client := setup()
	db := NewAccountRepositoryDb(client)

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `accounts` (`customer_id`,`opening_date`,`account_type`,`amount`,`status`,`account_id`) VALUES (?,?,?,?,?,?)")).
		WithArgs(2, "2012-10-18 10:10:10", "saving", 6000, "1", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	account := Account{
		AccountId:   1,
		CustomerId:  2,
		OpeningDate: "2012-10-18 10:10:10",
		AccountType: "saving",
		Amount:      6000,
		Status:      "1",
	}
	_, err := db.Save(account)
	// Assert
	if want := errs.NewUnexpectedError("Unexpected error from database"); err == nil {
		t.Errorf("test failed while save account into database got: %v\nwant: %v", err, want)
	}
}

func TestAccountRepositoryDb_ByID_Success(t *testing.T) {
	// Arrange
	client := setup()
	db := NewAccountRepositoryDb(client)
	rows := sqlmock.NewRows([]string{"customer_id", "opening_date", "account_type", "amount", "status", "account_id"}).
		AddRow(2, "2012-10-18", "saving", 6000, 1, 1)

	mock.ExpectQuery(
		"^SELECT \\* FROM `accounts` WHERE account_id = \\?").
		WillReturnRows(rows)

	// Act
	_, err := db.ByID(1)
	// Assert
	if err != nil {
		t.Error("test failed while find account through id from database")
	}
	fmt.Println("error:", err)
}

func TestAccountRepositoryDb_ByID_Failed(t *testing.T) {
	// Arrange
	client := setup()
	db := NewAccountRepositoryDb(client)

	mock.ExpectQuery(
		"^SELECT \\* FROM `accounts` WHERE account_id = \\?").
		WillReturnError(sql.ErrNoRows)

	// Act
	_, err := db.ByID(1)
	// Assert
	if want := errs.NewNotFoundError("Account not found"); !reflect.DeepEqual(err, want) {
		t.Errorf("test failed while get account from database got: %v\nwant: %v", err, want)
	}
	fmt.Println("error ", err)
}

//func TestAccountRepositoryDb_SaveTransaction_Success(t *testing.T) {
//	// Arrange
//	client := setup()
//	db := NewAccountRepositoryDb(client)
//
//	mock.ExpectBegin()
//
//	mock.ExpectExec(
//		regexp.QuoteMeta("UPDATE `accounts` SET `amount`=?,`updated_at`=? WHERE `account_id` = ?")).
//		WithArgs(10000.00, time.Now(), 1).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	mock.ExpectExec(
//		regexp.QuoteMeta("INSERT INTO `transactions` "+
//			"(`account_id`,`amount`,`new_balance`,`transaction_type`,`transaction_date`,`transaction_id`) "+
//			"VALUES (?,?,?,?,?,?)")).
//		WithArgs(1, 6000.00, 10000.00, "deposit", "2012-10-18", 1).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	mock.ExpectCommit()
//
//	transaction := Transaction{
//		TransactionId:   1,
//		AccountId:       1,
//		Amount:          6000,
//		NewBalance:      4000,
//		TransactionType: "deposit",
//		TransactionDate: "2012-10-18",
//	}
//
//	// Act
//	_, err := db.SaveTransaction(transaction)
//
//	// Assert
//	if err != nil {
//		t.Error("test failed while save transaction")
//	}
//}

func TestAccountRepositoryDb_SaveTransaction_Failed(t *testing.T) {
	// Arrange
	client := setup()
	db := NewAccountRepositoryDb(client)

	mock.ExpectBegin()

	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `accounts` SET `amount`=?,`updated_at`=? WHERE `account_id` = ?")).
		WithArgs(10000.00, time.Now(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//mock.ExpectExec(
	//	regexp.QuoteMeta("INSERT INTO `transactions` (`account_id`,`amount`,`transaction_type`,`transaction_date`,`transaction_id`) VALUES (?,?,?,?,?)")).
	//	WithArgs(1, 6000, "deposit", "2012-10-18", 1).
	//	WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	transaction := Transaction{
		TransactionId:   1,
		AccountId:       1,
		Amount:          6000,
		NewBalance:      4000,
		TransactionType: "deposit",
		TransactionDate: "2012-10-18",
	}

	// Act
	_, err := db.SaveTransaction(transaction)

	// Assert
	if want := errs.NewUnexpectedError("Unexpected database error when save transaction"); !reflect.DeepEqual(err, want) {
		t.Errorf("validating error failed while save transaction \ngot: %v\nwant: %v", err, want)
	}
}

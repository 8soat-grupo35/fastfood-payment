package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/8soat-grupo35/fastfood-payment/internal/entities"
)

type PaymentRepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo    PaymentRepository
	payment *entities.Payment
}

func (suite *PaymentRepositorySuite) SetupSuite() {
	var err error

	suite.conn, suite.mock, err = sqlmock.New()
	assert.NoError(suite.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 suite.conn,
		PreferSimpleProtocol: true,
	})

	suite.DB, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(suite.T(), err)

	suite.repo = NewPaymentRepository(suite.DB)
	assert.IsType(suite.T(), &paymentRepository{}, suite.repo)

	suite.payment = &entities.Payment{
		ID:      1,
		OrderID: 123,
		Status:  "PAID",
	}
}

func (suite *PaymentRepositorySuite) AfterTest(_, _ string) {
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

// Teste de sucesso: GetPaymentStatus retorna o status corretamente
func (suite *PaymentRepositorySuite) TestGetPaymentStatus_Success() {
	rows := sqlmock.NewRows([]string{"id", "order_id", "status"}).
		AddRow(suite.payment.ID, suite.payment.OrderID, suite.payment.Status)

	suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "payments" WHERE "payments"."id" = $1 ORDER BY "payments"."id" LIMIT $2`)).
		WithArgs(suite.payment.ID, 1).
		WillReturnRows(rows)

	status, err := suite.repo.GetPaymentStatus(suite.payment.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.payment.Status, status)
}

// Teste de erro: GetPaymentStatus falha ao buscar o pagamento
func (suite *PaymentRepositorySuite) TestGetPaymentStatus_Error() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "payments" WHERE "payments"."id" = $1 ORDER BY "payments"."id" LIMIT $2`)).
		WithArgs(suite.payment.ID, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	status, err := suite.repo.GetPaymentStatus(suite.payment.ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "", status)
}

// Teste de sucesso: UpdatePaymentStatus atualiza o status corretamente
func (suite *PaymentRepositorySuite) TestUpdatePaymentStatus_Success() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "payments" SET "status"=$1 WHERE id = $2`)).
		WithArgs("COMPLETED", suite.payment.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	err := suite.repo.UpdatePaymentStatus(suite.payment.ID, "COMPLETED")
	assert.NoError(suite.T(), err)
}

// Teste de erro: UpdatePaymentStatus falha ao atualizar o status
func (suite *PaymentRepositorySuite) TestUpdatePaymentStatus_Error() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "payments" SET "status"=$1 WHERE id = $2`)).
		WithArgs("FAILED", suite.payment.ID).
		WillReturnError(errors.New("update error"))
	suite.mock.ExpectRollback()

	err := suite.repo.UpdatePaymentStatus(suite.payment.ID, "FAILED")
	assert.Error(suite.T(), err)
}

func (suite *PaymentRepositorySuite) TestCreate() {
	expectedSQL := "INSERT INTO \"payments\" (.+) VALUES (.+)"
	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	suite.mock.ExpectBegin()                                   // inicia a transação
	suite.mock.ExpectQuery(expectedSQL).WillReturnRows(addRow) // avalia o resultado
	suite.mock.ExpectCommit()                                  // commita a transação

	_, err := suite.repo.Create(*suite.payment) // chama o método Create do repository
	assert.NoError(suite.T(), err)              // avalia se não houve nenhum erro na execução
	assert.Nil(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *PaymentRepositorySuite) TestCreateReturnsErrorOnInsertFailure() {
	expectedSQL := "INSERT INTO \"payments\" (.+) VALUES (.+)"
	suite.mock.ExpectBegin() // inicia a transação
	suite.mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("insert error"))
	suite.mock.ExpectRollback() // commita a transação

	_, err := suite.repo.Create(*suite.payment)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "insert error", err.Error())
	assert.Nil(suite.T(), suite.mock.ExpectationsWereMet())
}

// Função para rodar a test suite
func TestPaymentRepositorySuite(t *testing.T) {
	suite.Run(t, new(PaymentRepositorySuite))
}

package user

import (
	"fmt"
	"os"
	"testing"

	"github.com/fajryhamzah/mkpay-user/db"
	"github.com/joho/godotenv"

	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	repo     RepoInterface
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error
	godotenv.Load("./../../.env")
	connection := db.GetTestInstance()
	dbConn := connection.GetConnection()
	repo = Repository{dbConn}

	fixtures, err = testfixtures.New(
		testfixtures.Database(dbConn),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("./../../schema/fixtures"),
	)

	defer connection.Close()

	if err != nil {
		fmt.Println("Failed on testfixtures")
		panic(err)
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Println("Failed to load fixture")
		panic(err)
	}
}

func TestFindByCode(t *testing.T) {
	prepareTestDatabase()
	var dummyCode = "e827435e-7a08-4711-8089-cf4f009694bc"
	model := repo.FindByCode(dummyCode)

	if model.GetCode() != dummyCode {
		t.Error("Wrong user code")
	}
}

func TestFindByEmail(t *testing.T) {
	prepareTestDatabase()
	var dummyEmail = "dummy@email.com"
	model := repo.FindByEmail(dummyEmail)

	if model.GetEmail() != dummyEmail {
		t.Error("Wrong user email")
	}
}

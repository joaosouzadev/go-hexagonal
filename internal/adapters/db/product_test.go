package db_test

import (
	"database/sql"
	"github.com/joaosouzadev/go-hexagonal-arch/internal/adapters/db"
	"github.com/joaosouzadev/go-hexagonal-arch/internal/application"
	"github.com/joaosouzadev/go-hexagonal-arch/pkg/utils"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

const dbUrl = "root:admin@tcp(localhost:3307)/hexagonal?parseTime=true"

func setUp() {
	Db, err = sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	err := Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func createRandomProduct() *application.Product {
	name := utils.RandomString(40)
	price := rand.Intn(10000)
	product := application.NewProduct(name, price, false, false)

	return product
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	product := createRandomProduct()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetUuid(), result.GetUuid())
	require.Equal(t, product.GetName(), result.GetName())
	require.Equal(t, product.GetPrice(), result.GetPrice())
	require.Equal(t, product.IsActive(), result.IsActive())
	require.Equal(t, product.GetOnStock(), result.GetOnStock())

	oldName := product.GetName()
	oldPrice := product.GetPrice()

	product.Name = "GeForce RTX 3060"
	product.Price = rand.Intn(10000)

	result, err = productDb.Save(product)
	require.Nil(t, err)
	require.NotEqual(t, oldName, result.GetName())
	require.NotEqual(t, oldPrice, result.GetPrice())
}

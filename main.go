package main

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"

	"log"
	"net/http"

	_ "github.com/lib/pq"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB
var Nil uuid.UUID

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DB_URI")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// e.POST("/virtual-accounts", CreateVA)
	// e.GET("/virtual-accounts/:id", CreateVA)

	e.POST("/callbacks/virtual-accounts", UpdateVACallback)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type VirtualAccount struct {
	ID             string     `json:"id,omitempty"` // will be empty when requesting
	BankCode       string     `json:"bank_code"`
	IsClosed       bool       `json:"is_closed,omitempty"`       // might be empty when requesting
	ExpectedAmount int32      `json:"expected_amount,omitempty"` // might be empty when requesting
	ExternalID     string     `json:"external_id"`
	AccountNumber  string     `json:"account_number,omitempty"` // might be empty when requesting
	Name           string     `json:"name"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"` // might be empty when requesting
	Status         string     `json:"status,omitempty"`          // might be empty when requesting
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

// func GetVA(c echo.Context) error {
// 	id := c.QueryParam("id")

// 	va, err := GetVAByID(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	return c.JSON(http.StatusOK, va)
// }

// func CreateVA(c echo.Context) error {
// 	va := &VirtualAccount{}

// 	if err := c.Bind(va); err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	id, err := SaveVA(*va)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	jsonBody, err := json.Marshal(va)
// 	if err != nil {
// 		log.Print(err)
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	res, err := SendRequest("POST", "callback_virtual_accounts", jsonBody)
// 	if err != nil {
// 		log.Print(err)
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	var jsonResponse VirtualAccount
// 	json.Unmarshal(res, &jsonResponse)

// 	err = UpdateVAByID(id, VirtualAccount{
// 		IsClosed:       jsonResponse.IsClosed,
// 		AccountNumber:  jsonResponse.AccountNumber,
// 		ExpirationDate: jsonResponse.ExpirationDate,
// 		Status:         jsonResponse.Status,
// 		VendorID:       jsonResponse.ID,
// 	})
// 	if err != nil {
// 		log.Print(err)
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	return c.JSON(http.StatusCreated, va)
// }

// func SendRequest(method string, endpoint string, body []byte) ([]byte, error) {
// 	baseUrl := os.Getenv("BASE_VENDOR_URL")
// 	apiKey := os.Getenv("API_KEY")

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, baseUrl+endpoint, bytes.NewBuffer(body))
// 	if err != nil {
// 		log.Print(err)
// 		return nil, err
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.SetBasicAuth(apiKey, "")
// 	res, err := client.Do(req)
// 	if err != nil {
// 		log.Print(err)
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	byteRes, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Print(err)
// 		return nil, err
// 	}

// 	return byteRes, nil
// }

func UpdateVACallback(c echo.Context) error {
	va := new(VirtualAccount)

	if err := c.Bind(va); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err := SaveVA(*va)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, va)
}

func GetVAByID(id string) (VirtualAccount, error) {
	va := VirtualAccount{}

	sql := `SELECT id, bank_code, is_closed, expected_amount, external_id, account_number, "name", expiration_date, status 
	FROM virtual_accounts WHERE id = $1`

	row := db.QueryRow(sql, id)
	err := row.Scan(&va.ID,
		&va.BankCode,
		&va.IsClosed,
		&va.ExpectedAmount,
		&va.ExternalID,
		&va.AccountNumber,
		&va.Name,
		&va.ExpirationDate,
		&va.Status)

	if err != nil {
		log.Print(err)
		return VirtualAccount{}, err
	}

	return va, nil

}

func SaveVA(va VirtualAccount) error {
	sql := `INSERT INTO virtual_accounts
	(id, vendor_id, bank_code, is_closed, expected_amount, external_id, account_number, "name", expiration_date, status)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	_, err := db.Exec(sql,
		va.BankCode,
		va.IsClosed,
		va.ExpectedAmount,
		va.ExternalID,
		va.AccountNumber,
		va.Name,
		va.ExpirationDate,
		va.Status)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// func UpdateVAByID(id uuid.UUID, va VirtualAccount) (err error) {
// 	sql := `UPDATE virtual_accounts SET
// 	is_closed = $1,
// 	account_number = $2,
// 	expiration_date = $3,
// 	status = $4,
// 	vendor_id = $5,
// 	updated_at = now()
// 	WHERE id = $6;`

// 	_, err = db.Exec(sql,
// 		va.IsClosed,
// 		va.AccountNumber,
// 		va.ExpirationDate,
// 		va.Status,
// 		va.VendorID,
// 		id)

// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	return
// }

// func UpdateVAByVendorID(va VirtualAccount) (err error) {
// 	sql := `UPDATE virtual_accounts SET
// 	status = $1,
// 	updated_at = now()
// 	WHERE vendor_id = $2;`

// 	_, err = db.Exec(sql, va.Status, va.VendorID)

// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	return
// }

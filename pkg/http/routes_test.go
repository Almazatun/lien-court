package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/almazatun/lien-court/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type Res[T any] struct {
	Mess   string `json:"message"`
	Res    T
	Status bool `json:"success"`
}

var accessToken string = ""

func TestRoutes(t *testing.T) {
	// Load .env.test file from the root folder.
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	err := database.ConnectDB()

	if err != nil {
		fmt.Println("🔴 Database connection error")
		log.Fatal(err)
	}

	// Define a structure for specifying input and output data of a single test case.
	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		tokenString   string // input token
		body          io.Reader
		expectedError bool
		expectedCode  int
		isPrivate     bool
	}{
		{
			description: "register user",
			route:       "/api/v1/users/register",
			method:      "POST",
			tokenString: "",
			body: strings.NewReader(
				`{
					"username": "bla",
					"email": "bla@gmail.com",
					"password": "hello@__"
				}`,
			),
			expectedError: false,
			expectedCode:  200,
			isPrivate:     false,
		},
		{
			description: "login user",
			route:       "/api/v1/auth/login",
			method:      "POST",
			tokenString: "",
			body: strings.NewReader(
				`{
					"email": "bla@gmail.com",
					"password": "hello@__"
				}`,
			),
			expectedError: false,
			expectedCode:  200,
			isPrivate:     false,
		},
		{
			description: "register user with invalid email",
			route:       "/api/v1/users/register",
			method:      "POST",
			tokenString: "",
			body: strings.NewReader(
				`{
					"username": "bla",
					"email": "bla",
					"password": "bla001"
				}`,
			),
			expectedError: false,
			expectedCode:  400,
			isPrivate:     false,
		},
		{
			description: "login user with invalid email",
			route:       "/api/v1/auth/login",
			method:      "POST",
			tokenString: "",
			body: strings.NewReader(
				`{
					"email": "bla",
					"password": "bla001"
				}`,
			),
			expectedError: false,
			expectedCode:  400,
			isPrivate:     false,
		},
		{
			description:   "auth me user",
			route:         "/api/v1/auth/me",
			method:        "GET",
			tokenString:   "Bearer ",
			body:          nil,
			expectedError: false,
			expectedCode:  200,
			isPrivate:     true,
		},
	}

	// Define a new Fiber app.
	app := fiber.New()

	// Define routes.
	PublicRoutes(app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Authorization", (test.tokenString + accessToken))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency

		if test.description == "login user" {
			// Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error:", err)
				log.Fatal(err)
			}

			// Parse the JSON response body
			var responseBody Res[struct {
				AccessToken string `json:"accessToken"`
			}]
			err = json.Unmarshal(body, &responseBody)
			if err != nil {
				fmt.Println("Error:", err)
				log.Fatal(err)
			}

			accessToken = responseBody.Res.AccessToken
			fmt.Println("✅ Token", accessToken)
		}

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		// Verify, if the status code is as expected.
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}

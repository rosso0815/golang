package main

import (
	"database/sql"
	"encoding/base32"
	"flag"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	//	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

// go run main.go -user pfistera -conn drone:drone@/drone
func main() {
	var conn string
	var user string

	// get username and database connection details from command line flags
	flag.StringVar(&conn, "conn", "", "database connection string")
	flag.StringVar(&user, "user", "", "username (i.e. octocat)")
	flag.Parse()

	// generate the user hash
	hash := base32.StdEncoding.EncodeToString(
		securecookie.GenerateRandomKey(32),
	)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	// insert the user into the database
	_, err = db.Exec(stmt, user, hash)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("user account created [%s]", user)
	}

	// generate and sign the jwt token for the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		StandardClaims: jwt.StandardClaims{},
		Type:           "user",
		Text:           user,
	})
	secret, err := token.SignedString([]byte(hash))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("user token generated [%s]", secret)
}

type claims struct {
	jwt.StandardClaims
	Type string `json:"type"`
	Text string `json:"text"`
}

const stmt = `
INSERT INTO users (
 user_login
,user_token
,user_secret
,user_expiry
,user_email
,user_avatar
,user_active
,user_admin
,user_hash
) values (?,'1','1','1','1','1',true,false,?)
`

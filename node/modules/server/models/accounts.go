package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/syahrul12345/Blocknalytics/node/modules/server/utils"
	"golang.org/x/crypto/bcrypt"
)

//Token Comment
type Token struct {
	UserID   uint
	UserName string
	jwt.StandardClaims
}

//Account Comment
type Account struct {
	gorm.Model
	Email    string
	Password string
	Token    string
}

//Validate comment
func (acc *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(acc.Email, "@") {
		return utils.Message(false, "Email Address Required"), false
	}
	if len(acc.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}
	//Email must be unique
	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", acc.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection Error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email already in use"), false
	}
	return utils.Message(false, "Requirement passed"), true

}

//Create account
func (acc *Account) Create() map[string]interface{} {
	response, ok := acc.Validate()
	if !ok {
		return response
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	acc.Password = string(hashedPassword)
	fmt.Println("before adding to database")
	fmt.Println(acc)

	//stores the account into the database
	GetDB().Create(acc)
	fmt.Println("Added to database")
	fmt.Println(acc)
	if acc.ID <= 0 {
		return utils.Message(false, "Failed to create account")
	}
	//create a new JWT token
	tk := &Token{UserID: acc.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Token = tokenString
	acc.Password = ""

	response = utils.Message(true, "Account has been created")
	response["account"] = acc
	return response
}

//Login comment
func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email")
		}
		return utils.Message(false, "Connection error. Please retry")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Invalid Login Credentials")
	}
	//worked, logged in
	account.Password = ""
	//create JWT token

	tk := &Token{}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_pasword")))
	account.Token = tokenString
	resp := utils.Message(true, "Succesfully logged in")
	resp["account"] = account
	return resp
}

//GetUser comment
func GetUser(u uint) *Account {
	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}
	acc.Password = ""
	return acc
}

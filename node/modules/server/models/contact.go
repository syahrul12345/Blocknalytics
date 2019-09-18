package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syahrul12345/Blocknalytics/node/modules/server/utils"
)

//Contact comment
type Contact struct {
	gorm.Model
	Name   string
	Phone  string
	UserID uint
}

//Validate comment
func (contact *Contact) Validate() (map[string]interface{}, bool) {
	//lets check that the contacts that are being sent is valid

	if contact.Name == "" {
		return utils.Message(false, "Contact name should be on the payload"), false
	}
	if contact.Phone == "" {
		return utils.Message(false, "Contact phone should be on the payload"), false
	}
	if contact.UserID <= 0 {
		return utils.Message(false, "User is not recognized"), false
	}
	return utils.Message(true, "success"), true
}

//Create comment
func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := utils.Message(true, "success")
	resp["contact"] = contact
	return resp

}

//GetContact comment
func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

//GetContacts comment
func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}

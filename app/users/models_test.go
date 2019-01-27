package users

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestUserCreate(t *testing.T) {
	user :=&User{string(bson.NewObjectId()),"name_test11111","password_test","email_test"}

	result,err:=createUser(user)

	assert.Nil(t, err)
	assert.NotNil(t,result)
	assert.Equal(t,user.Username,result.Username)

	result,err=getUser(user.Username)

	assert.Nil(t, err)
	assert.NotNil(t,result)
	assert.Equal(t,user.Username,result.Username)


	result,err=deleteUser(user.Username)



}

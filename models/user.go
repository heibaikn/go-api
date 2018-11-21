package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
	//orm.RegisterModel(new(UserInfo))
}

type UserInfo struct {
	Id          int
	UserId   string
	Username string
}
type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

// GetBookById retrieves Book by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *UserInfo, err error)  {
	//var aa UserInfo
	o := orm.NewOrm()
	//ids := []int{1, 2, 3}
	res, err := o.Raw("SELECT * FROM user WHERE userid=1").Exec()
	//res, err := o.Raw("SELECT * FROM user WHERE userid=1").QueryRows(&aa)
	if err == nil {
		//num, _ := res.RowsAffected()
		//fmt.Println("mysql row affected nums: ", num)
		fmt.Print(id)
		fmt.Print(res)
	}
	return nil,err
	//v = &UserInfo{Id: id}
	//if err = o.QueryTable(new(UserInfo)).Filter("Id", id).RelatedSel().One(v); err == nil {
	//	return v, nil
	//}
	//return nil, err
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

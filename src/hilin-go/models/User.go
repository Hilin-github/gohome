package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


func (user *User) TableName() string {
	return "h_user"
}


// 定义 user 实体
type User struct {
	Id	uint
	Name string `orm:"size(60)"description:"用户名"`
	Password string `orm:"size(120)"description:"密码"`
	Role uint8 `valid:"Min(0)";description:"角色默认 0"`
	Ip string `description:"登陆 ip 地址"`
	Address string `description:"住址"`
	LastLoginTime time.Time `orm:"null;auto_now;type(datetime)"description:"最后登陆时间"`
}


// 根据用户名和密码获取单条数据
func (user *User) QueryUserOneByNameAndPwd(username, password string) (*User, error){
	m := User{}
	err := orm.NewOrm().QueryTable("h_user").Filter("name",username).Filter("password",password).One(&m)
	if err != nil{
		return nil, err
	}
	return &m, nil
}

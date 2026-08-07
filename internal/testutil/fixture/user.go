// Package fixture はテストで必要となるダミーデータを用意するパッケージ
package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Suiren91/GoExampleWebApp/internal/entity"
)

// User は与えられたentity.Userをもとに，要素が足りない部分にはダミーデータを挿入し，
// すべての要素が存在するentity.Userを返す
func User(u *entity.User) *entity.User {
	res := &entity.User{
		ID:       entity.UserID(rand.Int()),
		Name:     "suiren" + strconv.Itoa(rand.Int())[:5],
		Password: "password",
		Role:     "admin",
		Created:  time.Now(),
		Modified: time.Now(),
	}

	if u == nil {
		return res
	}
	if u.ID != 0 {
		res.ID = u.ID
	}
	if u.Name != "" {
		res.Name = u.Name
	}
	if u.Password != "" {
		res.Password = u.Password
	}
	if u.Role != "" {
		res.Role = u.Role
	}
	if !u.Created.IsZero() {
		res.Created = u.Created
	}
	if !u.Modified.IsZero() {
		res.Modified = u.Modified
	}
	return res
}

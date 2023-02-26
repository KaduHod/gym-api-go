package entitys

import "time"

type PersonalE struct {
	Id        int       `json:"id"        gorm:"primary_key -all"`
	Name      string    `json:"name"      gorm:"-all"`
	Nickname  string    `json:"nickname"  gorm:"-all"`
	Email     string    `json:"email"     gorm:"-all"`
	Password  string    `json:"password"  gorm:"-all"`
	Cellphone string    `json:"cellphone" gorm:"-all"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

func (u PersonalE) GetId() int {
	return u.Id
}

func (u PersonalE) GetName() string {
	return u.Name
}

func (u PersonalE) GetNickname() string {
	return u.Nickname
}

func (u PersonalE) GetEmail() string {
	return u.Email
}
func (u PersonalE) GetPassword() string {
	return u.Password
}
func (u PersonalE) GetCellphone() string {
	return u.Cellphone
}
func (u PersonalE) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u PersonalE) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

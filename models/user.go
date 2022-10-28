package models

type User struct {
	ID        int    `orm:"column(id)" json:"id"`
	Name      string `orm:"column(name)" json:"name"`
	Password  string `json:"password"`
	Status    int    `json:"status"`
	CreatTime int64  `json:"creatTime"`
}

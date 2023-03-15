package database

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Message struct {
	Id       int64
	Author   int64
	Content  string
	IpAddr   string
	DateTime time.Time
}

func (m Message) String() string {
	return fmt.Sprintf("Messsage<%d %d %s %s>", m.Id, m.Author, m.Content, m.DateTime.Format("2006-01-02T15:04:05Z07:00"))
}

func (m Message) ToJson() gin.H {
	return gin.H{"id": m.Id, "author": m.Author, "content": m.Content, "ip_addr": m.IpAddr, "date": m.DateTime.String()}

}

type User struct {
	Id   int64
	Name string
	Icon string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.Id, u.Name, u.Icon)
}

type Guild struct {
	Id    int64
	Owner int64
	Name  string
	Icon  string
}

type Channel struct {
	Id     int64
	Belong int64
	Name   string
}

func create_schema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Message)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false, //一時作成の場合はTrue
		})
		if err != nil {
			return err
		}
	}
	return nil
}

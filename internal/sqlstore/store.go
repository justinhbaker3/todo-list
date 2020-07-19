package sqlstore

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/justinhbaker3/todo-list/internal/list"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	db *sql.DB
}

func New() (*Client, error) {
	db, err := sql.Open("mysql", "root:pass@tcp(docker.for.mac.localhost:3306)/main?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("failed to open: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping: %s", err)
	}

	return &Client{db: db}, nil
}

func (c *Client) Get(title string) (*list.List, error) {
	query := `SELECT items.title, items.description, items.opened
				FROM lists JOIN items ON lists.id=items.list
				WHERE lists.title=?`

	l := list.NewList(title)

	rows, err := c.db.Query(query, title)
	if err != nil {
		return nil, err
	}
	for rows.Next() == true {
		var itemTitle, itemDescription string
		var opened time.Time
		err := rows.Scan(&itemTitle, &itemDescription, &opened)
		if err != nil {
			return nil, err
		}
		err = l.Append(list.NewItem(itemTitle, itemDescription))
		if err != nil {
			return nil, err
		}
	}
	return l, nil
}

func (c *Client) Upsert(list *list.List) {
	insertList := `INSERT INTO lists (title) values ("?")
					ON DUPLICATE KEY UPDATE title=?`
	insertItem := `INSERT INTO items (title, description, list)
					values ("?", "?", ?)
					ON DUPLICATE KEY UPDATE title = VALUES(title), description = VALUES(description), list = VALUES(list)`

	res, _ := c.db.Exec(insertList, list.Title)
	listID, _ := res.LastInsertId()

	for _, item := range list.Items {
		c.db.Exec(insertItem, item.Title, item.Description, listID)
	}

}

func (c *Client) Delete(title string) {
	// TODO
	return
}

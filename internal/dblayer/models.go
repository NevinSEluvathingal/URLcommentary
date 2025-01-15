// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dblayer

import (
	"time"
)

type Comment struct {
	Commentid   int32     `json:"commentid"`
	Pageid      int32     `json:"pageid"`
	Userid      int32     `json:"userid"`
	Createdat   time.Time `json:"createdat"`
	Editedbool  bool      `json:"editedbool"`
	Upvotes     int32     `json:"upvotes"`
	Downvotes   int32     `json:"downvotes"`
	Commentdata string    `json:"commentdata"`
	Parentid    int32     `json:"parentid"`
}

type Page struct {
	Pageurl       string    `json:"pageurl"`
	Pageid        int32     `json:"pageid"`
	Commentscount int32     `json:"commentscount"`
	Createdat     time.Time `json:"createdat"`
}

type User struct {
	Userid    int32  `json:"userid"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Emailid   string `json:"emailid"`
}

type Userrelation struct {
	Vieweruserid    int32  `json:"vieweruserid"`
	Commentoruserid int32  `json:"commentoruserid"`
	Tag             string `json:"tag"`
	Positivity      bool   `json:"positivity"`
}

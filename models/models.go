// store the database schema
// The User struct is a representation of users table
package models

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	Age int64 `json:"age"`
}

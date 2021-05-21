import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time `gorm:"index"`
}

type Mentor struct {
	ID int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time `gorm:"index"`
}

type Booking struct {
	ID int `gorm:"primaryKey"`
	ClassProgram string
	BookingDate time.Time
	DurasiJam int 

}

type Role struct {
	ID int `gorm:"primarykey"`
	LoginAs string
	LoginDate time.Time
}
package entity

import (
	// "time"

	// "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// -------------------------------------------<< User >>------------------------------------

type Gender struct {
	gorm.Model
	Name string
	User []User `gorm:"foreignKey:GenderID"`
}

type Province struct {
	gorm.Model
	Name     string
	District []District `gorm:"foreignKey:ProvinceID"`
}

type District struct {
	gorm.Model
	Name     string
	Postcode string

	ProvinceID *uint
	Province   Province

	User []User `gorm:"foreignKey:DistrictID"`
}

type Occupation struct {
	gorm.Model
	Name string
	User []User `gorm:"foreignKey:OccupationID"`
}

type Salary struct {
	gorm.Model
	Salary string
	User   []User `gorm:"foreignKey:SalaryID"`
}

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Role      string

	GenderID *uint
	Gender   Gender

	DistrictID *uint
	District   District

	OccupationID *uint
	Occupation   Occupation

	SalaryID *uint
	Salary   Salary

	Product []Product `gorm:"foreignKey:UserID"`
}

// -------------------------------------------<< Product >>------------------------------------

type ProductType struct {
	gorm.Model
	Name    string
	Product []Product `gorm:"foreignKey:ProductTypeID"`
}

type Promotion struct {
	gorm.Model
	Name       string
	Percentage int
	Product    []Product `gorm:"foreignKey:PromotionID"`
}

type Product struct {
	gorm.Model
	Name         string
	ProductPhoto string
	Description  string

	ProductTypeID *uint
	ProductType   ProductType

	PromotionID *uint
	Promotion   Promotion

	UserID *uint
	User   User
}

// -------------------------------------------<< Purchase >>------------------------------------

type PaymentMethod struct {
	gorm.Model
	Name     string
	Method   string     // QRCODE/Bank Transfer
	Purchase []Purchase `gorm:"foreignKey:PaymentMethodID"`
}

type Receipt struct {
	gorm.Model
	ReceiptCode string

	Purchase []Purchase `gorm:"foreignKey:ReceiptID"`
}

type Purchase struct {
	gorm.Model
	Name         string
	ProductImage string
	Description  string

	PaymentMethodID *uint
	PaymentMethod   PaymentMethod

	ReceiptID *uint
	Receipt   Receipt

	UserID *uint
	User   User

	ProductID *uint
	Product   Product
}

// func init() {
// 	// Custom valid tag
// 	govalidator.TagMap["image"] = govalidator.Validator(func(str string) bool {
// 		pattern := "^data:image/(jpeg|jpg|png|svg|gif|tiff|tif|bmp|apng|eps|jfif|pjp|xbm|dib|jxl|svgz|webp|ico|pjpeg|avif);base64,[A-Za-z0-9+/]+={0,2}$"
// 		return govalidator.Matches(str, pattern)
// 	})

// 	govalidator.CustomTypeTagMap.Set("notpast30min", func(i interface{}, context interface{}) bool {
// 		t := i.(time.Time)
// 		return t.After(time.Now().Add(time.Minute * -30))
// 	})

// 	govalidator.CustomTypeTagMap.Set("notfuture30min", func(i interface{}, context interface{}) bool {
// 		t := i.(time.Time)
// 		return t.Before(time.Now().Add(time.Minute * 30))
// 	})

// 	// govalidator.TagMap["age"] = govalidator.IsPositive(Trainer);
// 	// govalidator.CustomTypeTagMap.Set("age", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
// 	// 	s, ok := i.(float64)
// 	// 	if !ok {
// 	// 		fmt.Print(s, "Aonon")
// 	// 		return false
// 	// 	}
// 	// 	return s < 1
// 	// }))

// }

// Firstname   string `valid:"required~Firstname cannot be blank"`
// Lastname    string `valid:"required~Lastname cannot be blank"`
// Email       string `valid:"email~Invalid email format,required~Email cannot be blank"`
// Password    string `valid:"required~Password cannot be blank,minstringlength(8)~Password must be no less than 8 characters long"`

// Name        string `valid:"maxstringlength(50)~Course Name must be no more than 50 characters,required~Course Name cannot be blank"`
// CoverPage   string `valid:"image~Cover Page must be images file"`
// Description string `valid:"maxstringlength(300)~Description must be no more than 300 characters,required~Description cannot be blank"`
// Goal        string `valid:"maxstringlength(100)~Goal must be no more than 100 characters,required~Goal cannot be blank"`

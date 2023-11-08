package model

type Unit struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Unit *Unit  `json:"unit" gorm:"foreignkey:ID"`
}

type Phone struct {
	ID          int       `json:"id"`
	NumberPhone string    `json:"numberPhone"`
	EmployeeID  int       `json:"employeeId"`
	Employee    *Employee `json:"employee" gorm:"foreignkey:EmployeeID"`
	Unit        *Unit     `json:"unit" gorm:"foreignkey:ID"`
}

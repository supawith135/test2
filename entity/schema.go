package entity

type User struct {
	StudentID string `valid:"matches(^[B]\\d{7}$)~กรุณาใส่ชื่อขึ้นต้นด้วย Bxxxxxxx, required~กรุณากรอกรหัสนักศึกษา"`
	Phone     string 
	Url       string `valid:"url~กรุณากรอกให้ถูกต้อง"`
	Email     string `valid:"email~กรุณากรอกอีเมลให้ถูกต้อง"`
}

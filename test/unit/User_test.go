package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/supawith135/labtest2/entity"
)

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`check StudentID`, func(t *testing.T) {
		User := entity.User{
			StudentID: "M6432140",
			Phone:     "",
			Url:       "github.com/asaskevich/govalidator",
			Email:     "test123@hotmail.com",
		}
		ok, err := govalidator.ValidateStruct(User)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณาใส่ชื่อขึ้นต้นด้วย Bxxxxxxx"))
	})
}

func TestUserRequired(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`check StudentID Required`, func(t *testing.T) {
		User := entity.User{
			StudentID: "",
			Phone:     "",
			Url:       "github.com/asaskevich/govalidator",
			Email:     "test123@hotmail.com",
		}
		ok, err := govalidator.ValidateStruct(User)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกรหัสนักศึกษา"))

	})

}


func TestURl(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`check URL`, func(t *testing.T) {
		User := entity.User{
			StudentID: "B6432140",
			Phone:     "",
			Url:       "github/asaskevich/govalidator",
			Email:     "test123@hotmail.com",
		}
		ok, err := govalidator.ValidateStruct(User)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกให้ถูกต้อง"))
	})
}


func TestEmail(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`check Email`, func(t *testing.T){
		User := entity.User{
			StudentID: "B6432140",
			Phone:     "",
			Url:       "github.com/asaskevich/govalidator",
			Email:     "test123sdcom",
		}
		ok, err := govalidator.ValidateStruct(User)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกอีเมลให้ถูกต้อง"))
	})
}
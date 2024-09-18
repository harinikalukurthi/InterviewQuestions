/*
1. The code should have functionality to create users. Users can be either customer or admin. All users will have a unique identifier: username.
2. Admin can create a Loan in the system for a customer.
3. While creating a loan, admin_username, customer_username, principal amount, interest rate and time (loan tenure) need to be taken as input.
The interest for the loan is calculated by I = (P*N*R)/100 where P is the principal amount, N is the number
of years and R is the rate of interest. The total amount to repay will be A = P + I The amount should be paid back monthly in the form of EMIs. Each EMI = A/(N*12)
Users should be able to fetch loan info for their loans only. Fetching a loan should return the loan info along with all the emi payments done and EMIs remaining[optional].
4. Admin should be able to fetch all loans for all customers.
5. All the functions should take username as one of the arguments, and all user level validation should happen against this username.
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Loan struct {
	Admin   string
	User    string
	Amount  int
	Span    int
	Interst int
}
type CustomerLoanDetails struct{
	Username string
	TotalAmount int
	EachEmi int
}

var LoanDetails []CustomerLoanDetails

func main() {
	r := gin.Default()

	auth := gin.BasicAuth(gin.Accounts{
		"Harini": "123",
	})

	admin := r.Group("/admin", auth)
	{
		admin.POST("/create-loan",CreateLoan)
		admin.GET("/loans",getLoanByAdmin)
	}
	customer:=r.Group("/user",auth)
	{
		customer.GET("/loan/:user",getLoanByCustomer)
	}

	r.Run()

}
func CreateLoan(c *gin.Context) {
	var l Loan

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot take loan details"})
		return
	}
interest:=(l.Amount*l.Span*l.Interst)/100
totalAmount:=l.Amount+interest
EMI:=totalAmount/(l.Span*12)
var y CustomerLoanDetails
y.Username=l.User
y.TotalAmount=totalAmount
y.EachEmi=EMI
LoanDetails=append(LoanDetails, y)

}
func getLoanByCustomer(c *gin.Context){
	user:=c.Param("user")
	 for _,x:=range LoanDetails{
		if x.Username==user{
			c.JSON(http.StatusOK,gin.H{"completed Emis",},gin.H{"remaining emis":x.TotalAmount/x.EachEmi})
			return
		}
	 }
	 c.JSON(http.StatusNotFound,gin.H{"message":"user not found"})
}

func getLoanByAdmin(c *gin.Context){
	c.JSON(http.StatusOK,LoanDetails)
}
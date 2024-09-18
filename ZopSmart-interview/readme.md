Change password: 

/changepassword,<handler-name>

package main
m:=make(map[string]string)


func main(){
 
r:=gin.Default()
r.PUT(“/:change”,changePass)
r.Run()
}



func changePass(r *gin.Context){
newpass:=r.Param("change")
If pass,ok:=m[user-name];ok{
If pass!=newpass{
m[user-name]=newpass
}else{
r.JSON(statuscode,”string”)
}
}

}

Basic auth
JWT
Oidc
Saml
Bearer token


 Curl -u “user”: “password” -X PUT https://localhost:<port/<end-point>


=============================

Restaurant : [Name, Photo, FoodName, FoodPhoto, Price]

Restaurant:[Name,Photo,FoodId] 
Items:[FoodId,FoodName,FoodPhoto,Price]


Select * 
From Restaurant r
JOIN Items  i ON i.FoodId=r.FoodId
Where r.Name=A 

A and B



Restaurant:
[A, PhotoA, Rice]
[B, PhotoB, Rice]

Items:
[Rice, RiceA, 25]
[Rice, RiceB, 40]


Example Query:
Select * 
From Restaurant r
JOIN Items  i ON i.FoodName=r.FoodName 

What will the output for the above example table?

[A,photoA,Rice,RiceA,25]
[A,photoA,Rice,RiceB,40]
[B,photB,rice,riceA,25]
[B,photB,rice,riceB,45]

restaurant:
[A,photoA,1]
[B,photoB,2]
[B,photoC,3]

Items:
[1,Rice, RiceA, 25]
[2,Rice, RiceB, 40]


==============


Employee: [Name, ID, MgrID] where Name is the employee name and ID is the identity of an employee.

Write a SQL query to list the emp name and his/her manager’s name.

Select e.Name, m.Name
From Employee e
Inner Join Employee m on m.ID=e.MgrID
         


//use single channel to connect both goroutines.
// I want to send the data from one channel and want to receive the channel f

what is goroutines leak
let say we have two goroutines,one go routinue
what is that we have to look while writing a testcase for a function in go
what is hashing and how can we do that.




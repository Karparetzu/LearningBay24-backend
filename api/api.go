package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"learningbay24.de/backend/config"
	"learningbay24.de/backend/course"
	"learningbay24.de/backend/models"
)



func GetCourseById(c *gin.Context) {
    db := config.SetupDbHandle()
	//Get given ID from the Context
	//Convert data type from str to int to use ist as param 
	id,err := strconv.Atoi(c.Param("id"))
	//Fetch Data from Database with Backend function
    course,err := course.GetCourse(db,id)
	//Return Status and Data in JSON-Format
    c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, course)
	fmt.Println("course ", course)
	fmt.Println(err)
}


func DeleteUserFromCourse(c *gin.Context) {
    db := config.SetupDbHandle()
	//Get given ID from the Context
	//Convert data type from str to int to use ist as param 
	user_id, err := strconv.Atoi(c.Param("user_id"))
	id,err := strconv.Atoi(c.Param("id"))
	//Fetch Data from Database with Backend function
    user,err := course.DeleteUserFromCourse(db,id,user_id)
	//Return Status and Data in JSON-Format
    c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, user)
	fmt.Println("course ", user)
	fmt.Println(err)
}


func GetUsersInCourse(c *gin.Context) {
    db := config.SetupDbHandle()
	//Get given ID from the Context
	//Convert data type from str to int to use ist as param 
	id,err := strconv.Atoi(c.Param("id"))
	//Fetch Data from Database with Backend function
    users,err := course.GetUserCourses(db,id)
	//Return Status and Data in JSON-Format
    c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, users)
	fmt.Println("course ", users)
	fmt.Println(err)
}



func GetUserCourses(c *gin.Context) {
    db := config.SetupDbHandle()
	//Get given ID from the Context
	//Convert data type from str to int to use ist as param 
	user_id, err := strconv.Atoi(c.Param("user_id"))
	//Fetch Data from Database with Backend function
    courses,err := course.GetUserCourses(db,user_id)
	//Return Status and Data in JSON-Format
    c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, courses)
	fmt.Println("course ", courses)
	fmt.Println(err)
}



func DeleteCourseById(c *gin.Context) {
	db := config.SetupDbHandle()
	//Get given ID from the Context
	//Convert data type from str to int to use ist as param 
	id,err := strconv.Atoi(c.Param("id"))
	//Delete Data from Database with Backend function
    course,err := course.DeleteCourse(db,id)
	//Return Status and Data in JSON-Format
    c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, course)
	fmt.Println("course ", course)
	fmt.Println(err)
}


func DeactivateCourse(c *gin.Context) {
	db := config.SetupDbHandle()
	//Get given ID from the Context
	//Convert data type from str to int to use ist as param 
	id,err := strconv.Atoi(c.Param("id"))
	//Deactivate Data from Database with Backend function
    course,err := course.DeactivateCourse(db,id)
	//Return Status and Data in JSON-Format
    c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, course)
	fmt.Println("course ", course)
	fmt.Println(err)
}




func CreateCourse(c *gin.Context) {
	db := config.SetupDbHandle()
	var newCourse models.Course
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}
	id,err := course.CreateCourse(db, newCourse.Name,newCourse.Description, newCourse.EnrollKey,  user_id)
	if err != nil {
		fmt.Println(err.Error())
		panic("error creating course")
	}
	newCourse.ID = id
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, newCourse)
}

func EnrollUser(c *gin.Context) {
	db := config.SetupDbHandle()
	
	id,err := strconv.Atoi(c.Param("id"))
	user_id, err := strconv.Atoi(c.Param("user_id"))
	var newCourse models.Course
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}
	user,err := course.EnrollUser(db, user_id,id,newCourse.EnrollKey)

	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, newCourse)
}




func UpdateCourseById(c *gin.Context) {
	db := config.SetupDbHandle()
	id, err := strconv.Atoi(c.Param("id"))
	var newCourse models.Course
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}
	course,err := course.UpdateCourse(db,id, newCourse.Name,newCourse.Description, newCourse.EnrollKey)
	if err != nil {
		fmt.Println(err.Error())
		panic("error creating course")
	}
	newCourse.ID = id
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, newCourse)
}


/* func GetCourses(c *gin.Context) {
	db := config.SetupDbHandle()
	var courses []models.Course

	err := queries.Raw("select * from course").Bind(context.Background(), db, &courses)
	if err != nil {
		fmt.Println(err.Error())
		panic("error raw query")
	}

	c.Header("Access-Control-Allow-Origin", "*")
	//Return Status and Data in JSON-Format
	c.IndentedJSON(http.StatusOK, courses)
	fmt.Println("courses ", courses)
	fmt.Println(err)
} */
package main
import "github.com/gin-gonic/gin"
type event struct{
	 ID string
	 Title string
	 Description  string

}
type allevents []event
var events = allevents{
	{
ID: "1",
Title: "Gin Gonic event",
Description: "problem solving with gin ",
	},
}



// handlers 
func getAllEvents(c *gin.Context){
	//returning a map with all event 
	c.JSON(200,gin.H{
   "data":events,
	})

}



// one event 

func getOneEvent(c *gin.Context){
	// from url 
eventID:=c.Param("id")
for _ , singleEvent:= range events{
	if singleEvent.ID ==eventID{
		c.JSON(200,gin.H{
			"data":singleEvent,
			 })
	}
}

}

func createEvent(c *gin.Context){
	id:=c.PostForm("ID")
	Title:=c.PostForm("Title")
	Description:=c.PostForm("Description")
	newEvent:=event{id,Title,Description}
	events=append(events, newEvent)
	c.JSON(201,gin.H{
     "status":201,
	 "message":"new event was created",
	  "event":newEvent	},
	
	)

}
func update(c *gin.Context){
	eventID:=c.Param("id")
	id:=c.PostForm("ID")
	Title:=c.PostForm("Title")
	Description:=c.PostForm("Description")


	for i , singleEvent:= range events{
		if singleEvent.ID ==eventID{
			singleEvent.ID =id
			singleEvent.Title=Title
			singleEvent.Description=Description

			events=append(events[:i], singleEvent)
	c.JSON(201,gin.H{
     "status":200,
	 "message":"updated event ",
	  "event":singleEvent	},
	
	)
		}
	}
}


func delete(c *gin.Context){
	eventID:=c.Param("id")
	for i,singleevent:=range events{
		if singleevent.ID == eventID{
		events = append(events[:i],events[i+1:]...)	
		c.JSON(201,gin.H{
			"status":200,
			"message":" event was deleted",
			 "event":singleevent	},
		   
		   )
		}
	}
}














 func main(){
	// creation router 
	r:=gin.Default()
	// to organize and modularize routes based on their functionality or versioning
	v1:=r.Group("/api/v1")
	{
		//api/v1
		v1.GET("/",func (c *gin.Context)  {
			c.JSON(200,gin.H{
				"message":"hello",
			})
		})
		v1.GET("/events", getAllEvents)
		v1.GET("/events/:id", getOneEvent)
		v1.POST("/event",createEvent )
		v1.PATCH("/events/:id", update)
		v1.DELETE("/events/:id", delete)
	}
	//listen an serve
r.Run(":8080")

 }
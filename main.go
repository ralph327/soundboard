package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"log"
	"strings"
)

type AudioLibrary struct{
	Lib map[string]*Audio
}

func main(){
	
	/* Initiate the audio library
	 */ 
	audioLibrary := new(AudioLibrary)
	audioLibrary.Lib = make(map[string]*Audio) 
	
	/* Build the audio library
	 */ 
	audioLibrary.Build()
	
	/* Initiate the Server
	 */ 
	r := gin.Default()
	html := template.Must(template.ParseGlob("style/tmpl/*"))
     r.SetHTMLTemplate(html)
     
     /* Serve Static files
      */
     r.Static("/style/css","style/css")
     r.Static("/scripts", "./scripts")
     r.Static("/images", "./images")
	
	/* Set the Routes
	 */ 
	r.GET("/", audioLibrary.Home())
	r.POST("/play", audioLibrary.Play())
	
	/* Start the Server
	 */ 
	r.Run(":7777") 
 }
 
 func (al *AudioLibrary) Build(){
	 files, err := ioutil.ReadDir("./audio")
	 
	 if err != nil {
		log.Fatal(err) 
	 }
	 
	 for _, file := range files {
		 substr := strings.Split(file.Name(), ".")
		 name := substr[0]
		 al.Lib[name] = newAudio(file.Name())
	 }
 }

func (al *AudioLibrary) Home() gin.HandlerFunc {
	data := gin.H{"domainName": "Soundboard"}
	
	var body template.HTML
	
	body += template.HTML(`<ul>`)
	
	for _, file := range al.Lib {
		body += template.HTML(`
			<li class="block" id="`+file.Name+`">
			<P>` + file.Filename + `</P>
			</li>
		`)
	}
	
	body += template.HTML(`</ul>`)
	
	data["body"] = body
	
	return func(c *gin.Context){
		c.HTML(http.StatusOK, "base", data)
	}
}

func (al *AudioLibrary) Play() gin.HandlerFunc {
	return func(c *gin.Context){
		
		name := c.PostForm("name")
		
		err := al.Lib[name].Cmd.Run()
		
		if err != nil {
			log.Fatal(err)
		}
		
		c.JSON(http.StatusOK, gin.H{"message": name,})
	}
}

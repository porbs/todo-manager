package controllers

import (
	"log"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/revel/revel"
	"todo-manager/app/models"
	"todo-manager/app/database"

)

type Notes struct {
	*revel.Controller
}

func (c Notes) Get() revel.Result {

	result := []models.Note{}

	if err := database.Notes.Find(bson.M{}).All(&result); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(result)
}

func (c Notes) Create(note models.Note) revel.Result {

	if err := database.Notes.Insert(note); err != nil {
		log.Fatal(err)
		c.Response.Status = http.StatusInternalServerError
		return c.RenderText("Could not be saved")
	}

	c.Response.Status = http.StatusCreated
	return c.RenderJSON(note)
}

func (c Notes) Update(note models.Note) revel.Result {

	if err := database.Notes.UpdateId(note.ID, note); err != nil {
		log.Fatal(err)
		c.Response.Status = http.StatusInternalServerError
		return c.RenderText("Could not be saved")
	}

	return c.RenderJSON(note)
}

func (c Notes) Delete(id string) revel.Result{

	if !bson.IsObjectIdHex(id) {
		c.Response.Status = http.StatusBadRequest
		return c.RenderText("ID is not valid")
	} else if obj := bson.ObjectIdHex(id); !obj.Valid() {
		c.Response.Status = http.StatusBadRequest
		return c.RenderText("ID is not valid")
	} else if err := database.Notes.RemoveId(obj); err != nil {
		log.Fatal(err)
		c.Response.Status = http.StatusInternalServerError
		return c.RenderText("Could not perform operation")
	}

	return c.RenderText(id)



}
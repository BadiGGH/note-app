package handlers

import (
	"net/http"
	"strconv"

	"github.com/BadiGGH/note-app/database"
	"github.com/BadiGGH/note-app/models"
	"github.com/gin-gonic/gin"
)

// CreateNewNote godoc
// @Summary Create a new note
// @Description Create a new note with title and content
// @Tags notes
// @Accept json
// @Produce json
// @Param note body models.Note true "Note data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /notes [post]
func CreateNewNote(c *gin.Context) {
	var newNote models.Note

	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data!"})
		return
	}
	id, err := database.InsertNote(newNote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Inserted", "note_id": id})
}

// ReadAllNotesHandler godoc
// @Summary Get all notes
// @Tags notes
// @Produce json
// @Success 200 {array} models.Note
// @Failure 500 {object} map[string]string
// @Router /notes [get]
func ReadAllNotesHandler(c *gin.Context) {
	notes, err := database.AllNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}

// ReadNoteByIdHandler godoc
// @Summary Get note by ID
// @Tags notes
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} models.Note
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [get]
func ReadNoteByIdHandler(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id!"})
		return
	}
	note, err := database.NoteById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, note)
}

// UpdateNoteHandler godoc
// @Summary Update a note
// @Tags notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Param note body models.Note true "Updated note"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /notes/{id} [put]
func UpdateNoteHandler(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id!"})
		return
	}
	note, err := database.NoteById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data!"})
		return
	}
	affectedRowsCount, err := database.UpdateNote(id, note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if affectedRowsCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated", "affected_rows": affectedRowsCount})
}

// DeleteNoteByIdHandler godoc
// @Summary Delete a note
// @Tags notes
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /notes/{id} [delete]
func DeleteNoteByIdHandler(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id!"})
		return
	}
	affectedRowsCount, err := database.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if affectedRowsCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted", "affected_rows": affectedRowsCount})
}

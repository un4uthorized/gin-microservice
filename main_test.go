package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"

	"github.com/gin-gonic/gin"
)

func TestGetAlbums(t *testing.T) {	
	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)

	GetAlbums(ctx)

	if w.Code != http.StatusOK {
		t.Errorf("GetAlbums() returned %d, want %d", w.Code, http.StatusOK)
	}
}

func TestReturnSpecifAlbum (t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	GetAlbumByID(ctx)

	if w == nil {
		t.Error("GetAlbumByID() returned nil")
	}

	if w.Code != http.StatusOK {
		t.Errorf("GetAlbumByID() returned %d, want %d", w.Code, http.StatusOK)
	}
}

func TestCreateAlbums (t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/albums", bytes.NewBufferString(`{"id": "4", "title": "Test", "artist": "Test", "price": 1.99}`))
	CreateAlbums(ctx)

	if w.Code != http.StatusCreated {
		t.Errorf("CreateAlbums() returned %d, want %d", w.Code, http.StatusCreated)
	}
}

func TestReturnNotFound (t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: "any number"}}
	GetAlbumByID(ctx)

	if w.Code != http.StatusNotFound {
		t.Errorf("GetAlbumByID() returned %d, want %d", w.Code, http.StatusNotFound)
	}
}




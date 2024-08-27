package handlers

import (
	"github.com/gorilla/mux"
	"library-management/middlewares/logging"
	"library-management/models"
	"library-management/utils"
	"net/http"
	"strconv"
)

func GetAllBooks(res http.ResponseWriter, req *http.Request) {

	queryParams := map[string]string{
		"title":    req.URL.Query().Get("title"),
		"author":   req.URL.Query().Get("author"),
		"minPrice": req.URL.Query().Get("minPrice"),
		"maxPrice": req.URL.Query().Get("maxPrice"),
		"sortBy":   req.URL.Query().Get("sortBy"),
		"sortDir":  req.URL.Query().Get("sortDir"),
	}

	books, err := BooksRepo.FindAll(queryParams)
	if err != nil {
		utils.SendErrorResponse(res, "Error getting books", http.StatusInternalServerError)
		return
	}
	err = utils.JsonEncode(res, books)
	if err != nil {
		utils.SendErrorResponse(res, "Error encoding books", http.StatusInternalServerError)
	}
}

func GetBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.SendErrorResponse(res, "Error parsing id", http.StatusBadRequest)
		return
	}
	book, err := BooksRepo.FindById(id)
	if err != nil {
		utils.SendErrorResponse(res, "Error getting book", http.StatusInternalServerError)
		return
	}
	err = utils.JsonEncode(res, book)
	if err != nil {
		logging.Log.Error("Error encoding books", err.Error())
	}
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	var book models.Book
	err := utils.JsonDecode(req, &book)
	if err != nil {
		logging.Log.Error("Error encoding books", err)
		utils.SendErrorResponse(res, "Error decoding book", http.StatusBadRequest)
		return
	}
	err = BooksRepo.Create(&book)
	if err != nil {
		utils.SendErrorResponse(res, "Error creating book", http.StatusInternalServerError)
		return
	}
	err = utils.JsonEncode(res, book)
	if err != nil {
		logging.Log.Error("Error encoding book", err.Error())
	}
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.SendErrorResponse(res, "Error parsing id", http.StatusBadRequest)
		return
	}

	book, err := BooksRepo.FindById(id)
	if err != nil {
		utils.SendErrorResponse(res, "Error getting book", http.StatusInternalServerError)
		return
	}

	if book.ID == 0 {
		utils.SendErrorResponse(res, "Book with Id"+string(rune(id))+"not found", http.StatusExpectationFailed)
		return
	}

	err = utils.JsonDecode(req, &book)
	if err != nil {
		utils.SendErrorResponse(res, "Invalid book data", http.StatusBadRequest)
		return
	}

	err = BooksRepo.Update(book)
	if err != nil {
		utils.SendErrorResponse(res, "Error updating book", http.StatusInternalServerError)
		return
	}

	utils.JsonEncode(res, book)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logging.Log.Error("Error parsing id", http.StatusBadRequest)
		return
	}
	book, err := BooksRepo.FindById(id)
	if err != nil {
		utils.SendErrorResponse(res, "Error getting book", http.StatusInternalServerError)
		return
	}
	err = BooksRepo.Delete(book)
	if err != nil {
		utils.SendErrorResponse(res, "Error deleting book", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

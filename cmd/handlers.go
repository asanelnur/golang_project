package main

/*
func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	category := &model.Category{
		Title:       input.Title,
		Description: input.Description,
	}

	err = app.models.Categories.Insert(category)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, category)
}

func (app *application) getCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["categoryId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := app.models.Categories.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, category)
}

func (app *application) updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["categoryId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := app.models.Categories.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.Title != nil {
		category.Title = *input.Title
	}

	if input.Description != nil {
		category.Description = *input.Description
	}

	err = app.models.Categories.Update(category)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, category)
}

func (app *application) deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["categoryId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	err = app.models.Categories.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}
*/

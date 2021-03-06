package actions

import (

  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "github.com/gobuffalo/x/responder"
  "github.com/hurtownia/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Kategorie)
// DB Table: Plural (kategories)
// Resource: Plural (Kategories)
// Path: Plural (/kategories)
// View Template Folder: Plural (/templates/kategories/)

// KategoriesResource is the resource for the Kategorie model
type KategoriesResource struct{
  buffalo.Resource
}

// List gets all Kategories. This function is mapped to the path
// GET /kategories
func (v KategoriesResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  kategories := &models.Kategories{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Kategories from the DB
  if err := q.All(kategories); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    c.Set("kategories", kategories)
    return c.Render(http.StatusOK, r.HTML("/kategories/index.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(kategories))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(kategories))
  }).Respond(c)
}

// Show gets the data for one Kategorie. This function is mapped to
// the path GET /kategories/{kategorie_id}
func (v KategoriesResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Kategorie
  kategory := &models.Kategorie{}

  // To find the Kategorie the parameter kategorie_id is used.
  if err := tx.Find(kategory, c.Param("kategorie_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    c.Set("kategory", kategory)

    return c.Render(http.StatusOK, r.HTML("/kategories/show.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(kategory))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(kategory))
  }).Respond(c)
}

// New renders the form for creating a new Kategorie.
// This function is mapped to the path GET /kategories/new
func (v KategoriesResource) New(c buffalo.Context) error {
  c.Set("kategory", &models.Kategorie{})

  return c.Render(http.StatusOK, r.HTML("/kategories/new.plush.html"))
}
// Create adds a Kategorie to the DB. This function is mapped to the
// path POST /kategories
func (v KategoriesResource) Create(c buffalo.Context) error {
  // Allocate an empty Kategorie
  kategory := &models.Kategorie{}

  // Bind kategory to the html form elements
  if err := c.Bind(kategory); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(kategory)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the new.html template that the user can
      // correct the input.
      c.Set("kategory", kategory)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/kategories/new.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "kategory.created.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/kategories/%v", kategory.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.JSON(kategory))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.XML(kategory))
  }).Respond(c)
}

// Edit renders a edit form for a Kategorie. This function is
// mapped to the path GET /kategories/{kategorie_id}/edit
func (v KategoriesResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Kategorie
  kategory := &models.Kategorie{}

  if err := tx.Find(kategory, c.Param("kategorie_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  c.Set("kategory", kategory)
  return c.Render(http.StatusOK, r.HTML("/kategories/edit.plush.html"))
}
// Update changes a Kategorie in the DB. This function is mapped to
// the path PUT /kategories/{kategorie_id}
func (v KategoriesResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Kategorie
  kategory := &models.Kategorie{}

  if err := tx.Find(kategory, c.Param("kategorie_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Kategorie to the html form elements
  if err := c.Bind(kategory); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(kategory)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the edit.html template that the user can
      // correct the input.
      c.Set("kategory", kategory)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/kategories/edit.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "kategory.updated.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/kategories/%v", kategory.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(kategory))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(kategory))
  }).Respond(c)
}

// Destroy deletes a Kategorie from the DB. This function is mapped
// to the path DELETE /kategories/{kategorie_id}
func (v KategoriesResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Kategorie
  kategory := &models.Kategorie{}

  // To find the Kategorie the parameter kategorie_id is used.
  if err := tx.Find(kategory, c.Param("kategorie_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(kategory); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a flash message
    c.Flash().Add("success", T.Translate(c, "kategory.destroyed.success"))

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "/kategories")
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(kategory))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(kategory))
  }).Respond(c)
}

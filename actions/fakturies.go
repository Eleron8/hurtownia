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
// Model: Singular (Faktury)
// DB Table: Plural (fakturies)
// Resource: Plural (Fakturies)
// Path: Plural (/fakturies)
// View Template Folder: Plural (/templates/fakturies/)

// FakturiesResource is the resource for the Faktury model
type FakturiesResource struct{
  buffalo.Resource
}

// List gets all Fakturies. This function is mapped to the path
// GET /fakturies
func (v FakturiesResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  fakturies := &models.Fakturies{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Fakturies from the DB
  if err := q.All(fakturies); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    c.Set("fakturies", fakturies)
    return c.Render(http.StatusOK, r.HTML("/fakturies/index.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(fakturies))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(fakturies))
  }).Respond(c)
}

// Show gets the data for one Faktury. This function is mapped to
// the path GET /fakturies/{faktury_id}
func (v FakturiesResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Faktury
  faktury := &models.Faktury{}

  // To find the Faktury the parameter faktury_id is used.
  if err := tx.Find(faktury, c.Param("faktury_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    c.Set("faktury", faktury)

    return c.Render(http.StatusOK, r.HTML("/fakturies/show.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(faktury))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(faktury))
  }).Respond(c)
}

// New renders the form for creating a new Faktury.
// This function is mapped to the path GET /fakturies/new
func (v FakturiesResource) New(c buffalo.Context) error {
  c.Set("faktury", &models.Faktury{})

  return c.Render(http.StatusOK, r.HTML("/fakturies/new.plush.html"))
}
// Create adds a Faktury to the DB. This function is mapped to the
// path POST /fakturies
func (v FakturiesResource) Create(c buffalo.Context) error {
  // Allocate an empty Faktury
  faktury := &models.Faktury{}

  // Bind faktury to the html form elements
  if err := c.Bind(faktury); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(faktury)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the new.html template that the user can
      // correct the input.
      c.Set("faktury", faktury)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/fakturies/new.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "faktury.created.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/fakturies/%v", faktury.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.JSON(faktury))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.XML(faktury))
  }).Respond(c)
}

// Edit renders a edit form for a Faktury. This function is
// mapped to the path GET /fakturies/{faktury_id}/edit
func (v FakturiesResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Faktury
  faktury := &models.Faktury{}

  if err := tx.Find(faktury, c.Param("faktury_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  c.Set("faktury", faktury)
  return c.Render(http.StatusOK, r.HTML("/fakturies/edit.plush.html"))
}
// Update changes a Faktury in the DB. This function is mapped to
// the path PUT /fakturies/{faktury_id}
func (v FakturiesResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Faktury
  faktury := &models.Faktury{}

  if err := tx.Find(faktury, c.Param("faktury_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Faktury to the html form elements
  if err := c.Bind(faktury); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(faktury)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the edit.html template that the user can
      // correct the input.
      c.Set("faktury", faktury)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/fakturies/edit.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "faktury.updated.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/fakturies/%v", faktury.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(faktury))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(faktury))
  }).Respond(c)
}

// Destroy deletes a Faktury from the DB. This function is mapped
// to the path DELETE /fakturies/{faktury_id}
func (v FakturiesResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Faktury
  faktury := &models.Faktury{}

  // To find the Faktury the parameter faktury_id is used.
  if err := tx.Find(faktury, c.Param("faktury_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(faktury); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a flash message
    c.Flash().Add("success", T.Translate(c, "faktury.destroyed.success"))

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "/fakturies")
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(faktury))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(faktury))
  }).Respond(c)
}

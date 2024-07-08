package main

import "snippetbox.pulkitvyas08/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

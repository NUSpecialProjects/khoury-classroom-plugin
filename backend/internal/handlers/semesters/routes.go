package semesters

import (
	"github.com/CamPlume1/khoury-classroom/internal/handlers/semesters/assignments"
	"github.com/CamPlume1/khoury-classroom/internal/handlers/semesters/file_tree"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	protected := app.Group("/semester/:semesterID")
	file_tree.Routes(protected, params)
	assignments.Routes(protected, params)
}

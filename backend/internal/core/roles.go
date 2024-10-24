package core

import "github.com/CamPlume1/khoury-classroom/internal/models"

var Prof_Role = models.OrganizationTemplateRole{
	Name:        "Professor",
	Description: "Professor",
	Permissions: []string{},
	BaseRole:    "admin",
}

var TA_Role = models.OrganizationTemplateRole{
	Name:        "TA",
	Description: "Teaching Assistant",
	Permissions: []string{},
	BaseRole:    "maintain",
}

var Student_Role = models.OrganizationTemplateRole{
	Name:        "Student",
	Description: "Student",
	Permissions: []string{},
	BaseRole:    "read",
}

var Role_Map = map[string]models.OrganizationTemplateRole{
	Prof_Role.Name:    Prof_Role,
	TA_Role.Name:      TA_Role,
	Student_Role.Name: Student_Role,
}

func GenerateRoleName(semester models.Semester, role models.OrganizationTemplateRole) string {
	return role.Name + "-" + semester.GetName()
}

func createSemesterTemplateRole(role models.OrganizationTemplateRole, semester models.Semester) models.OrganizationTemplateRole {
	return models.OrganizationTemplateRole{
		Name:        GenerateRoleName(semester, role),
		Description: role.Description + " for " + semester.GetName(),
		Permissions: role.Permissions,
		BaseRole:    role.BaseRole,
	}
}

func GetSemesterTemplateRoles(semester models.Semester) []models.OrganizationTemplateRole {
	return []models.OrganizationTemplateRole{
		Prof_Role,
		createSemesterTemplateRole(TA_Role, semester),
		createSemesterTemplateRole(Student_Role, semester),
	}
}

package models



type AutoGrade struct {
	/*Assignment ID*/
	AssignmentID uint64 `json:"assignment_id"`

	/*Number of points awarded */
	PointsAwarded int16 `json:"points_awarded"`

	/* Number of points available */
	PointsAvailable int16 `json:"points_available"`

	/* Student identifier */
	RosterIdentifier string `json:"roster_identifier"`

	/* Github Username */
	GithubUsername string `json:"github_username"`
}
CREATE TABLE IF NOT EXISTS classrooms (
    id SERIAL PRIMARY KEY,
    classroom_name VARCHAR(255) NOT NULL,
    org_id INTEGER NOT NULL,
    org_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS classroom_tokens (
    token VARCHAR(255) PRIMARY KEY, 
    expires_in TIMESTAMP NOT NULL,
    classroom_id INTEGER NOT NULL,
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id)
);

CREATE TYPE USER_ROLE AS
ENUM('PROFESSOR', 'TA', 'STUDENT');

CREATE TABLE IF NOT EXISTS user_to_classroom (
    github_username VARCHAR(255) PRIMARY KEY, 
    github_user_id INTEGER UNIQUE NOT NULL,
    role USER_ROLE NOT NULL,
    classroom_id INTEGER NOT NULL,
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id)
);

CREATE TABLE IF NOT EXISTS assignment_outlines (
    id SERIAL PRIMARY KEY,
    template_repo_owner VARCHAR(255) NOT NULL,
    template_repo_id VARCHAR(255) NOT NULL,
    created_date TIMESTAMP DEFAULT NOW(),
    released_date TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    classroom_id INTEGER NOT NULL,
    group_assignment BOOLEAN DEFAULT FALSE NOT NULL,
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id)
);

CREATE TABLE IF NOT EXISTS assignment_tokens (
    token VARCHAR(255) PRIMARY KEY,
    expires_in TIMESTAMP NOT NULL,
    assignment_outline_id INTEGER NOT NULL,
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)
);

CREATE TABLE IF NOT EXISTS rubric_items (
    id SERIAL PRIMARY KEY,
    assignment_outline_id INTEGER NOT NULL,
    point_value INTEGER NOT NULL,
    explanation VARCHAR(255) NOT NULL,
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)
);

CREATE TABLE IF NOT EXISTS student_works (
    id SERIAL PRIMARY KEY,
    assignment_outline_id INTEGER NOT NULL,
    repo_name VARCHAR(255),
    due_date TIMESTAMP NOT NULL,
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)
);

CREATE TABLE IF NOT EXISTS students_to_student_work (
    github_user_id INTEGER NOT NULL,
    student_work_id INTEGER NOT NULL,
    FOREIGN KEY (github_user_id) REFERENCES user_to_classroom(github_user_id),
    FOREIGN KEY (student_work_id) REFERENCES student_works(id)
);

CREATE TYPE GRADING_STATE AS 
ENUM('GRADING_ASSIGNED', 'GRADING_COMPLETED', 'GRADES_PUBLISHED');

CREATE TABLE IF NOT EXISTS submissions (
    id SERIAL PRIMARY KEY, 
    student_work_id INTEGER NOT NULL,
    repo_name VARCHAR(255) NOT NULL,
    grading_state GRADING_STATE NOT NULL,
    manual_feedback_score INTEGER,
    auto_grader_score INTEGER,
    submission_timestamp TIMESTAMP NOT NULL,
    grades_published_timestamp TIMESTAMP,
    pull_request_number INTEGER NOT NULL,
    FOREIGN KEY (student_work_id) REFERENCES student_works(id)
);

CREATE TABLE IF NOT EXISTS feedback_comment (
    id SERIAL PRIMARY KEY,
    submission_id INTEGER NOT NULL,
    rubric_item_id INTEGER NOT NULL,
    grader_gh_user_id INTEGER NOT NULL,
    FOREIGN KEY (submission_id) REFERENCES submissions(id),
    FOREIGN KEY (rubric_item_id) REFERENCES rubric_items(id),
    FOREIGN KEY (grader_gh_user_id) REFERENCES user_to_classroom(github_user_id)
);


CREATE TYPE REGRADE_STATE AS 
ENUM('NO_REGRADE_REQUESTED', 'REGRADE_REQUESTED', 'REGRADE_FINALIZED');

CREATE TABLE IF NOT EXISTS regrade_requests (
    id SERIAL PRIMARY KEY, 
    feedback_comment_id INTEGER NOT NULL,
    regrade_state REGRADE_STATE NOT NULL,
    student_comment VARCHAR(255) NOT NULL,
    FOREIGN KEY (feedback_comment_id) REFERENCES feedback_comment(id)

);

CREATE TABLE IF NOT EXISTS sessions (
  github_user_id INTEGER PRIMARY KEY,
  access_token VARCHAR(255) NOT NULL,
  token_type VARCHAR(255),
  refresh_token VARCHAR(255),
  expires_in INTEGER
);




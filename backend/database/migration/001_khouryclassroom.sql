CREATE TABLE IF NOT EXISTS semesters (
  classroom_id INTEGER UNIQUE PRIMARY KEY,
  org_name VARCHAR(255) NOT NULL,
  classroom_name VARCHAR(255) NOT NULL,
  active BOOLEAN NOT NULL,
  org_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS classrooms (
    id SERIAL PRIMARY KEY,
    classroom_name VARCHAR(255) NOT NULL,
    org_id INTEGER NOT NULL,
    org_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS classroom_tokens (
    token STRING PRIMARY KEY, 
    expires_in TIMESTAMP NOT NULL,
    classroom_id INTEGER NOT NULL,
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id)
);

CREATE TABLE IF NOT EXISTS user_to_classroom (
    github_username STRING PRIMARY KEY, 
    github_user_id INTEGER NOT NUll,
    role STRING NOT NULL,
    classroom_id INTEGER NOT NULL,
);

CREATE TABLE IF NOT EXISTS assignment_outlines (
    id SERIAL PRIMARY KEY,
    template_repo_owner STRING NOT NULL,
    template_repo_id STRING NOT NULL,
    created_date TIMESTAMP DEFAULT NOW(),
    name STRING NOT NULL,
    classroom_id INTEGER NOT NULL,
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id)
);

CREATE TABLE IF NOT EXISTS assignment_tokens (
    token STRING PRIMARY KEY,
    expires_in TIMESTAMP NOT NULL,
    assignment_outline_id INTEGER NOT NULL,
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)
);

CREATE TABLE IF NOT EXISTS rubric_items (
    id SERIAL PRIMARY KEY,
    assignment_outline_id INTEGER NOT NULL,
    point_value INTEGER NOT NULL,
    explanation STRING NOT NULL,
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)
);


CREATE TABLE IF NOT EXISTS student_works (
    id SERIAL PRIMARY KEY,
    assignment_outline_id INTEGER NOT NULL,
    repo_name STRING,
    due_date TIMESTAMP NOT NULL,
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)

);

CREATE TABLE IF NOT EXISTS students_to_student_work (
    github_user_id STRING NOT NULL,
    student_work_id INTEGER NOT NULL,
    FOREIGN KEY (github_user_id) REFERENCES user_to_classroom(github_user_id),
    FOREIGN KEY (student_work_id) REFERENCES student_works(id)
);

CREATE TABLE IF NOT EXISTS submissions (
    id SERIAL PRIMARY KEY, 
    student_work_id INTEGER NOT NULL,
    repo_name STRING NOT NULL,
    grading_completed BOOLEAN DEFAULT FALSE NOT NULL,
    grades_published BOOLEAN DEFAULT FALSE NOT NULL,
    manual_feedback_score INTEGER,
    auto_grader_score INTEGER,
    submission_timestamp TIMESTAMP NOT NULL,
    grades_published_timestamp TIMESTAMP,
    commit_hash STRING NOT NULL,
    FOREIGN KEY (student_work_id) REFERENCES student_works(id)
);

CREATE TABLE IF NOT EXISTS feedback_comment (
    id SERIAL PRIMARY KEY,
    submission_id INTEGER NOT NULL,
    rubric_item_id INTEGER NOT NULL,
    grader_gh_user_id INTEGER NOT NULL,
    FOREIGN KEY (submission_id) REFERENCES submissions(id)
    FOREIGN KEY (rubric_item_id) REFERENCES rubric_items(id)
    FOREIGN KEY (grader_gh_user_id) REFERENCES user_to_classroom(github_user_id)
);

CREATE TABLE IF NOT EXISTS sessions (
    github_user_id INTEGER PRIMARY KEY,
    access_token STRING NOT NULL,
    token_type STRING NOT NULL,
    refresh_token STRING NOT NULL,
    expires_in INTEGER NOT NULL
)



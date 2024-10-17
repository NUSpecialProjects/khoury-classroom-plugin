CREATE TABLE IF NOT EXISTS semesters (
  id SERIAL PRIMARY KEY,
  classroom_id INTEGER UNIQUE NOT NULL,
  active BOOLEAN NOT NULL,
  org_id INTEGER UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS rubrics (
  id SERIAL PRIMARY KEY,
  content VARCHAR(255) NOT NULL 
);

CREATE TABLE IF NOT EXISTS assignments (
  id SERIAL PRIMARY KEY,
  rubric_id INTEGER,
  classroom_id INTEGER NOT NULL,
  assignment_classroom_id INTEGER NOT NULL,
  active BOOLEAN NOT NULL, 
  FOREIGN KEY (rubric_id) REFERENCES rubrics(id)
);

CREATE TABLE IF NOT EXISTS student_assignments (
  id SERIAL PRIMARY KEY, 
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  student_gh_username VARCHAR(255) NOT NULL,
  ta_gh_username VARCHAR(255) NOT NULL,
  assignment_id INTEGER NOT NULL,
  completed BOOLEAN NOT NULL,
  started BOOLEAN NOT NULL,
  FOREIGN KEY (assignment_id) REFERENCES assignments(id)
);

CREATE TABLE IF NOT EXISTS due_dates (
  id SERIAL PRIMARY KEY,
  due TIMESTAMP DEFAULT NOW() NOT NULL,
  student_assignment_id INTEGER NOT NULL,
  FOREIGN KEY (student_assignment_id) REFERENCES student_assignments(id)
);

CREATE TABLE IF NOT EXISTS regrades (
  id SERIAL PRIMARY KEY,
  student_gh_username VARCHAR(255) NOT NULL,
  ta_gh_username VARCHAR(255) NOT NULL,
  due_date_id INTEGER NOT NULL,
  FOREIGN KEY (due_date_id) REFERENCES due_dates(id)
);

CREATE TABLE IF NOT EXISTS sessions (
  github_user_id INTEGER UNIQUE PRIMARY KEY,
  access_token VARCHAR(255) NOT NULL,
  token_type VARCHAR(255),
  refresh_token VARCHAR(255),
  expires_in INTEGER
);

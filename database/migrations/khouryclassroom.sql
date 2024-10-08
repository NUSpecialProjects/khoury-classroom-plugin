CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  role VARCHAR(50) NOT NULL,
  name VARCHAR(255) NOT NULL,
  gh_username VARCHAR(255) NOT NULL,
  joined_on TIMESTAMP
);

CREATE TABLE IF NOT EXISTS professor_ta (
  prof_id INTEGER NOT NULL,
  ta_id INTEGER,
  PRIMARY KEY (prof_id, ta_id),
  FOREIGN KEY (prof_id) REFERENCES users(id),
  FOREIGN KEY (ta_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS rubrics (
  id SERIAL PRIMARY KEY,
  content VARCHAR(255) NOT NULL 
);

CREATE TABLE IF NOT EXISTS assignment_templates (
  id SERIAL PRIMARY KEY,
  rubric_id INTEGER,
  FOREIGN KEY (rubric_id) REFERENCES rubrics(id)
);

CREATE TABLE IF NOT EXISTS assignments (
  id SERIAL PRIMARY KEY, 
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  student_gh_username VARCHAR(255) NOT NULL,
  ta_id INTEGER NOT NULL,
  github_classroom_id INTEGER NOT NULL,
  template_id INTEGER NOT NULL,
  completed BOOLEAN NOT NULL,
  started BOOLEAN NOT NULL,
  FOREIGN KEY (ta_id) REFERENCES users(id),
  FOREIGN KEY (template_id) REFERENCES assignment_templates(id)
);

CREATE TABLE IF NOT EXISTS due_dates (
  id SERIAL PRIMARY KEY,
  due TIMESTAMP DEFAULT NOW() NOT NULL,
  assignment_id INTEGER NOT NULL,
  FOREIGN KEY (assignment_id) REFERENCES assignments(id)
);

CREATE TABLE IF NOT EXISTS regrades (
  id SERIAL PRIMARY KEY,
  student_id INTEGER NOT NULL,
  ta_id INTEGER NOT NULL,
  due_date_id INTEGER NOT NULL,
  FOREIGN KEY (student_id) REFERENCES users(id),
  FOREIGN KEY (ta_id) REFERENCES users(id),
  FOREIGN KEY (due_date_id) REFERENCES due_dates(id)
);



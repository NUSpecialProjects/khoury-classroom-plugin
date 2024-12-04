CREATE TABLE IF NOT EXISTS classrooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    org_id INTEGER NOT NULL,
    org_name VARCHAR(255) NOT NULL,
    student_team_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (name, org_id)
);

DO $$ BEGIN
    CREATE TYPE USER_ROLE AS 
    ENUM('PROFESSOR', 'TA', 'STUDENT');
EXCEPTION 
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE USER_STATUS AS 
    ENUM('NOT_IN_ORG', 'REMOVED', 'REQUESTED', 'ORG_INVITED', 'ACTIVE'); -- intentionally don't have a "NONE" status, as any user in our DB has "interacted" with our system in some way
EXCEPTION 
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255), --TODO: this should be not null eventually
    last_name VARCHAR(255), --TODO: this should be not null eventually
    github_username VARCHAR(255) NOT NULL, 
    github_user_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- TODO: Impose length on tokens
CREATE TABLE IF NOT EXISTS classroom_tokens (
    token VARCHAR(255) PRIMARY KEY, 
    expires_at TIMESTAMP,
    classroom_id INTEGER NOT NULL,
    classroom_role USER_ROLE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id)
);

CREATE TABLE IF NOT EXISTS classroom_membership ( 
    user_id INTEGER NOT NULL,
    classroom_id INTEGER NOT NULL,
    classroom_role USER_ROLE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    status USER_STATUS NOT NULL, -- represents whether the user has "requested" to join the org, been invited to the org, or is in the org
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id),
    PRIMARY KEY (user_id, classroom_id)
);

CREATE TABLE IF NOT EXISTS assignment_templates (
    template_repo_id INTEGER PRIMARY KEY,
    template_repo_owner VARCHAR(255) NOT NULL,
    template_repo_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS assignment_base_repos (
    base_repo_id INTEGER PRIMARY KEY,
    base_repo_owner VARCHAR(255) NOT NULL,
    base_repo_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS rubrics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    org_id INTEGER NOT NULL,
    classroom_id INTEGER NOT NULL,
    reusable BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id)
);

CREATE TABLE IF NOT EXISTS rubric_items (
    id SERIAL PRIMARY KEY,
    rubric_id INTEGER,
    point_value INTEGER NOT NULL,
    explanation VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (rubric_id) REFERENCES rubrics(id)
);

CREATE TABLE IF NOT EXISTS assignment_outlines (
    id SERIAL PRIMARY KEY,
    template_id INTEGER,
    base_repo_id INTEGER UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    released_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    classroom_id INTEGER NOT NULL,
    rubric_id INTEGER,
    group_assignment BOOLEAN DEFAULT FALSE NOT NULL,
    main_due_date TIMESTAMP,
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id),
    FOREIGN KEY (template_id) REFERENCES assignment_templates(template_repo_id),
    FOREIGN KEY (base_repo_id) REFERENCES assignment_base_repos(base_repo_id)
);

CREATE TABLE IF NOT EXISTS assignment_outline_tokens (
    token VARCHAR(255) PRIMARY KEY, 
    expires_at TIMESTAMP,
    assignment_outline_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)
);

-- TODO: Impose length on tokens
CREATE TABLE IF NOT EXISTS assignment_tokens (
    token VARCHAR(255) PRIMARY KEY,
    expires_at TIMESTAMP,
    assignment_outline_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)
);

DO $$ BEGIN
    CREATE TYPE WORK_STATE AS 
    ENUM('ACCEPTED', 'GRADING_ASSIGNED', 'GRADING_COMPLETED', 'GRADE_PUBLISHED');
EXCEPTION 
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS student_works (
    id SERIAL PRIMARY KEY,
    assignment_outline_id INTEGER NOT NULL,
    repo_name VARCHAR(255),
    unique_due_date TIMESTAMP,
    manual_feedback_score INTEGER,
    auto_grader_score INTEGER,
    grades_published_timestamp TIMESTAMP,
    work_state WORK_STATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (assignment_outline_id) REFERENCES assignment_outlines(id)
);

CREATE TABLE IF NOT EXISTS work_contributors (
    user_id INTEGER NOT NULL,
    student_work_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (student_work_id) REFERENCES student_works(id),
    PRIMARY KEY (user_id, student_work_id)
);

CREATE TABLE IF NOT EXISTS feedback_comment (
    id SERIAL PRIMARY KEY,
    student_work_id INTEGER NOT NULL,
    rubric_item_id INTEGER NOT NULL,
    ta_user_id INTEGER NOT NULL,
    file_path VARCHAR(255),
    file_line INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (student_work_id) REFERENCES student_works(id),
    FOREIGN KEY (rubric_item_id) REFERENCES rubric_items(id),
    FOREIGN KEY (ta_user_id) REFERENCES users(id),
    -- if file path exists, enforce that file line also exists.
    -- cannot comment on an entire file (for now), only lines and entire work
    CONSTRAINT if_file_path_then_file_line
        CHECK (NOT (file_path IS NOT NULL AND file_line IS NULL))
);


DO $$ BEGIN
    CREATE TYPE REGRADE_STATE AS 
    ENUM('NO_REGRADE_REQUESTED', 'REGRADE_REQUESTED', 'REGRADE_FINALIZED');
EXCEPTION 
    WHEN duplicate_object THEN null;
END $$;


CREATE TABLE IF NOT EXISTS regrade_requests (
    id SERIAL PRIMARY KEY, 
    feedback_comment_id INTEGER NOT NULL,
    regrade_state REGRADE_STATE NOT NULL,
    student_comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (feedback_comment_id) REFERENCES feedback_comment(id)
);

CREATE TABLE IF NOT EXISTS sessions (
    github_user_id INTEGER PRIMARY KEY,
    access_token VARCHAR(255) NOT NULL,
    token_type VARCHAR(255),
    refresh_token VARCHAR(255),
    expires_in INTEGER,
    created_at TIMESTAMP DEFAULT NOW()
);

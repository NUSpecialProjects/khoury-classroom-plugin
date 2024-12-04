-- Insert into classrooms
INSERT INTO classrooms (id, name, org_id, org_name, created_at, student_team_name)
VALUES
(1, 'Spring 2025', 182810684, 'NUSpecialProjects', NOW(), 'spring-2025-students-MOCK'),
(2, 'Fall 2025', 182810684, 'NUSpecialProjects', NOW(), 'fall-2025-students-MOCK');
SELECT setval('classrooms_id_seq', (SELECT MAX(id) FROM classrooms));

-- Create users
INSERT INTO users (id, first_name, last_name, github_username, github_user_id)
VALUES
(1, 'Kenny', 'Chen', 'kennybc', 54950614),
(2, 'Alex', 'Angione', 'alexangione419', 111721125),
(3, 'Nick', 'Tietje2', 'NickTietje', 183017928),
(4, 'Seby', 'Tremblay', 'sebytremblay', 91509344),
(5, 'Cam', 'Plume', 'CamPlume1', 116120547),
(6, 'Nick', 'Tietje', 'ntietje1', 124538220),
(7, 'Nandini', 'Ghosh', 'nandini-ghosh', 93226556),
(8, 'Mark', 'Fontenot', 'MarkFontenot', 1777629);
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));

-- Insert into classroom_membership
INSERT INTO classroom_membership (user_id, classroom_id, classroom_role, created_at, status)
VALUES
(1, 1, 'TA', NOW(), 'ACTIVE'),
(2, 1, 'PROFESSOR', NOW(), 'ACTIVE'),
(3, 1, 'STUDENT', NOW(), 'ACTIVE'),
(4, 1, 'TA', NOW(), 'ACTIVE'),
(5, 1, 'TA', NOW(), 'ACTIVE'),
(6, 1, 'TA', NOW(), 'ACTIVE'),
(7, 1, 'TA', NOW(), 'ACTIVE'),
(8, 1, 'PROFESSOR', NOW(), 'ACTIVE'),
(1, 2, 'STUDENT', NOW(), 'ACTIVE'),
(2, 2, 'PROFESSOR', NOW(), 'ACTIVE'),
(3, 2, 'TA', NOW(), 'ACTIVE'),
(4, 2, 'STUDENT', NOW(), 'ACTIVE'),
(5, 2, 'STUDENT', NOW(), 'ACTIVE'),
(6, 2, 'STUDENT', NOW(), 'ACTIVE'),
(7, 2, 'STUDENT', NOW(), 'ACTIVE'),
(8, 2, 'PROFESSOR', NOW(), 'ACTIVE');

-- Insert into rubrics
INSERT INTO rubrics (id, name, org_id, classroom_id, reusable) VALUES 
(1, 'Generic Assignment Rubric', 1, 1, true);
SELECT setval('rubrics_id_seq', (SELECT MAX(id) FROM rubrics));

-- Insert into rubric_items
INSERT INTO rubric_items (id, rubric_id, point_value, explanation, created_at)
VALUES
(1, 1, 1, 'The code works well', NOW()),
(2, 1, -1, 'The code is really bad', NOW()),
(3, 1, 0, 'You wrote code', NOW());
SELECT setval('rubric_items_id_seq', (SELECT MAX(id) FROM rubric_items));


-- Insert into classrooms
INSERT INTO classrooms (id, name, org_id, org_name, created_at, student_team_name)
VALUES
(1, 'Spring 2025', 182810684, 'NUSpecialProjects', NOW(), 'kennys-coding-classroom-students'),
SELECT setval('classrooms_id_seq', (SELECT MAX(id) FROM classrooms));

-- Create users (Kenny and others)
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
(8, 1, 'PROFESSOR', NOW(), 'ACTIVE');

-- Insert into rubrics
INSERT INTO rubrics (id, name, org_id, classroom_id, reusable) VALUES 
(1, 'Assignment 1 Rubric', 1, 1, true);
SELECT setval('rubrics_id_seq', (SELECT MAX(id) FROM rubrics));

-- Insert into rubric_items
INSERT INTO rubric_items (id, rubric_id, point_value, explanation, created_at)
VALUES
(1, 1, 1, 'The code works well', NOW()),
(2, 1, -1, 'The code is really bad', NOW()),
(3, 1, 0, 'You wrote code', NOW());
SELECT setval('rubric_items_id_seq', (SELECT MAX(id) FROM rubric_items));

-- Insert into assignment_templates
INSERT INTO assignment_templates (template_repo_owner, template_repo_id, created_at, template_repo_name)
VALUES
()


-- Insert into assignment_outlines
INSERT INTO assignment_outlines (id, template_id, base_repo_id, created_at, released_at, name, rubric_id, classroom_id, group_assignment)
VALUES
(1, 1001, 1000, NOW(), '2023-01-01 09:00:00', 'Running and Chocolate Tracker App', 1, 1, FALSE);
SELECT setval('assignment_outlines_id_seq', (SELECT MAX(id) FROM assignment_outlines));

-- Insert into student_works
INSERT INTO student_works (id, assignment_outline_id, repo_name, unique_due_date, manual_feedback_score, auto_grader_score, grades_published_timestamp, work_state, created_at)
VALUES
(1, 1, 'khoury-classroom-plugin', '2023-02-01 23:59:59', 28, 20, '2023-02-05 10:00:00', 'GRADE_PUBLISHED', NOW()),
(2, 2, 'kennysmith/compiler-design', '2023-03-01 23:59:59', 25, 22, '2023-03-05 12:00:00', 'GRADE_PUBLISHED', NOW()),
(3, 3, 'alanturing/linux-module', '2023-04-01 23:59:59', 40, 35, '2023-04-05 14:00:00', 'GRADE_PUBLISHED', NOW()),
(4, 4, 'adalovelace/encryption-algorithms', '2023-05-01 23:59:59', 45, 40, '2023-05-05 16:00:00', 'GRADE_PUBLISHED', NOW()),
(5, 5, 'mhamilton/mobile-app', '2023-06-01 23:59:59', 18, 15, '2023-06-05 18:00:00', 'GRADE_PUBLISHED', NOW()),
(6, 6, 'kennysmith/ai-chatbot', '2023-07-01 23:59:59', 50, 45, '2023-07-05 20:00:00', 'GRADE_PUBLISHED', NOW()),
(7, 7, 'bliskov/website-dev', '2023-08-01 23:59:59', 28, 25, '2023-08-05 22:00:00', 'GRADE_PUBLISHED', NOW()),
(8, 8, 'linustorvalds/cloud-deployment', '2023-09-01 23:59:59', 40, 35, '2023-09-05 12:00:00', 'GRADE_PUBLISHED', NOW()),
(9, 9, 'mhamilton/cybersecurity-analysis', '2023-10-01 23:59:59', 48, 45, '2023-10-05 14:00:00', 'GRADE_PUBLISHED', NOW()),
(10, 10, 'kennysmith/software-design-patterns', '2023-11-01 23:59:59', 38, 35, '2023-11-05 16:00:00', 'GRADE_PUBLISHED', NOW()),
(11, 1, 'kenny-assignment-josevaca1231', '2023-04-01 23:59:59', 40, 35, '2023-04-05 14:00:00', 'GRADE_PUBLISHED', NOW()),
(12, 1, 'kenny-assignment-josevaca1231', '2023-04-01 23:59:59', 40, 35, '2023-04-05 14:00:00', 'GRADE_PUBLISHED', NOW());
SELECT setval('student_works_id_seq', (SELECT MAX(id) FROM student_works));

-- Insert into work_contributors+
INSERT INTO work_contributors (user_id, student_work_id, created_at)
VALUES
(1, 1, NOW()),
(2, 1, NOW()),
(3, 1, NOW()),
(1, 2, NOW()),
(2, 2, NOW()),
(3, 3, NOW()),
(4, 4, NOW()),
(6, 5, NOW()),
(1, 6, NOW()),
(8, 7, NOW()),
(5, 8, NOW()),
(6, 9, NOW()),
(1, 10, NOW()),
(3, 11, NOW()),
(4, 12, NOW());

-- Insert into feedback_comment
INSERT INTO feedback_comment (id, student_work_id, rubric_item_id, ta_user_id, created_at)
VALUES
(1, 1, 1, 1, NOW()),
(2, 1, 2, 1, NOW()),
(3, 1, 3, 1, NOW()),
(4, 2, 4, 1, NOW()),
(5, 2, 5, 1, NOW()),
(6, 3, 6, 1, NOW()),
(7, 3, 7, 1, NOW()),
(8, 4, 8, 1, NOW()),
(9, 4, 9, 1, NOW()),
(10, 5, 10, 1, NOW()),
(11, 5, 11, 1, NOW()),
(12, 6, 12, 1, NOW()),
(13, 7, 13, 1, NOW()),
(14, 7, 14, 1, NOW()),
(15, 8, 15, 1, NOW()),
(16, 9, 16, 1, NOW()),
(17, 10, 17, 1, NOW());
SELECT setval('feedback_comment_id_seq', (SELECT MAX(id) FROM feedback_comment));

-- Insert into regrade_requests
INSERT INTO regrade_requests (id, feedback_comment_id, regrade_state, student_comment, created_at)
VALUES
(1, 2, 'REGRADE_REQUESTED', 'I believe my app correctly records chocolate consumption. Could you please re-evaluate?', NOW()),
(2, 5, 'REGRADE_REQUESTED', 'Could you check the parsing module again? I made some updates.', NOW()),
(3, 9, 'REGRADE_REQUESTED', 'I think my encryption algorithm meets the efficiency standards.', NOW()),
(4, 11, 'REGRADE_REQUESTED', 'The GPS tracking should be accurate now.', NOW()),
(5, 13, 'REGRADE_REQUESTED', 'I improved the website responsiveness.', NOW()),
(6, 16, 'REGRADE_REQUESTED', 'Found additional vulnerabilities, please review.', NOW());
SELECT setval('regrade_requests_id_seq', (SELECT MAX(id) FROM regrade_requests));


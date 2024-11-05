-- Insert into classrooms
INSERT INTO classrooms (name, org_id, org_name, created_at)
VALUES
('Kennys Coding Classroom', 98765, 'KennyCodeOrg', NOW()),
('Advanced Running Analytics', 98766, 'RunTechOrg', NOW()),
('Chocolate Lovers Unite', 98767, 'ChocoOrg', NOW()),
('Data Structures and Algorithms', 98768, 'CodeMasters', NOW()),
('Mobile App Development', 98769, 'AppDevs', NOW()),
('AI and Machine Learning', 98770, 'AIMLGroup', NOW()),
('Web Development Bootcamp', 98771, 'WebCoders', NOW()),
('Cloud Computing', 98772, 'CloudExperts', NOW()),
('Cybersecurity Fundamentals', 98773, 'SecureNet', NOW()),
('Software Engineering Principles', 98774, 'SoftEngOrg', NOW());

-- Create users (Kenny and others)
INSERT INTO users (first_name, last_name, github_username, github_user_id, role)
VALUES
('Kenny', 'Smith', 'kennysmith', 123456, 'STUDENT'),
('Grace', 'Hopper', 'gracehopper', 789012, 'PROFESSOR'),
('Alan', 'Turing', 'alanturing', 345678, 'STUDENT'),
('Ada', 'Lovelace', 'adalovelace', 901234, 'STUDENT'),
('Linus', 'Torvalds', 'linustorvalds', 567890, 'TA'),
('Margaret', 'Hamilton', 'mhamilton', 234567, 'STUDENT'),
('Tim', 'Berners-Lee', 'timbl', 890123, 'PROFESSOR'),
('Barbara', 'Liskov', 'bliskov', 678901, 'STUDENT'),
('Dennis', 'Ritchie', 'dritchie', 112233, 'PROFESSOR'),
('Ken', 'Thompson', 'kthompson', 445566, 'TA');

-- Insert into classroom_membership
INSERT INTO classroom_membership (user_id, classroom_id, created_at)
VALUES
(1, 1, NOW()),
(2, 1, NOW()),
(3, 1, NOW()),
(4, 2, NOW()),
(5, 2, NOW()),
(6, 3, NOW()),
(7, 3, NOW()),
(8, 4, NOW()),
(9, 4, NOW()),
(10, 5, NOW()),
(1, 6, NOW()),
(3, 7, NOW()),
(6, 8, NOW()),
(8, 9, NOW()),
(5, 10, NOW());

-- Insert into assignment_template
INSERT INTO assignment_template (template_repo_owner, template_repo_id, created_at)
VALUES
('kennysmith', '1000', NOW()),
('gracehopper', '1001', NOW()),
('linustorvalds', '1002', NOW()),
('alanturing', '1003', NOW()),
('adalovelace', '1004', NOW()),
('mhamilton', '1005', NOW()),
('timbl', '1006', NOW()),
('bliskov', '1007', NOW()),
('dritchie', '1008', NOW()),
('kthompson', '1009', NOW());

-- Insert into assignment_outlines
INSERT INTO assignment_outlines (template_id, created_at, released_at, name, classroom_id, group_assignment)
VALUES
(1, NOW(), '2023-01-01 09:00:00', 'Running and Chocolate Tracker App', 1, FALSE),
(2, NOW(), '2023-02-01 09:00:00', 'Compiler Design', 1, TRUE),
(3, NOW(), '2023-03-01 09:00:00', 'Linux Kernel Module', 2, FALSE),
(4, NOW(), '2023-04-01 09:00:00', 'Encryption Algorithms', 2, TRUE),
(5, NOW(), '2023-05-01 09:00:00', 'Mobile App for Runners', 3, FALSE),
(6, NOW(), '2023-06-01 09:00:00', 'AI Chatbot', 4, TRUE),
(7, NOW(), '2023-07-01 09:00:00', 'Website Development', 5, FALSE),
(8, NOW(), '2023-08-01 09:00:00', 'Cloud Deployment', 6, TRUE),
(9, NOW(), '2023-09-01 09:00:00', 'Cybersecurity Analysis', 7, FALSE),
(10, NOW(), '2023-10-01 09:00:00', 'Software Design Patterns', 8, TRUE);


-- Insert into rubric_items
INSERT INTO rubric_items (assignment_outline_id, point_value, explanation, created_at)
VALUES
(1, 10, 'Tracks running distances accurately', NOW()),
(1, 10, 'Records chocolate consumption correctly', NOW()),
(1, 10, 'Code is well-documented and clean', NOW()),
(2, 15, 'Lexical analysis implemented', NOW()),
(2, 15, 'Parsing and syntax tree generation', NOW()),
(3, 20, 'Kernel module loads without errors', NOW()),
(3, 20, 'Module performs expected operations', NOW()),
(4, 25, 'Encryption algorithm efficiency', NOW()),
(4, 25, 'Security level meets standards', NOW()),
(5, 10, 'User interface is intuitive', NOW()),
(5, 10, 'GPS tracking is accurate', NOW()),
(6, 30, 'AI chatbot responds correctly', NOW()),
(7, 15, 'Website layout is responsive', NOW()),
(7, 15, 'Accessibility standards are met', NOW()),
(8, 20, 'Cloud deployment is successful', NOW()),
(9, 25, 'Identified security vulnerabilities', NOW()),
(10, 20, 'Implemented design patterns correctly', NOW());

-- Insert into student_works
INSERT INTO student_works (assignment_outline_id, repo_name, due_date, submitted_pr_number, manual_feedback_score, auto_grader_score, submission_timestamp, grades_published_timestamp, work_state, created_at)
VALUES
(1, 'kennysmith/running-chocolate-tracker', '2023-02-01 23:59:59', 10, 28, 20, '2023-01-31 20:00:00', '2023-02-05 10:00:00', 'GRADE_PUBLISHED', NOW()),
(2, 'kennysmith/compiler-design', '2023-03-01 23:59:59', 11, 25, 22, '2023-02-28 18:00:00', '2023-03-05 12:00:00', 'GRADE_PUBLISHED', NOW()),
(3, 'alanturing/linux-module', '2023-04-01 23:59:59', 12, 40, 35, '2023-03-30 22:00:00', '2023-04-05 14:00:00', 'GRADE_PUBLISHED', NOW()),
(4, 'adalovelace/encryption-algorithms', '2023-05-01 23:59:59', 13, 45, 40, '2023-04-29 19:00:00', '2023-05-05 16:00:00', 'GRADE_PUBLISHED', NOW()),
(5, 'mhamilton/mobile-app', '2023-06-01 23:59:59', 14, 18, 15, '2023-05-31 21:00:00', '2023-06-05 18:00:00', 'GRADE_PUBLISHED', NOW()),
(6, 'kennysmith/ai-chatbot', '2023-07-01 23:59:59', 15, 50, 45, '2023-06-30 20:00:00', '2023-07-05 20:00:00', 'GRADE_PUBLISHED', NOW()),
(7, 'bliskov/website-dev', '2023-08-01 23:59:59', 16, 28, 25, '2023-07-31 23:00:00', '2023-08-05 22:00:00', 'GRADE_PUBLISHED', NOW()),
(8, 'linustorvalds/cloud-deployment', '2023-09-01 23:59:59', 17, 40, 35, '2023-08-30 20:00:00', '2023-09-05 12:00:00', 'GRADE_PUBLISHED', NOW()),
(9, 'mhamilton/cybersecurity-analysis', '2023-10-01 23:59:59', 18, 48, 45, '2023-09-29 21:00:00', '2023-10-05 14:00:00', 'GRADE_PUBLISHED', NOW()),
(10, 'kennysmith/software-design-patterns', '2023-11-01 23:59:59', 19, 38, 35, '2023-10-31 20:00:00', '2023-11-05 16:00:00', 'GRADE_PUBLISHED', NOW());

-- Insert into work_contributors
INSERT INTO work_contributors (user_id, student_work_id, created_at)
VALUES
(1, 1, NOW()),
(1, 2, NOW()),
(3, 3, NOW()),
(4, 4, NOW()),
(6, 5, NOW()),
(1, 6, NOW()),
(8, 7, NOW()),
(5, 8, NOW()),
(6, 9, NOW()),
(1, 10, NOW());

-- Insert into feedback_comment
INSERT INTO feedback_comment (student_work_id, rubric_item_id, grader_gh_user_id, created_at)
VALUES
(1, 1, 789012, NOW()),
(1, 2, 789012, NOW()),
(1, 3, 789012, NOW()),
(2, 4, 789012, NOW()),
(2, 5, 789012, NOW()),
(3, 6, 789012, NOW()),
(3, 7, 789012, NOW()),
(4, 8, 789012, NOW()),
(4, 9, 789012, NOW()),
(5, 10, 789012, NOW()),
(5, 11, 789012, NOW()),
(6, 12, 789012, NOW()),
(7, 13, 789012, NOW()),
(7, 14, 789012, NOW()),
(8, 15, 789012, NOW()),
(9, 16, 789012, NOW()),
(10, 17, 789012, NOW());

-- Insert into regrade_requests
INSERT INTO regrade_requests (feedback_comment_id, regrade_state, student_comment, created_at)
VALUES
(2, 'REGRADE_REQUESTED', 'I believe my app correctly records chocolate consumption. Could you please re-evaluate?', NOW()),
(5, 'REGRADE_REQUESTED', 'Could you check the parsing module again? I made some updates.', NOW()),
(9, 'REGRADE_REQUESTED', 'I think my encryption algorithm meets the efficiency standards.', NOW()),
(11, 'REGRADE_REQUESTED', 'The GPS tracking should be accurate now.', NOW()),
(13, 'REGRADE_REQUESTED', 'I improved the website responsiveness.', NOW()),
(16, 'REGRADE_REQUESTED', 'Found additional vulnerabilities, please review.', NOW());

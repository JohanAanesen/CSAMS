INSERT INTO users (user_login, user_pass, user_email)
VALUES
('sveiad', 'password', 'sveiad@stud.ntnu.no'),
('christopher', 'password', 'christopher.frantz@ntnu.no'),
('mariusz', 'password', 'mariusz.nowostawski@ntnu.no');

INSERT INTO usermeta (user_ID, meta_key, meta_value)
VALUES
(1, 'role', 'student'),
(2, 'role', 'teacher'),
(3, 'role', 'teacher');

INSERT INTO courses (course_code, course_name, course_description, course_year, course_semester)
VALUES
('IMT2681', 'Cloud Technologies', '', '2019', 'H');

INSERT INTO assignments (assignment_course_ID, assignment_name, assignment_description, assignment_publish, assignment_deadline)
VALUES
(1, 'Heightmaps', 'Make heightmaps', '2019-02-14 12:00:00', '2019-03-01 23:59:59');

INSERT INTO reviewforms (reviewform_assignment_ID)
VALUES
(1);

INSERT INTO reviewformmeta (reviewform_ID, meta_key, meta_value)
VALUES
(1, 'field', '{"type":"textfield","name":"Enter your student identifier","description":"This is the one you use to log onto FEIDE. Ensure to get that right, so we can match your response with your exam.","weight":1,"order":1}'),
(1, 'field', '{"type":"radio","name":"Task 1: Did you load the heightmap successfully?","description":"If you selected \"Partially\", please clarify in the Comments field. Mark only one oval.","weight":1,"order":2,"choices":["Yes","Partially","No"]}');
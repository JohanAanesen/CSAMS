-- Adminer 4.7.1 MySQL dump

SET NAMES utf8;
SET time_zone = '+01:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

INSERT INTO `adminfaq` (`id`, `timestamp`, `questions`) VALUES
(1, '2019-03-06 15:53:00', 'Q: How do I make a course?\n--------------------------------\n**A:** Dashboard -> Courses -> new. And create the course there\n\nQ: How do I invite students to course?\n--------------------------------\n**A:** You go to [admin/course](/admin/course) or [admin/](/admin) and on the course card, click the copy button to get the `join course through link` and send that to all students in preferred way (ex: email)\n\nQ: How do I make an assignment?\n--------------------------------\n**A:** Dashboard -> Assignments-> new. And create the assignment there');

INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (1, 'Ken Thompson', 'hei@gmail.com', 1, 'mannen@harmannenfalt.no', '$2a$14$MZj24p41j2NNGn6JDsQi0OsDb56.0LcfrIdgjE6WmZzp58O6V/VhK');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (2, 'Frode Haug', 'frodehg@teach.ntnu.no', 1, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (3, 'Ola Nordmann', 'olanor@stud.ntnu.no', 1, 'swag-meister69@ggmail.com', '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (4, 'Johan Klausen', 'johkl@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (5, 'Stian Fjerdingstad', 'stianfj@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (6, 'Svein Nilsen', 'sveini@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (7, 'Kjell Are-Kjelterud', 'kjellak@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (8, 'Marius Lillevik', 'mariuslil@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (9, 'Jorun Skaalnes', 'jorunska@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO users (id, name, email_student, teacher, email_private, password) VALUES (10, 'Klaus Aanesen', 'klausaa@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');

INSERT INTO course (id, hash, coursecode, coursename, teacher, description, year, semester) VALUES (1, 'bi4d2164gh0gbb7r94qg', 'IMT2681', 'Cloud Technologies', 3, '# IMT2681 Cloud Technologies', 2019, 'fall');

INSERT INTO usercourse (userid, courseid) VALUES (3, 1);
INSERT INTO usercourse (userid, courseid) VALUES (4, 1);
INSERT INTO usercourse (userid, courseid) VALUES (5, 1);
INSERT INTO usercourse (userid, courseid) VALUES (6, 1);
INSERT INTO usercourse (userid, courseid) VALUES (7, 1);
INSERT INTO usercourse (userid, courseid) VALUES (8, 1);
INSERT INTO usercourse (userid, courseid) VALUES (9, 1);
INSERT INTO usercourse (userid, courseid) VALUES (10, 1);

INSERT INTO `forms` (`id`, `prefix`, `name`, `created`) VALUES
(1, 'git_repository_w_comment', 'Git Repository w/Comment', '2019-03-22 08:10:16'),
(2, 'lab_1_cloud_2019', 'Lab 1 - Cloud 2019', '2019-03-22 08:14:42');

INSERT INTO `fields` (`id`, `form_id`, `type`, `name`, `description`, `label`, `hasComment`, `priority`, `weight`, `choices`) VALUES
(1, 1, 'url', 'git_repository_w_comment_0', 'Make sure it is public', 'Git Repository', 1, 0, 0, ''),
(2, 2, 'paragraph', 'lab_1_cloud_2019_1', 'Before you start filling out, just know that you are a beatiful person! I love you', 'Some information', 0, 0, 0, ''),
(3, 2, 'checkbox', 'lab_1_cloud_2019_0', '', 'Does the program compile?', 0, 1, 1, ''),
(4, 2, 'radio', 'lab_1_cloud_2019_2', 'Mark the closest one.\nEg.: 84,5% -> 80%', 'Test covarage', 0, 2, 4, '10%,20%,30%,40%,50%,60%,70%,80%,90%,100%'),
(5, 2, 'checkbox', 'lab_1_cloud_2019_3', 'Does the request work?', 'GET /api', 0, 3, 1, ''),
(6, 2, 'text', 'lab_1_cloud_2019_4', 'Is the duration fomratted in ISO 20915092835209 ?', 'GET /api', 0, 4, 1, ''),
(7, 2, 'text', 'lab_1_cloud_2019_5', 'Does it work?', 'GET /api/foo', 0, 5, 1, ''),
(8, 2, 'text', 'lab_1_cloud_2019_6', 'Does it work?', 'POST /api/foo', 0, 6, 1, '');

INSERT INTO `reviews` (`id`, `form_id`) VALUES
(1, 2);

INSERT INTO `submissions` (`id`, `form_id`) VALUES
(1, 1);

INSERT INTO `assignments` (`id`, `name`, `description`, `created`, `publish`, `deadline`, `course_id`, `submission_id`, `review_id`, `review_deadline`, `validation_id`, `reviewers`) VALUES
(1, 'Lab 1', 'Make easiest rest api', '2019-03-22 09:15:15', '2019-03-22 10:00:00', '2019-03-23 16:00:00', 1, 1, 1, '2019-03-23 17:00:00' NULL, 2);
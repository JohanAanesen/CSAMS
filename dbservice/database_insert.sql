INSERT INTO `adminfaq` (`id`, `timestamp`, `questions`)
VALUES (1, '2019-03-06 15:53:00',
        'Q: How do I make a course?\n--------------------------------\n**A:** Dashboard -> Courses -> new. And create the course there\n\nQ: How do I invite students to course?\n--------------------------------\n**A:** You go to [admin/course](/admin/course) or [admin/](/admin) and on the course card, click the copy button to get the `join course through link` and send that to all students in preferred way (ex: email)\n\nQ: How do I make an assignment?\n--------------------------------\n**A:** Dashboard -> Assignments-> new. And create the assignment there');


INSERT INTO `users` (`id`, `name`, `email_student`, `teacher`, `email_private`, `password`)
VALUES (1, 'Test User', 'hei@gmail.com', 1, 'test@yahoo.com',
        '$2a$14$MZj24p41j2NNGn6JDsQi0OsDb56.0LcfrIdgjE6WmZzp58O6V/VhK'),
       (2, 'Frode Haug', 'frodehg@teach.ntnu.no', 1, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (3, 'Ola Nordmann', 'olanor@stud.ntnu.no', 1, 'swag-meister69@ggmail.com',
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (4, 'Johan Klausen', 'johkl@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (5, 'Stian Fjerdingstad', 'stianfj@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (6, 'Svein Nilsen', 'sveini@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (7, 'Kjell Are-Kjelterud', 'kjellak@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (8, 'Marius Lillevik', 'mariuslil@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (9, 'Jorun Skaalnes', 'jorunska@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (10, 'Klaus Aanesen', 'klausaa@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');

INSERT INTO `course` (`id`, `hash`, `coursecode`, `coursename`, `teacher`, `description`, `year`, `semester`)
VALUES (1, '3876438629b786', 'IMT1031', 'Grunnleggende Programmering', 2, 'Write hello, world in C++', 2019, 'fall'),
       (2, '12387teg817eg18', 'IMT1082', 'Objekt-orientert programmering', 2, 'Write Wazz up world in Python', 2019,
        'fall'),
       (3, '12e612eg1e17ge1', 'IMT2021', 'Algoritmiske metoder', 2, 'Object orientation and algorithmic methods in C#',
        2019, 'spring');

INSERT INTO `usercourse` (`userid`, `courseid`)
VALUES (1, 3),
       (2, 3),
       (3, 1),
       (3, 3),
       (4, 2),
       (4, 3),
       (5, 3),
       (6, 3),
       (7, 3),
       (8, 3),
       (9, 3),
       (10, 3);

INSERT INTO `forms` (`id`, `prefix`, `name`, `description`, `created`)
VALUES (1, 'github_form', 'Github form', 'Form to Github', '2019-02-28 15:23:23');

INSERT INTO `fields` (`id`, `form_id`, `type`, `name`, `label`, `description`, `priority`, `weight`, `choices`)
VALUES (1, 1, 'text', 'github_form_text_0', 'Github handle', 'Username on github', 0, 0, ''),
       (2, 1, 'url', 'github_form_url_1', 'Github url', 'url to github', 1, 0, ''),
       (3, 1, 'textarea', 'github_form_textarea_2', 'Comments', 'Comments about your work', 2, 0, '');

INSERT INTO `reviews` (`id`, `form_id`) VALUES
(1, 1);

INSERT INTO `submissions` (`id`, `form_id`)
VALUES (1, 1);

INSERT INTO `assignments` (`id`, `name`, `description`, `created`, `publish`, `deadline`, `course_id`, `submission_id`, `review_id`)
VALUES (1, 'Assignment 1', '# this is an assignment\r\n* sub\r\n* 2\r\n* pew', '2019-02-28 15:25:10',
        '2019-02-28 16:24:00', '2019-03-29 11:11:00', 3, 1, NULL);

INSERT INTO `peer_reviews` (`id`, `submission_id`, `assignment_id`, `user_id`, `review_user_id`)
VALUES (1, 1, 1, 3, 4),
       (2, 1, 1, 3, 5),
       (3, 1, 1, 9, 4);

INSERT INTO `user_reviews` (`id`, `user_reviewer`, `user_target`, `review_id`, `assignment_id`, `type`, `name`, `label`, `answer`, `submitted`) VALUES
(3, 3, 4, 1, 1, 'text', 'test_text_0', 'A', 'good', '2019-03-08 01:08:43'),
(4, 3, 4, 1, 1, 'text', 'test_text_1', 'B', 'bad', '2019-03-08 01:08:43');

INSERT INTO `user_submissions` (`id`, `user_id`, `assignment_id`, `submission_id`, `type`, `answer`, `submitted`)
VALUES (1, 4, 1, 1, 'text', 'JohanKlausen', CURRENT_TIMESTAMP),
       (2, 4, 1, 1, 'url', 'https://github.com/JohanKlausen/yeet', CURRENT_TIMESTAMP),
       (3, 4, 1, 1, 'textarea', 'I did good!', CURRENT_TIMESTAMP),
       (4, 5, 1, 1, 'text', 'StianFjerdingstad', '2019-03-01 23:59:59'),
       (5, 5, 1, 1, 'url', 'https://github.com/StianFjerdingstad/Sudoku', '2019-03-01 23:59:59'),
       (6, 5, 1, 1, 'textarea', 'I did sexy good!', '2019-03-01 23:59:59'),
       (7, 10, 1, 1, 'text', 'KlausAanesen', '2019-02-28 15:23:23'),
       (8, 10, 1, 1, 'url', 'https://github.com/KlausAanesen/1337yeet420', '2019-02-28 15:23:23'),
       (9, 10, 1, 1, 'textarea', 'I did bad :(', '2019-02-28 15:23:23');

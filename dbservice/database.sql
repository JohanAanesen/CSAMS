CREATE DATABASE IF NOT EXISTS cs53 COLLATE = utf8_general_ci;

USE cs53;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+01:00";
-- Norwegian time zone! --

--
-- Database: `cs53`
--

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `adminfaq`
--

CREATE TABLE `adminfaq`
(
  `id`        int(11)                     NOT NULL,
  `timestamp` datetime                    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `questions` text COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `adminfaq`
--

INSERT INTO `adminfaq` (`id`, `timestamp`, `questions`)
VALUES (1, '2019-03-06 15:53:00',
        'Q: How do I make a course?\n--------------------------------\n**A:** Dashboard -> Courses -> new. And create the course there\n\nQ: How do I invite students to course?\n--------------------------------\n**A:** You go to [admin/course](/admin/course) or [admin/](/admin) and on the course card, click the copy button to get the `join course through link` and send that to all students in preferred way (ex: email)\n\nQ: How do I make an assignment?\n--------------------------------\n**A:** Dashboard -> Assignments-> new. And create the assignment there');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `assignments`
--

CREATE TABLE `assignments`
(
  `id`            int(11)     NOT NULL,
  `name`          varchar(64) NOT NULL,
  `description`   text,
  `created`       timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `publish`       datetime    NOT NULL,
  `deadline`      datetime    NOT NULL,
  `course_id`     int(11)     NOT NULL,
  `submission_id` int(11)              DEFAULT NULL,
  `review_id`     int(11)              DEFAULT NULL,
  `validation_id` int(11)     NULL,
  `reviewers`     int(11)     NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `assignments`
--

INSERT INTO `assignments` (`id`, `name`, `description`, `created`, `publish`, `deadline`, `course_id`, `submission_id`,
                           `review_id`)
VALUES (1, 'Assignment 1', '# this is an assignment\r\n* sub\r\n* 2\r\n* pew', '2019-02-28 15:25:10',
        '2019-02-28 16:24:00', '2019-03-29 11:11:00', 3, 1, NULL);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `course`
--

CREATE TABLE `course`
(
  `id`          int(11)                                       NOT NULL,
  `hash`        varchar(64) COLLATE utf8_danish_ci            NOT NULL,
  `coursecode`  varchar(10) COLLATE utf8_danish_ci            NOT NULL,
  `coursename`  varchar(64) COLLATE utf8_danish_ci            NOT NULL,
  `teacher`     int(11)                                       NOT NULL,
  `description` text                                          NOT NULL,
  `year`        int(11)                                       NOT NULL,
  `semester`    enum ('fall','spring') COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `course`
--

INSERT INTO `course` (`id`, `hash`, `coursecode`, `coursename`, `teacher`, `description`, `year`, `semester`)
VALUES (1, '3876438629b786', 'IMT1031', 'Grunnleggende Programmering', 2, 'Write hello, world in C++', 2019, 'fall'),
       (2, '12387teg817eg18', 'IMT1082', 'Objekt-orientert programmering', 2, 'Write Wazz up world in Python', 2019,
        'fall'),
       (3, '12e612eg1e17ge1', 'IMT2021', 'Algoritmiske metoder', 2, 'Object orientation and algorithmic methods in C#',
        2019, 'spring');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `fields`
--

CREATE TABLE `fields`
(
  `id`          int(11)     NOT NULL,
  `form_id`     int(11)     NOT NULL,
  `type`        varchar(64) NOT NULL,
  `name`        varchar(64) NOT NULL,
  `label`       varchar(64) DEFAULT NULL,
  `description` text        NOT NULL,
  `priority`    int(11)     NOT NULL,
  `weight`      int(11)     DEFAULT NULL,
  `choices`     varchar(64) DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `fields`
--

INSERT INTO `fields` (`id`, `form_id`, `type`, `name`, `label`, `description`, `priority`, `weight`, `choices`)
VALUES (1, 1, 'text', 'github_form_text_0', 'Github handle', 'Username on github', 0, 0, ''),
       (2, 1, 'url', 'github_form_url_1', 'Github url', 'url to github', 1, 0, ''),
       (3, 1, 'textarea', 'github_form_textarea_2', 'Comments', 'Comments about your work', 2, 0, '');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `forms`
--

CREATE TABLE `forms`
(
  `id`          int(11)     NOT NULL,
  `prefix`      varchar(64) NOT NULL,
  `name`        varchar(64)      DEFAULT NULL,
  `description` text,
  `created`     timestamp   NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `forms`
--

INSERT INTO `forms` (`id`, `prefix`, `name`, `description`, `created`)
VALUES (1, 'github_form', 'Github form', 'Form to Github', '2019-02-28 15:23:23');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `logs`
--

CREATE TABLE `logs`
(
  `userid`       int(11)                            NOT NULL,
  `timestamp`    datetime                           NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `activity`     varchar(32) COLLATE utf8_danish_ci NOT NULL,
  `assignmentid` int(11)                                     DEFAULT NULL,
  `courseid`     int(11)                                     DEFAULT NULL,
  `submissionid` int(11)                                     DEFAULT NULL,
  `oldvalue`     text COLLATE utf8_danish_ci,
  `newValue`     text COLLATE utf8_danish_ci
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `peer_reviews`
--

CREATE TABLE `peer_reviews`
(
  `id`             int(11) NOT NULL,
  `submission_id`  int(11) NOT NULL,
  `assignment_id`  int(11) NOT NULL,
  `user_id`        int(11) NOT NULL,
  `review_user_id` int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `reviews`
--

CREATE TABLE `reviews`
(
  `id`      int(11) NOT NULL,
  `form_id` int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `submissions`
--

CREATE TABLE `submissions`
(
  `id`      int(11) NOT NULL,
  `form_id` int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `submissions`
--

INSERT INTO `submissions` (`id`, `form_id`)
VALUES (1, 1);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `usercourse`
--

CREATE TABLE `usercourse`
(
  `userid`   int(11) NOT NULL,
  `courseid` int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `usercourse`
--

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

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `users`
--

CREATE TABLE `users`
(
  `id`            int(11)                            NOT NULL,
  `name`          varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `email_student` varchar(64) COLLATE utf8_danish_ci NOT NULL,
  `teacher`       tinyint(1)                         NOT NULL,
  `email_private` varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `password`      varchar(64) COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `users`
--

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

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `user_reviews`
--

CREATE TABLE `user_reviews`
(
  `id`        int(11)  NOT NULL,
  `user_id`   int(11)  NOT NULL,
  `review_id` int(11)  NOT NULL,
  `data`      longtext NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `user_submissions`
--

CREATE TABLE `user_submissions`
(
  `id`            int(11)     NOT NULL,
  `user_id`       int(11)     NOT NULL,
  `assignment_id` int(11)     NOT NULL,
  `submission_id` int(11)     NOT NULL,
  `type`          varchar(64) NOT NULL,
  `answer`        mediumtext  NULL,
  `submitted`     timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Tabellstruktur for tabell `schedule_tasks`
--

CREATE TABLE `schedule_tasks`
(
  `id`             int(11)     NOT NULL,
  `submission_id`  int(11)     not null,
  `assignment_id`  int(11)     not null,
  `scheduled_time` datetime    not null,
  `task`           varchar(32) not null,
  `data`           blob        not null
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `user_submissions`
--

INSERT INTO `user_submissions` (`id`, `user_id`, `assignment_id`, `submission_id`, `type`, `answer`)
VALUES (1, 4, 1, 1, 'text', 'JohanKlausen'),
       (2, 4, 1, 1, 'url', 'https://github.com/JohanKlausen/yeet'),
       (3, 4, 1, 1, 'textarea', 'I did good!'),
       (4, 5, 1, 1, 'text', 'StianFjerdingstad'),
       (5, 5, 1, 1, 'url', 'https://github.com/StianFjerdingstad/Sudoku'),
       (6, 5, 1, 1, 'textarea', 'I did sexy good!'),
       (7, 10, 1, 1, 'text', 'KlausAanesen'),
       (8, 10, 1, 1, 'url', 'https://github.com/KlausAanesen/1337yeet420'),
       (9, 10, 1, 1, 'textarea', 'I did bad :(');

INSERT INTO `peer_reviews` (`id`, `submission_id`, `assignment_id`, `user_id`, `review_user_id`)
VALUES (1, 1, 1, 3, 4),
       (2, 1, 1, 3, 5),
       (3, 1, 2, 9, 4);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `adminfaq`
--
ALTER TABLE `adminfaq`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `assignments`
--
ALTER TABLE `assignments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `assignments_courses_id_fk` (`course_id`),
  ADD KEY `assignments_reviews_id_fk` (`review_id`),
  ADD KEY `assignments_submissions_id_fk` (`submission_id`);

--
-- Indexes for table `course`
--
ALTER TABLE `course`
  ADD PRIMARY KEY (`id`),
  ADD KEY `teacher` (`teacher`);

--
-- Indexes for table `fields`
--
ALTER TABLE `fields`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fields_forms_id_fk` (`form_id`);

--
-- Indexes for table `forms`
--
ALTER TABLE `forms`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `logs`
--
ALTER TABLE `logs`
  ADD KEY `userid` (`userid`),
  ADD KEY `courseid` (`courseid`),
  ADD KEY `assignmentid` (`assignmentid`),
  ADD KEY `submissionid` (`submissionid`);


--
-- Indexes for table `reviews`
--
ALTER TABLE `reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `reviews_forms_id_fk` (`form_id`);

--
-- Indexes for table `submissions`
--
ALTER TABLE `submissions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `submissions_forms_id_fk` (`form_id`);

--
-- Indexes for table `usercourse`
--
ALTER TABLE `usercourse`
  ADD KEY `userid` (`userid`),
  ADD KEY `courseid` (`courseid`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email_student` (`email_student`),
  ADD UNIQUE KEY `email_private` (`email_private`);

--
-- Indexes for table `user_reviews`
--
ALTER TABLE `user_reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_reviews_reviews_id_fk` (`review_id`),
  ADD KEY `user_reviews_users_id_fk` (`user_id`);

--
-- Indexes for table `user_submissions`
--
ALTER TABLE `user_submissions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_submissions_submissions_id_fk` (`submission_id`),
  ADD KEY `user_submissions_users_id_fk` (`user_id`),
  ADD KEY `user_submission_assignment_id_fk` (`assignment_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `adminfaq`
--
ALTER TABLE `adminfaq`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 2;
--
-- AUTO_INCREMENT for table `assignments`
--
ALTER TABLE `assignments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 2;
--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 4;
--
-- AUTO_INCREMENT for table `fields`
--
ALTER TABLE `fields`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 4;
--
-- AUTO_INCREMENT for table `forms`
--
ALTER TABLE `forms`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 2;

--
-- AUTO_INCREMENT for table `reviews`
--
ALTER TABLE `reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `submissions`
--
ALTER TABLE `submissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 2;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 11;
--
-- AUTO_INCREMENT for table `user_reviews`
--
ALTER TABLE `user_reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `user_submissions`
--
ALTER TABLE `user_submissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 4;


ALTER TABLE `schedule_tasks`
  ADD PRIMARY KEY (`id`),
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- Begrensninger for dumpede tabeller
--

--
-- Indexes for table `peer_reviews`
--
ALTER TABLE `peer_reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `peer_reviews_submissions_id_fk` (`submission_id`),
  ADD KEY `peer_reviews_assignment_id_fk` (`assignment_id`),
  ADD KEY `peer_reviews_user_id_fk` (`user_id`),
  ADD KEY `peer_reviews_review_user_id_fk` (`review_user_id`);

--
-- AUTO_INCREMENT for table `peer_reviews`
--
ALTER TABLE `peer_reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;


--
-- Begrensninger for tabell `assignments`
--
ALTER TABLE `assignments`
  ADD CONSTRAINT `assignments_courses_id_fk` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`),
  ADD CONSTRAINT `assignments_reviews_id_fk` FOREIGN KEY (`review_id`) REFERENCES `reviews` (`id`),
  ADD CONSTRAINT `assignments_submissions_id_fk` FOREIGN KEY (`submission_id`) REFERENCES `submissions` (`id`);

--
-- Begrensninger for tabell `course`
--
ALTER TABLE `course`
  ADD CONSTRAINT `course_ibfk_1` FOREIGN KEY (`teacher`) REFERENCES `users` (`id`)
    ON UPDATE CASCADE;

--
-- Begrensninger for tabell `fields`
--
ALTER TABLE `fields`
  ADD CONSTRAINT `fields_forms_id_fk` FOREIGN KEY (`form_id`) REFERENCES `forms` (`id`);


--
-- Begrensninger for tabell `logs`
--
ALTER TABLE `logs`
  ADD CONSTRAINT `logs_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE CASCADE,
  ADD CONSTRAINT `logs_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`)
    ON DELETE NO ACTION
    ON UPDATE CASCADE;


--
-- Begrensninger for tabell `reviews`
--
ALTER TABLE `reviews`
  ADD CONSTRAINT `reviews_forms_id_fk` FOREIGN KEY (`form_id`) REFERENCES `forms` (`id`);


--
-- Begrensninger for tabell `submissions`
--
ALTER TABLE `submissions`
  ADD CONSTRAINT `submissions_forms_id_fk` FOREIGN KEY (`form_id`) REFERENCES `forms` (`id`);


--
-- Begrensninger for tabell `usercourse`
--
ALTER TABLE `usercourse`
  ADD CONSTRAINT `usercourse_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  ADD CONSTRAINT `usercourse_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE;


--
-- Begrensninger for tabell `user_reviews`
--
ALTER TABLE `user_reviews`
  ADD CONSTRAINT `user_reviews_reviews_id_fk` FOREIGN KEY (`review_id`) REFERENCES `reviews` (`id`),
  ADD CONSTRAINT `user_reviews_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Begrensninger for tabell `user_submissions`
--
ALTER TABLE `user_submissions`
  ADD CONSTRAINT `user_submission_assignment_id_fk` FOREIGN KEY (`assignment_id`) REFERENCES `assignments` (`id`),
  ADD CONSTRAINT `user_submissions_submissions_id_fk` FOREIGN KEY (`submission_id`) REFERENCES `submissions` (`id`),
  ADD CONSTRAINT `user_submissions_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);


--
-- Begrensninger for tabell `schedule_tasks`
--
ALTER TABLE `schedule_tasks`
  ADD CONSTRAINT schedule_tasks_submission_id_fk FOREIGN KEY (submission_id) REFERENCES submissions (id),
  ADD CONSTRAINT schedule_tasks_assignment_id_fk FOREIGN KEY (assignment_id) REFERENCES assignments (id);

--
-- Begrensninger for tabell `peer_reviews`
--
ALTER TABLE `peer_reviews`
  ADD CONSTRAINT `peer_reviews_review_user_id_fk` FOREIGN KEY (`review_user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `peer_reviews_submissions_id_fk` FOREIGN KEY (`submission_id`) REFERENCES `submissions` (`id`),
  ADD CONSTRAINT `peer_reviews_assignment_id_fk` FOREIGN KEY (`assignment_id`) REFERENCES `assignments` (`id`),
  ADD CONSTRAINT `peer_reviews_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);


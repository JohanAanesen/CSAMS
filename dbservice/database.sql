-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 29. Jan, 2019 17:17 PM
-- Server-versjon: 10.1.28-MariaDB
-- PHP Version: 7.1.10

CREATE DATABASE IF NOT EXISTS cs53 COLLATE = utf8_general_ci;

USE cs53;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+01:00"; -- Norwegian time zone! --


/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cs53`
--

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `course`
--

CREATE TABLE `course`
(
  `id`          int(11)                                        NOT NULL,
  `hash`        varchar(64)                                    NOT NULL,
  `coursecode`  varchar(10) COLLATE utf8_danish_ci             NOT NULL,
  `coursename`  varchar(64) COLLATE utf8_danish_ci             NOT NULL,
  `teacher`     int(11)                                        NOT NULL,
  `description` text COLLATE utf8_danish_ci,
  `year`        int(11) COLLATE utf8_danish_ci                 NOT NULL,
  `semester`    ENUM ('fall', 'spring') COLLATE utf8_danish_ci NOT NULL
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

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
  `oldvalue`     text COLLATE utf8_danish_ci                 DEFAULT NULL,
  `newValue`     text COLLATE utf8_danish_ci                 DEFAULT NULL
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;
-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `usercourse`
--

CREATE TABLE `usercourse`
(
  `userid`   int(11) NOT NULL,
  `courseid` int(11) NOT NULL
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

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
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;


CREATE TABLE `adminfaq`
(
  `id`        int(11)  NOT NULL,
  `timestamp` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `questions` text     NOT NULL
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

--
-- Indexes for dumped tables
--
--
-- Indexes for dumped tables
--


-- Insert Test info --
INSERT INTO `users` (`id`, `name`, `email_student`, `teacher`, `email_private`, `password`)
VALUES (1,
        'Test User',
        'hei@gmail.com',
        1,
        'test@yahoo.com',
        '$2a$14$MZj24p41j2NNGn6JDsQi0OsDb56.0LcfrIdgjE6WmZzp58O6V/VhK'),
       (2,
        'Frode Haug',
        'frodehg@teach.ntnu.no',
        1,
        NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'), -- Password is 123abc --
       (3,
        'Ola Nordmann',
        'olanor@stud.ntnu.no',
        1,
        'swag-meister69@ggmail.com',
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'), -- Password is 123abc --
       (4,
        'Johan Klausen',
        'johkl@stu.ntnu.no',
        0,
        NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'); -- Password is 123abc --


INSERT INTO `course` (`id`, `hash`, `coursecode`, `coursename`, `teacher`, `year`, `semester`, `description`)
VALUES (1, '3876438629b786', 'IMT1031', 'Grunnleggende Programmering', 2, 2019, 'fall', 'Write hello, world in C++'),
       (2,
        '12387teg817eg18',
        'IMT1082',
        'Objekt-orientert programmering',
        2,
        2019,
        'fall',
        'Write Wazz up world in Python'),
       (3, '12e612eg1e17ge1', 'IMT2021', 'Algoritmiske metoder', 2, 2019, 'spring', 'Write an AI in C#');

INSERT INTO `usercourse` (`userid`, `courseid`)
VALUES (3, 1),
       (3, 2),
       (4, 2);

INSERT INTO `adminfaq` (`id`, `timestamp`, `questions`)
VALUES ('1',
        '1997-02-13 13:37:00',
        'Q: How do I make a course + link?\r\n--------------------------------\r\n**A:** Dashboard -> Courses -> new. And create the course there\r\n\r\nQ: How do I make an assignment?\r\n--------------------------------\r\n**A:** Dashboard -> Assignments-> new. And create the assignment there\r\n\r\nQ: How do I invite students to the course?\r\n--------------------------------\r\n**A:** Create a link for the course and email the students the link\r\n\r\nQ: How do I import database?\r\n--------------------------------\r\n**A:** Start xampp and go to import in phpmyadmin\r\n\r\nQ: How do I export database?\r\n--------------------------------\r\n**A:** Start xampp and go to export in phpmyadmin\r\n\r\nQ: How do I sign up?\r\n--------------------------------\r\n**A:** You go to `/register` and register a user there\n\n![Reddit](https://external-preview.redd.it/lzcL5WbUuBr7pI9zIM9ZbUSrETZR1UNb-g6C5DehYss.jpg?width=960&crop=smart&auto=webp&s=4b483a024ac9103bfe6df2e98599043bbed29146)');
-- end --

--
-- Indexes for table `course`
--
ALTER TABLE `course`
  ADD PRIMARY KEY (`id`),
  ADD KEY `teacher` (`teacher`);

--
-- Indexes for table `logs`
--
ALTER TABLE `logs`
  ADD KEY `userid` (`userid`),
  ADD KEY `courseid` (`courseid`),
  ADD KEY `assignmentid` (`assignmentid`),
  ADD KEY `submissionid` (`submissionid`);

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

ALTER TABLE `adminfaq`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;


ALTER TABLE `adminfaq`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Begrensninger for dumpede tabeller
--

--
-- Begrensninger for tabell `course`
--
ALTER TABLE `course`
  ADD CONSTRAINT `course_ibfk_1` FOREIGN KEY (`teacher`) REFERENCES `users` (`id`)
  ON UPDATE CASCADE;

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
-- Begrensninger for tabell `usercourse`
--
ALTER TABLE `usercourse`
  ADD CONSTRAINT `usercourse_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`)
  ON DELETE CASCADE
  ON UPDATE CASCADE,
  ADD CONSTRAINT `usercourse_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`)
  ON DELETE CASCADE
  ON UPDATE CASCADE;
COMMIT;


create table forms
(
  id          int auto_increment
    primary key,
  prefix      varchar(64)                         not null,
  name        varchar(64)                         null,
  description text                                null,
  created     timestamp default CURRENT_TIMESTAMP null
);

create table fields
(
  id          int auto_increment
    primary key,
  form_id     int         not null,
  type        varchar(64) not null,
  name        varchar(64) not null,
  label       varchar(64) null,
  description text        not null,
  priority     int         not null,
  weight      int         null,
  choices     varchar(64) null,
  constraint fields_forms_id_fk
  foreign key (form_id) references forms (id)
);

create table reviews
(
  id      int auto_increment
    primary key,
  form_id int not null,
  constraint reviews_forms_id_fk
  foreign key (form_id) references forms (id)
);

create table submissions
(
  id      int auto_increment
    primary key,
  form_id int not null,
  constraint submissions_forms_id_fk
  foreign key (form_id) references forms (id)
);

create table assignments
(
  id            int auto_increment
    primary key,
  name          varchar(64)                         not null,
  description   text                                null,
  created       timestamp default CURRENT_TIMESTAMP not null,
  publish       datetime                            not null,
  deadline      datetime                            not null,
  course_id     int                                 not null,
  submission_id int                                 null,
  review_id     int                                 null,
  constraint assignments_courses_id_fk
  foreign key (course_id) references course (id),
  constraint assignments_reviews_id_fk
  foreign key (review_id) references reviews (id),
  constraint assignments_submissions_id_fk
  foreign key (submission_id) references submissions (id)
);

create table user_reviews
(
  id        int auto_increment
    primary key,
  user_id   int      not null,
  review_id int      not null,
  data      longtext not null,
  constraint user_reviews_reviews_id_fk
  foreign key (review_id) references reviews (id),
  constraint user_reviews_users_id_fk
  foreign key (user_id) references users (id)
);

create table user_submissions
(
  id            int auto_increment
    primary key,
  user_id       int        not null,
  assignment_id int not null,
  submission_id int        not null,
  type    varchar(64)   not null,
  answer        mediumtext null,
  constraint user_submissions_submissions_id_fk
  foreign key (submission_id) references submissions (id),
  constraint user_submissions_users_id_fk
  foreign key (user_id) references users (id),
  constraint user_submission_assignment_id_fk
  foreign key (assignment_id) references assignments (id)
);

create table peer_reviews
(
  id                   int auto_increment primary key,
  submission_id        int not null,
  user_id              int not null,
  review_submission_id int not null,
  constraint peer_reviews_submissions_id_fk
  foreign key (submission_id) references submissions (id),
  constraint peer_reviews_user_id_fk
  foreign key (user_id) references users (id),
  constraint peer_reviews_review_submission_id_fk
  foreign key (review_submission_id) references user_submissions (id)
);

/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS = @OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION = @OLD_COLLATION_CONNECTION */;
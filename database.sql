-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 29. Jan, 2019 17:17 PM
-- Server-versjon: 10.1.28-MariaDB
-- PHP Version: 7.1.10

DROP SCHEMA
  IF EXISTS cs53;
CREATE SCHEMA cs53 COLLATE = utf8_general_ci;

USE cs53;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cs53`
--

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `assignments`
--

CREATE TABLE `assignments`
(
  `id`       int(11)    NOT NULL,
  `courseid` int(11)    NOT NULL,
  `created`  timestamp  NOT NULL                DEFAULT CURRENT_TIMESTAMP,
  `due`      date       NOT NULL,
  `peer`     tinyint(1) NOT NULL                DEFAULT '0',
  `auto`     tinyint(1) NOT NULL                DEFAULT '0',
  `language` varchar(20) COLLATE utf8_danish_ci DEFAULT NULL,
  `tasktext` text COLLATE utf8_danish_ci,
  `payload`  int(11)                            DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `course`
--

CREATE TABLE `course`
(
  `id`          int(11)                                        NOT NULL,
  `coursecode`  varchar(10) COLLATE utf8_danish_ci             NOT NULL,
  `coursename`  varchar(64) COLLATE utf8_danish_ci             NOT NULL,
  `teacher`     int(11)                                        NOT NULL,
  `description` text COLLATE utf8_danish_ci,
  `year`        int(11) COLLATE utf8_danish_ci                 NOT NULL,
  `semester`    ENUM ('fall', 'spring') COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `logs`
--

CREATE TABLE `logs`
(
  `userid` int(11)                             NOT NULL,
  `log`    varchar(256) COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `peerreviews`
--

CREATE TABLE `peerreviews`
(
  `assignmentid` int(11)                             NOT NULL,
  `submissionid` int(11)                             NOT NULL,
  `userid`       int(11)                             NOT NULL,
  `grade`        int(11)                             NOT NULL,
  `feedback`     varchar(512) COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `reviewerpairs`
--

CREATE TABLE `reviewerpairs`
(
  `assignmentid`   int(11) NOT NULL,
  `userid`         int(11) NOT NULL,
  `submissionsid1` int(11) NOT NULL,
  `submissionid2`  int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `submissions`
--

CREATE TABLE `submissions`
(
  `id`           int(11)                             NOT NULL,
  `userid`       int(11)                             NOT NULL,
  `assignmentid` int(11)                             NOT NULL,
  `repo`         varchar(128) COLLATE utf8_danish_ci NOT NULL,
  `deploy`       varchar(128) COLLATE utf8_danish_ci NOT NULL,
  `comment`      varchar(512) COLLATE utf8_danish_ci NOT NULL,
  `grade`        int(11) DEFAULT NULL,
  `test`         int(11) DEFAULT NULL,
  `vet`          int(11) DEFAULT NULL,
  `cycle`        int(11) DEFAULT NULL
) ENGINE = InnoDB
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
) ENGINE = InnoDB
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
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_danish_ci;

--
-- Indexes for dumped tables
--


-- Insert Test info --
INSERT INTO `users` (`id`, `name`, `email_student`, `teacher`, `email_private`, `password`)
VALUES (1, 'Test User', 'hei@gmail.com', 1, NULL, '$2a$14$MZj24p41j2NNGn6JDsQi0OsDb56.0LcfrIdgjE6WmZzp58O6V/VhK'),
       (2, 'Frode Haug', 'frodehg@teach.ntnu.no', 1, NULL,
        '$2a$14$vH/ibjwwXqBmGgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (3, 'Ola Nordmann', 'olanor@stud.ntnu.no', 0, 'swag-meister69@ggmail.com',
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (4, 'Johan Klausen', 'johkl@stu.ntnu.no', 0, NULL,
        '$2a$14$MZj24p41j2NNGn6JDsQi0OsDb56.0LcfrIdgjE6WmZzp58O6V/VhK');


INSERT INTO `course` (`id`, `coursecode`, `coursename`, `teacher`, `year`, `semester`, `description`)
VALUES (1, 'IMT1031', 'Grunnleggende Programmering', 2, 2019, 'fall', 'Write hello, world in C++'),
       (2, 'IMT1082', 'Objekt-orientert programmering', 2, 2019, 'fall', 'Write Wazz up world in Python'),
       (3, 'IMT2021', 'Algoritmiske metoder', 2, 2019, 'spring', 'Write an AI in C#');

INSERT INTO `usercourse` (`userid`, `courseid`)
VALUES (3, 1),
       (3, 2),
       (4, 3);

-- end --

--
-- Indexes for table `assignments`
--
ALTER TABLE `assignments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `courseid` (`courseid`);

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
  ADD KEY `userid` (`userid`);

--
-- Indexes for table `peerreviews`
--
ALTER TABLE `peerreviews`
  ADD KEY `assignmentid` (`assignmentid`),
  ADD KEY `submissionid` (`submissionid`),
  ADD KEY `userid` (`userid`);

--
-- Indexes for table `reviewerpairs`
--
ALTER TABLE `reviewerpairs`
  ADD KEY `assignmentid` (`assignmentid`),
  ADD KEY `userid` (`userid`),
  ADD KEY `submissionsid1` (`submissionsid1`),
  ADD KEY `submissionid2` (`submissionid2`);

--
-- Indexes for table `submissions`
--
ALTER TABLE `submissions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `userid` (`userid`),
  ADD KEY `assignmentid` (`assignmentid`);

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
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `assignments`
--
ALTER TABLE `assignments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `submissions`
--
ALTER TABLE `submissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Begrensninger for dumpede tabeller
--

--
-- Begrensninger for tabell `assignments`
--
ALTER TABLE `assignments`
  ADD CONSTRAINT `assignments_ibfk_1` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `course`
--
ALTER TABLE `course`
  ADD CONSTRAINT `course_ibfk_1` FOREIGN KEY (`teacher`) REFERENCES `users` (`id`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `logs`
--
ALTER TABLE `logs`
  ADD CONSTRAINT `logs_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE;

--
-- Begrensninger for tabell `peerreviews`
--
ALTER TABLE `peerreviews`
  ADD CONSTRAINT `peerreviews_ibfk_1` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `peerreviews_ibfk_2` FOREIGN KEY (`submissionid`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `peerreviews_ibfk_3` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE;

--
-- Begrensninger for tabell `reviewerpairs`
--
ALTER TABLE `reviewerpairs`
  ADD CONSTRAINT `reviewerpairs_ibfk_1` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_3` FOREIGN KEY (`submissionsid1`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_4` FOREIGN KEY (`submissionid2`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `submissions`
--
ALTER TABLE `submissions`
  ADD CONSTRAINT `submissions_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `submissions_ibfk_2` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `usercourse`
--
ALTER TABLE `usercourse`
  ADD CONSTRAINT `usercourse_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `usercourse_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS = @OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION = @OLD_COLLATION_CONNECTION */;


-- phpMyAdmin SQL Dump
-- version 4.7.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 20, 2019 at 01:19 PM
-- Server version: 10.1.26-MariaDB
-- PHP Version: 7.1.8

DROP SCHEMA
  IF EXISTS cs53;
CREATE SCHEMA cs53 COLLATE = utf8_general_ci;

USE cs53;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cs53`
--

--
-- Table structure for table `adminfaq`
--

CREATE TABLE `adminfaq` (
  `id` int(11) NOT NULL,
  `timestamp` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `questions` text COLLATE utf8_danish_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

--
-- Dumping data for table `adminfaq`
--

INSERT INTO `adminfaq` (`id`, `timestamp`, `questions`) VALUES
(1, '1997-02-13 13:37:00', 'Q: How do I make a course + link?\r\n--------------------------------\r\n**A:** Dashboard -> Courses -> new. And create the course there\r\n\r\nQ: How do I make an assignment?\r\n--------------------------------\r\n**A:** Dashboard -> Assignments-> new. And create the assignment there\r\n\r\nQ: How do I invite students to the course?\r\n--------------------------------\r\n**A:** Create a link for the course and email the students the link\r\n\r\nQ: How do I import database?\r\n--------------------------------\r\n**A:** Start xampp and go to import in phpmyadmin\r\n\r\nQ: How do I export database?\r\n--------------------------------\r\n**A:** Start xampp and go to export in phpmyadmin\r\n\r\nQ: How do I sign up?\r\n--------------------------------\r\n**A:** You go to `/register` and register a user there\n\n![Reddit](https://external-preview.redd.it/lzcL5WbUuBr7pI9zIM9ZbUSrETZR1UNb-g6C5DehYss.jpg?width=960&crop=smart&auto=webp&s=4b483a024ac9103bfe6df2e98599043bbed29146)');

-- --------------------------------------------------------

-- --------------------------------------------------------

--
-- Table structure for table `assignments`
--

CREATE TABLE `assignments` (
  `id` int(11) NOT NULL,
  `name` tinytext,
  `description` mediumtext,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `publish` datetime DEFAULT NULL,
  `deadline` datetime DEFAULT NULL,
  `course_id` int(11) DEFAULT NULL,
  `submission_id` int(11) DEFAULT NULL,
  `review_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `course`
--

CREATE TABLE `course` (
  `id` int(11) NOT NULL,
  `coursecode` varchar(10) COLLATE utf8_danish_ci NOT NULL,
  `coursename` varchar(64) COLLATE utf8_danish_ci NOT NULL,
  `teacher` int(11) NOT NULL,
  `description` text COLLATE utf8_danish_ci,
  `year` int(11) NOT NULL,
  `semester` enum('fall','spring') COLLATE utf8_danish_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

--
-- Dumping data for table `course`
--

INSERT INTO `course` (`id`, `coursecode`, `coursename`, `teacher`, `description`, `year`, `semester`) VALUES
(1, 'IMT1031', 'Grunnleggende Programmering', 2, 'Write hello, world in C++', 2019, 'fall'),
(2, 'IMT1082', 'Objekt-orientert programmering', 2, 'Write Wazz up world in Python', 2019, 'fall'),
(3, 'IMT2021', 'Algoritmiske metoder', 2, 'Write an AI in C#', 2019, 'spring');

-- --------------------------------------------------------

--
-- Table structure for table `fields`
--

CREATE TABLE `fields` (
  `id` int(11) NOT NULL,
  `form_id` int(11) NOT NULL,
  `type` tinytext NOT NULL,
  `name` tinytext NOT NULL,
  `label` tinytext NOT NULL,
  `description` mediumtext,
  `priority` int(11) NOT NULL,
  `weight` int(11) DEFAULT NULL,
  `choices` tinytext
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `forms`
--

CREATE TABLE `forms` (
  `id` int(11) NOT NULL,
  `prefix` tinytext NOT NULL,
  `name` tinytext NOT NULL,
  `description` mediumtext,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `forms`
--

INSERT INTO `forms` (`id`, `prefix`, `name`, `description`, `created`) VALUES
(16, 'rest_api', 'REST API', 'Form Description', '2019-02-20 10:06:58'),
(17, 'gitrepo,_comment', 'GitRepo, Comment', 'Testing testing', '2019-02-20 12:13:07');

-- --------------------------------------------------------

--
-- Table structure for table `logs`
--

CREATE TABLE `logs` (
  `userid` int(11) NOT NULL,
  `timestamp` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `activity` varchar(32) COLLATE utf8_danish_ci NOT NULL,
  `assignmentid` int(11) DEFAULT NULL,
  `courseid` int(11) DEFAULT NULL,
  `submissionid` int(11) DEFAULT NULL,
  `oldvalue` varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `newValue` varchar(64) COLLATE utf8_danish_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `peerreviews`
--

CREATE TABLE `peerreviews` (
  `assignmentid` int(11) NOT NULL,
  `submissionid` int(11) NOT NULL,
  `userid` int(11) NOT NULL,
  `grade` int(11) NOT NULL,
  `feedback` varchar(512) COLLATE utf8_danish_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `reviewerpairs`
--

CREATE TABLE `reviewerpairs` (
  `assignmentid` int(11) NOT NULL,
  `userid` int(11) NOT NULL,
  `submissionsid1` int(11) NOT NULL,
  `submissionid2` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `submissions`
--

CREATE TABLE `submissions` (
  `id` int(11) NOT NULL,
  `form_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `usercourse`
--

CREATE TABLE `usercourse` (
  `userid` int(11) NOT NULL,
  `courseid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

--
-- Dumping data for table `usercourse`
--

INSERT INTO `usercourse` (`userid`, `courseid`) VALUES
(3, 1),
(3, 2),
(4, 3);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `email_student` varchar(64) COLLATE utf8_danish_ci NOT NULL,
  `teacher` tinyint(1) NOT NULL,
  `email_private` varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `password` varchar(64) COLLATE utf8_danish_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email_student`, `teacher`, `email_private`, `password`) VALUES
(2, 'Frode Haug', 'frodehg@teach.ntnu.no', 1, NULL, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
(3, 'Ola Nordmann', 'olanor@stud.ntnu.no', 0, 'swag-meister69@ggmail.com', '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
(4, 'Johan Klausen', 'johkl@stu.ntnu.no', 0, NULL, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');

-- --------------------------------------------------------

--
-- Table structure for table `user_submissions`
--

CREATE TABLE `user_submissions` (
  `id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `submission_id` int(11) DEFAULT NULL,
  `field_name` tinytext NOT NULL,
  `answer` varchar(4369) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  ADD KEY `assignments_course_id_fk` (`course_id`);

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
-- Indexes for table `user_submissions`
--
ALTER TABLE `user_submissions`
  ADD KEY `user_submissions_submissions_id_fk` (`submission_id`),
  ADD KEY `user_submissions_users_id_fk` (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `adminfaq`
--
ALTER TABLE `adminfaq`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `assignments`
--
ALTER TABLE `assignments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT for table `fields`
--
ALTER TABLE `fields`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT for table `forms`
--
ALTER TABLE `forms`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;
--
-- AUTO_INCREMENT for table `submissions`
--
ALTER TABLE `submissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `assignments`
--
ALTER TABLE `assignments`
  ADD CONSTRAINT `assignments_course_id_fk` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`);

--
-- Constraints for table `course`
--
ALTER TABLE `course`
  ADD CONSTRAINT `course_ibfk_1` FOREIGN KEY (`teacher`) REFERENCES `users` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `fields`
--
ALTER TABLE `fields`
  ADD CONSTRAINT `fields_forms_id_fk` FOREIGN KEY (`form_id`) REFERENCES `forms` (`id`);

--
-- Constraints for table `logs`
--
ALTER TABLE `logs`
  ADD CONSTRAINT `logs_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  ADD CONSTRAINT `logs_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  ADD CONSTRAINT `logs_ibfk_3` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  ADD CONSTRAINT `logs_ibfk_4` FOREIGN KEY (`submissionid`) REFERENCES `submissions` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE;

--
-- Constraints for table `peerreviews`
--
ALTER TABLE `peerreviews`
  ADD CONSTRAINT `peerreviews_ibfk_1` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `peerreviews_ibfk_2` FOREIGN KEY (`submissionid`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `peerreviews_ibfk_3` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE;

--
-- Constraints for table `reviewerpairs`
--
ALTER TABLE `reviewerpairs`
  ADD CONSTRAINT `reviewerpairs_ibfk_1` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_3` FOREIGN KEY (`submissionsid1`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_4` FOREIGN KEY (`submissionid2`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `submissions`
--
ALTER TABLE `submissions`
  ADD CONSTRAINT `submissions_forms_id_fk` FOREIGN KEY (`form_id`) REFERENCES `forms` (`id`);

--
-- Constraints for table `usercourse`
--
ALTER TABLE `usercourse`
  ADD CONSTRAINT `usercourse_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `usercourse_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `user_submissions`
--
ALTER TABLE `user_submissions`
  ADD CONSTRAINT `user_submissions_submissions_id_fk` FOREIGN KEY (`submission_id`) REFERENCES `submissions` (`id`),
  ADD CONSTRAINT `user_submissions_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

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
SET time_zone = "+01:00";


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
  `hash`        varchar(64)                                    NOT NULL,
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
  `userid`       int(11)                            NOT NULL,
  `timestamp`    datetime                           NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `activity`     varchar(32) COLLATE utf8_danish_ci NOT NULL,
  `assignmentid` int(11)                                     DEFAULT NULL,
  `courseid`     int(11)                                     DEFAULT NULL,
  `submissionid` int(11)                                     DEFAULT NULL,
  `oldvalue`     text COLLATE utf8_danish_ci                 DEFAULT NULL,
  `newValue`     text COLLATE utf8_danish_ci                 DEFAULT NULL
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
VALUES (1, 'Test User', 'hei@gmail.com', 1, 'test@yahoo.com',
        '$2a$14$MZj24p41j2NNGn6JDsQi0OsDb56.0LcfrIdgjE6WmZzp58O6V/VhK'),
       (2, 'Frode Haug', 'frodehg@teach.ntnu.no', 1, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'), -- Password is 123abc --
       (3, 'Ola Nordmann', 'olanor@stud.ntnu.no', 0, 'swag-meister69@ggmail.com',
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'), -- Password is 123abc --
       (4, 'Johan Klausen', 'johkl@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'); -- Password is 123abc --


INSERT INTO `course` (`id`, `hash`, `coursecode`, `coursename`, `teacher`, `year`, `semester`, `description`)
VALUES (1, '3876438629b786', 'IMT1031', 'Grunnleggende Programmering', 2, 2019, 'fall', 'Write hello, world in C++'),
       (2, '12387teg817eg18', 'IMT1082', 'Objekt-orientert programmering', 2, 2019, 'fall',
        'Write Wazz up world in Python'),
       (3, '12e612eg1e17ge1', 'IMT2021', 'Algoritmiske metoder', 2, 2019, 'spring', 'Write an AI in C#'),
       (4, '1337-420-69', 'IMT1337', 'Quick Markdown Example', 3, 1814, 'spring',
        'An h1 header\r\n============\r\n\r\nParagraphs are separated by a blank line.\r\n\r\n2nd paragraph. *Italic*, **bold**, and `monospace`. Itemized lists\r\nlook like:\r\n\r\n  * this one\r\n  * that one\r\n  * the other one\r\n\r\nNote that --- not considering the asterisk --- the actual text\r\ncontent starts at 4-columns in.\r\n\r\n> Block quotes are\r\n> written like so.\r\n>\r\n> They can span multiple paragraphs,\r\n> if you like.\r\n\r\nUse 3 dashes for an em-dash. Use 2 dashes for ranges (ex., \"it\'s all\r\nin chapters 12--14\"). Three dots ... will be converted to an ellipsis.\r\nUnicode is supported. ☺\r\n\r\n\r\n\r\nAn h2 header\r\n------------\r\n\r\nHere\'s a numbered list:\r\n\r\n 1. first item\r\n 2. second item\r\n 3. third item\r\n\r\nNote again how the actual text starts at 4 columns in (4 characters\r\nfrom the left side). Here\'s a code sample:\r\n\r\n    # Let me re-iterate ...\r\n    for i in 1 .. 10 { do-something(i) }\r\n\r\nAs you probably guessed, indented 4 spaces. By the way, instead of\r\nindenting the block, you can use delimited blocks, if you like:\r\n\r\n~~~\r\ndefine foobar() {\r\n    print \"Welcome to flavor country!\";\r\n}\r\n~~~\r\n\r\n(which makes copying & pasting easier). You can optionally mark the\r\ndelimited block for Pandoc to syntax highlight it:\r\n\r\n~~~python\r\nimport time\r\n# Quick, count to ten!\r\nfor i in range(10):\r\n    # (but not *too* quick)\r\n    time.sleep(0.5)\r\n    print(i)\r\n~~~\r\n\r\n\r\n\r\n### An h3 header ###\r\n\r\nNow a nested list:\r\n\r\n 1. First, get these ingredients:\r\n\r\n      * carrots\r\n      * celery\r\n      * lentils\r\n\r\n 2. Boil some water.\r\n\r\n 3. Dump everything in the pot and follow\r\n    this algorithm:\r\n\r\n        find wooden spoon\r\n        uncover pot\r\n        stir\r\n        cover pot\r\n        balance wooden spoon precariously on pot handle\r\n        wait 10 minutes\r\n        goto first step (or shut off burner when done)\r\n\r\n    Do not bump wooden spoon or it will fall.\r\n\r\nNotice again how text always lines up on 4-space indents (including\r\nthat last line which continues item 3 above).\r\n\r\nHere\'s a link to [a website](http://foo.bar), to a [local\r\ndoc](local-doc.html), and to a [section heading in the current\r\ndoc](#an-h2-header). Here\'s a footnote [^1].\r\n\r\n[^1]: Some footnote text.\r\n\r\nTables can look like this:\r\n\r\nName           Size  Material      Color\r\n------------- -----  ------------  ------------\r\nAll Business      9  leather       brown\r\nRoundabout       10  hemp canvas   natural\r\nCinderella       11  glass         transparent\r\n\r\nTable: Shoes sizes, materials, and colors.\r\n\r\n(The above is the caption for the table.) Pandoc also supports\r\nmulti-line tables:\r\n\r\n--------  -----------------------\r\nKeyword   Text\r\n--------  -----------------------\r\nred       Sunsets, apples, and\r\n          other red or reddish\r\n          things.\r\n\r\ngreen     Leaves, grass, frogs\r\n          and other things it\'s\r\n          not easy being.\r\n--------  -----------------------\r\n\r\nA horizontal rule follows.\r\n\r\n***\r\n\r\nHere\'s a definition list:\r\n\r\napples\r\n  : Good for making applesauce.\r\n\r\noranges\r\n  : Citrus!\r\n\r\ntomatoes\r\n  : There\'s no \"e\" in tomatoe.\r\n\r\nAgain, text is indented 4 spaces. (Put a blank line between each\r\nterm and  its definition to spread things out more.)\r\n\r\nHere\'s a \"line block\" (note how whitespace is honored):\r\n\r\n| Line one\r\n|   Line too\r\n| Line tree\r\n\r\nand images can be specified like so:\r\n\r\n![example image](https://external-preview.redd.it/6PB4LMzhKCFDUH15pwTJHT4b1Y63kq5Zjemvj0qbnrY.jpg?width=640&crop=smart&auto=webp&s=d2cfb9f54a8fc18d65b185a80b8473aba188be9b \"An exemplary image\")\r\n\r\nInline math equation: $\\omega = d\\phi / dt$. Display\r\nmath should get its own line like so:\r\n\r\n$$I = \\int \\rho R^{2} dV$$\r\n\r\nAnd note that you can backslash-escape any punctuation characters\r\nwhich you wish to be displayed literally, ex.: \\`foo\\`, \\*bar\\*, etc.');

INSERT INTO `usercourse` (`userid`, `courseid`)
VALUES (3, 1),
       (3, 2),
       (4, 2);

INSERT INTO `assignments` (`id`, `courseid`, `created`, `due`, `peer`, `auto`, `language`, `tasktext`, `payload`)
VALUES ('1', 1, CURRENT_TIMESTAMP, '2019-02-14', '1', '0', 'English',
        '# Assignment 1\r\n* Doctors and nurses\r\n<!-- Hello -->', '13');

INSERT INTO `submissions` (`id`, `userid`, `assignmentid`, `repo`, `deploy`, `comment`, `grade`, `test`, `vet`, `cycle`)
VALUES ('1', '3', '1', 'www.github.com/user3/submission1', 'Hello', 'I am grate progrman', '6', NULL, NULL, NULL);
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
  ADD CONSTRAINT `logs_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  ADD CONSTRAINT `logs_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  ADD CONSTRAINT `logs_ibfk_3` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  ADD CONSTRAINT `logs_ibfk_4` FOREIGN KEY (`submissionid`) REFERENCES `submissions` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE;

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


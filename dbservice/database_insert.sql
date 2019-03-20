-- Adminer 4.7.1 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

INSERT INTO `adminfaq` (`id`, `timestamp`, `questions`) VALUES
(1,	'2019-03-06 15:53:00',	'Q: How do I make a course?\n--------------------------------\n**A:** Dashboard -> Courses -> new. And create the course there\n\nQ: How do I invite students to course?\n--------------------------------\n**A:** You go to [admin/course](/admin/course) or [admin/](/admin) and on the course card, click the copy button to get the `join course through link` and send that to all students in preferred way (ex: email)\n\nQ: How do I make an assignment?\n--------------------------------\n**A:** Dashboard -> Assignments-> new. And create the assignment there');

INSERT INTO `assignments` (`id`, `name`, `description`, `created`, `publish`, `deadline`, `course_id`, `submission_id`, `review_id`, `validation_id`, `reviewers`) VALUES
(1,	'Test Assignment',	'# Test Assignment\r\n## This is assignment\r\n### Good assignment\r\n`5/7`',	'2019-03-18 11:52:10',	'2019-03-18 11:51:00',	'2019-03-23 18:00:00',	2,	NULL,	NULL,	NULL,	1),
(2,	'Test assignment',	'# Testing',	'2019-03-20 10:03:54',	'2019-03-25 12:00:00',	'2019-03-28 23:59:00',	1,	NULL,	NULL,	NULL,	1);

INSERT INTO `course` (`id`, `hash`, `coursecode`, `coursename`, `teacher`, `description`, `year`, `semester`) VALUES
(1,	'bi7n4as48c6b4j2l0fv0',	'IMT3673',	'Mobile Programming',	2,	'Mobile Programming Course as part of Bachelor in Programming',	2019,	'spring'),
(2,	'bi7ng8k48c6b4j2l0fvg',	'IMT1337',	'Test Course',	3,	'# IMT1337 - Test Course\r\n## This is course\r\n### Good course\r\n`10/10`',	2019,	'spring');

INSERT INTO `fields` (`id`, `form_id`, `type`, `name`, `description`, `label`, `hasComment`, `priority`, `weight`, `choices`) VALUES
(3,	1,	'url',	'testing_svein_gitrepo_w_comment_url_2',	'',	'Git Repository',	1,	0,	0,	''),
(38,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_29',	'',	'The app has a custom icon (not the default Android one).',	0,	0,	1,	''),
(39,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_30',	'',	'The app MainActivity loads.',	0,	1,	1,	''),
(40,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_31',	'',	'The app\'s MainActivity contains all required UI elements as per SPEC.',	0,	2,	1,	''),
(41,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_32',	'Pressing Android\'s \"back button\" on the MainActivity always quits the app, AND, this behaviour is not hardcoded in the code, ie. the code does not handle the Android back button presses, but instead, relies on the default Android behaviour.',	'\"Back button\"',	0,	3,	1,	''),
(42,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_33',	'',	'Pressing btn_transactions moves the user to TransactionsActivity.',	0,	4,	1,	''),
(43,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_34',	'',	'The default founding transaction from Angel is done correctly and visible in the TransactionsActivity. The user balance in lbl_balance matches the funding transaction.',	0,	5,	1,	''),
(44,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_35',	'',	'TransactionsActivity shows new payments correctly.',	0,	6,	1,	''),
(45,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_36',	'',	'TransactionsActivity moves the user back to MainActivity on \"back botton\" press.',	0,	7,	1,	''),
(46,	2,	'checkbox',	'lab_1_mobile_2019_checkbox_37',	'',	'Pressing btn_transfer moves the user to TransferActivity.',	0,	8,	1,	'');

INSERT INTO `forms` (`id`, `prefix`, `name`, `created`) VALUES
(1,	'testing_svein_gitrepo_w_comment',	'Testing (Svein) GitRepo w/Comment',	'2019-03-19 09:27:55'),
(2,	'lab_1_mobile_2019',	'Lab 1 - Mobile 2019',	'2019-03-20 10:14:13');

INSERT INTO `logs` (`userid`, `timestamp`, `activity`, `assignmentid`, `courseid`, `submissionid`, `oldvalue`, `newValue`) VALUES
(2,	'2019-03-18 11:25:47',	'COURSE-CREATED',	NULL,	1,	NULL,	NULL,	NULL),
(2,	'2019-03-18 11:25:47',	'JOINED-COURSE',	NULL,	1,	NULL,	NULL,	NULL),
(3,	'2019-03-18 11:51:14',	'COURSE-CREATED',	NULL,	2,	NULL,	NULL,	NULL),
(3,	'2019-03-18 11:51:14',	'JOINED-COURSE',	NULL,	2,	NULL,	NULL,	NULL);


INSERT INTO `reviews` (`id`, `form_id`) VALUES
(1,	2);


INSERT INTO `submissions` (`id`, `form_id`) VALUES
(1,	1);

INSERT INTO `usercourse` (`userid`, `courseid`) VALUES
(2,	1),
(3,	2),
(1,	1);

INSERT INTO `users` (`id`, `name`, `email_student`, `teacher`, `email_private`, `password`) VALUES
(1,	'Svein Are Danielsen',	'sveiad@stud.ntnu.no',	1,	NULL,	'$2a$14$ZiThiqRkYDj9wS5wyJjbl.jtpB8JLeN2Zztl6Kudhr.2e.bPdmk9W'),
(2,	'Christopher Frantz',	'christopher.frantz@ntnu.no',	1,	NULL,	'$2a$14$flEKp.4Q136bdgTopw.9wOg2JQ2Jp.rdMYyrDr4Fxwfa0X26DTHrq'),
(3,	'Brede Fritjof Klausen',	'bredefk@stud.ntnu.no',	1,	NULL,	'$2a$14$CUQPJwmIxm1oCM3w/pt5IOFjx5tIXWz8GXJTPJXDoH.0RXVcl6oj2');



-- 2019-03-20 10:42:23
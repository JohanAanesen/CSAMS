CREATE SCHEMA IF NOT EXISTS cs53
  COLLATE = utf8_general_ci;

USE cs53;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+01:00";
-- Norwegian time zone!  TODO time --


CREATE TABLE `adminfaq`
(
  `id`        int(11)  NOT NULL AUTO_INCREMENT,
  `timestamp` datetime NOT NULL,
  `questions` text     NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `users`
(
  `id`            int(11)     NOT NULL AUTO_INCREMENT,
  `name`          varchar(64) DEFAULT NULL,
  `email_student` varchar(64) NOT NULL,
  `teacher`       tinyint(1)  NOT NULL,
  `email_private` varchar(64) DEFAULT NULL,
  `password`      varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `course`
(
  `id`          int(11)                 NOT NULL AUTO_INCREMENT,
  `hash`        varchar(64)             NOT NULL,
  `coursecode`  varchar(10)             NOT NULL,
  `coursename`  varchar(64)             NOT NULL,
  `teacher`     int(11)                 NOT NULL,
  `description` text                    NOT NULL,
  `year`        int(11)                 NOT NULL,
  `semester`    enum ('fall', 'spring') NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`teacher`) REFERENCES users (`id`)
);

CREATE TABLE `usercourse`
(
  `userid`   int(11) NOT NULL,
  `courseid` int(11) NOT NULL,
  FOREIGN KEY (`userid`) REFERENCES users (`id`),
  FOREIGN KEY (`courseid`) REFERENCES course (`id`)
);

CREATE TABLE `forms`
(
  `id`      int(11)     NOT NULL AUTO_INCREMENT,
  `prefix`  varchar(256) NOT NULL,
  `name`    varchar(256) DEFAULT NULL,
  `created` timestamp   NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `fields`
(
  `id`          int(11)     NOT NULL AUTO_INCREMENT,
  `form_id`     int(11)     NOT NULL,
  `type`        varchar(64) NOT NULL,
  `name`        varchar(256) NOT NULL,
  `description` text        NOT NULL,
  `label`       varchar(256) DEFAULT NULL,
  `hasComment`  int(1)      NOT NULL,
  `priority`    int(11)     NOT NULL,
  `weight`      int(11)     DEFAULT NULL,
  `choices`     varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`form_id`) REFERENCES forms (`id`)
);

CREATE TABLE `reviews`
(
  `id`      int(11) NOT NULL AUTO_INCREMENT,
  `form_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`form_id`) REFERENCES forms (`id`)
);

CREATE TABLE `submissions`
(
  `id`      int(11) NOT NULL AUTO_INCREMENT,
  `form_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`form_id`) REFERENCES forms (`id`)
);

CREATE TABLE `assignments`
(
  `id`            int(11)     NOT NULL AUTO_INCREMENT,
  `name`          varchar(64) NOT NULL,
  `description`   text,
  `created`       timestamp   NOT NULL,
  `publish`       datetime    NOT NULL,
  `deadline`      datetime    NOT NULL,
  `course_id`     int(11)     NOT NULL,
  `submission_id` int(11) DEFAULT NULL,
  `review_id`     int(11) DEFAULT NULL,
  `validation_id` int(11) DEFAULT NULL,
  `reviewers`     int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`course_id`) REFERENCES course (`id`),
  FOREIGN KEY (`submission_id`) REFERENCES submissions (`id`),
  FOREIGN KEY (`review_id`) REFERENCES reviews (`id`)
);

CREATE TABLE `peer_reviews`
(
  `id`             int(11) NOT NULL AUTO_INCREMENT,
  `submission_id`  int(11) NOT NULL,
  `assignment_id`  int(11) NOT NULL,
  `user_id`        int(11) NOT NULL,
  `review_user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`assignment_id`) REFERENCES assignments (`id`),
  FOREIGN KEY (`submission_id`) REFERENCES submissions (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users (`id`),
  FOREIGN KEY (`review_user_id`) REFERENCES users (`id`)
);

CREATE TABLE `user_reviews`
(
  `id`            int(11)     NOT NULL AUTO_INCREMENT,
  `user_reviewer` int(11)     NOT NULL,
  `user_target`   int(11)     NOT NULL,
  `review_id`     int(11)     NOT NULL,
  `assignment_id` int(11)     NOT NULL,
  `type`          varchar(64) NOT NULL,
  `name`          varchar(256) NOT NULL,
  `label`         varchar(256) NOT NULL,
  `answer`        text        NOT NULL,
  `comment`       text                 DEFAULT NULL,
  `submitted`     datetime    NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`assignment_id`) REFERENCES assignments (`id`),
  FOREIGN KEY (`review_id`) REFERENCES reviews (`id`),
  FOREIGN KEY (`user_reviewer`) REFERENCES users (`id`),
  FOREIGN KEY (`user_target`) REFERENCES users (`id`)
);

CREATE TABLE `user_submissions`
(
  `id`            int(11)     NOT NULL AUTO_INCREMENT,
  `user_id`       int(11)     NOT NULL,
  `assignment_id` int(11)     NOT NULL,
  `submission_id` int(11)     NOT NULL,
  `type`          varchar(64) NOT NULL,
  `answer`        mediumtext  NULL,
  `comment`       text                 DEFAULT NULL,
  `submitted`     timestamp   NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`assignment_id`) REFERENCES assignments (`id`),
  FOREIGN KEY (`submission_id`) REFERENCES submissions (`id`)
);

CREATE TABLE `schedule_tasks`
(
  `id`             int(11)     NOT NULL AUTO_INCREMENT,
  `submission_id`  int(11)     NOT NULL,
  `assignment_id`  int(11)     NOT NULL,
  `scheduled_time` datetime    NOT NULL,
  `task`           varchar(32) NOT NULL,
  `data`           blob        NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`assignment_id`) REFERENCES assignments (`id`),
  FOREIGN KEY (`submission_id`) REFERENCES submissions (`id`)
);

-- not attaching any foreign keys to log because we always want it to log something even if some of the data is missing
CREATE TABLE `logs`
(
  `userid`       int(11)                            NOT NULL,
  `timestamp`    datetime                           NOT NULL,
  `activity`     varchar(32) COLLATE utf8_danish_ci NOT NULL,
  `assignmentid` int(11) DEFAULT NULL,
  `courseid`     int(11) DEFAULT NULL,
  `submissionid` int(11) DEFAULT NULL,
  `oldvalue`     text    DEFAULT NULL,
  `newValue`     text    DEFAULT NULL
);
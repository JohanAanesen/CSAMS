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
    `id`       int(11) NOT NULL AUTO_INCREMENT,
    `userid`   int(11) NOT NULL,
    `courseid` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`userid`) REFERENCES users (`id`),
    FOREIGN KEY (`courseid`) REFERENCES course (`id`)
);

CREATE TABLE `forms`
(
    `id`      int(11)      NOT NULL AUTO_INCREMENT,
    `prefix`  varchar(256) NOT NULL,
    `name`    varchar(256) DEFAULT NULL,
    `created` timestamp    NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `fields`
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT,
    `form_id`     int(11)      NOT NULL,
    `type`        varchar(64)  NOT NULL,
    `name`        varchar(256) NOT NULL,
    `description` text         NOT NULL,
    `label`       text    DEFAULT NULL,
    `hasComment`  int(1)       NOT NULL,
    `priority`    int(11)      NOT NULL,
    `weight`      int(11) DEFAULT NULL,
    `choices`     text    DEFAULT NULL,
    `required`    int(1)  NOT NULL DEFAULT 0,
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
    `id`              int(11)     NOT NULL AUTO_INCREMENT,
    `name`            varchar(64) NOT NULL,
    `description`     text,
    `created`         timestamp   NOT NULL,
    `publish`         datetime    NOT NULL,
    `deadline`        datetime    NOT NULL,
    `course_id`       int(11)     NOT NULL,
    `submission_id`   int(11)  DEFAULT NULL,
    `review_enabled`  tinyint(1) NOT NULL DEFAULT 0,
    `review_id`       int(11)  DEFAULT NULL,
    `review_deadline` datetime DEFAULT NULL,
    `validation_id`   int(11)  DEFAULT NULL,
    `reviewers`       int(11)  DEFAULT NULL,
    `group_delivery`  int(1)   NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`course_id`) REFERENCES course (`id`),
    FOREIGN KEY (`submission_id`) REFERENCES submissions (`id`)
);

CREATE TABLE `peer_reviews`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `assignment_id`  int(11) NOT NULL,
    `user_id`        int(11) NOT NULL,
    `review_user_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`assignment_id`) REFERENCES assignments (`id`),
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
    `name`          text        NOT NULL,
    `label`         text        NOT NULL,
    `answer`        text        NOT NULL,
    `comment`       text    DEFAULT NULL,
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
    `name`          text        NOT NULL,
    `label`         text        NOT NULL,
    `answer`        text            NULL,
    `comment`       text    DEFAULT NULL,
    `submitted`     timestamp   NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`assignment_id`) REFERENCES assignments (`id`),
    FOREIGN KEY (`submission_id`) REFERENCES submissions (`id`)
);

CREATE TABLE `validation`
(
    `id`        int(11)     NOT NULL AUTO_INCREMENT,
    `hash`      varchar(64) NOT NULL,
    `user_id`   int(11)              DEFAULT NULL,
    `valid`     tinyint(1)  NOT NULL DEFAULT 1,
    `timestamp` datetime    NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `users_pending`
(
    `id`            int(11)     NOT NULL AUTO_INCREMENT,
    `name`          varchar(64) DEFAULT NULL,
    `email`         varchar(64) NOT NULL,
    `password`      varchar(64) DEFAULT NULL,
    `validation_id` int(11)     NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`validation_id`) REFERENCES validation (`id`)
);

-- not attaching any foreign keys to log because we always want it to log something even if some of the data is missing
CREATE TABLE `logs`
(
    `id`               int(11)  NOT NULL AUTO_INCREMENT,
    `user_id`          int(11)  NOT NULL,
    `timestamp`        datetime NOT NULL,
    `activity`         int(11)  NOT NULL,
    `assignment_id`    int(11) DEFAULT NULL,
    `course_id`        int(11) DEFAULT NULL,
    `submission_id`    int(11) DEFAULT NULL,
    `review_id`        int(11) DEFAULT NULL,
    `group_id`         int(11) DEFAULT NULL,
    `old_value`        text    DEFAULT NULL,
    `new_value`        text    DEFAULT NULL,
    `affected_user_id` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `groups`
(
    `id`            int(11)         NOT NULL AUTO_INCREMENT,
    `assignment_id` int(11)         NOT NULL,
    `name`          varchar(255)    NOT NULL,
    `user_id`       int(11)         NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`assignment_id`) REFERENCES assignments (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users (`id`)
);

CREATE TABLE `user_groups`
(
    `id`            int(11) NOT NULL AUTO_INCREMENT,
    `user_id`       int(11) NOT NULL,
    `group_id`      int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users (`id`),
    FOREIGN KEY (`group_id`) REFERENCES groups (`id`)
);
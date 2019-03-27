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

INSERT INTO course (id, hash, coursecode, coursename, teacher, description, year, semester) VALUES (4, 'bi4d2164gh0gbb7r94qg', 'IMT2681', 'Cloud Technologies', 3, '# IMT2681 Cloud Technologies', 2019, 'fall');

INSERT INTO usercourse (userid, courseid) VALUES (3, 4);
INSERT INTO usercourse (userid, courseid) VALUES (4, 4);
INSERT INTO usercourse (userid, courseid) VALUES (5, 4);
INSERT INTO usercourse (userid, courseid) VALUES (6, 4);
INSERT INTO usercourse (userid, courseid) VALUES (7, 4);
INSERT INTO usercourse (userid, courseid) VALUES (8, 4);
INSERT INTO usercourse (userid, courseid) VALUES (9, 4);
INSERT INTO usercourse (userid, courseid) VALUES (10, 4);
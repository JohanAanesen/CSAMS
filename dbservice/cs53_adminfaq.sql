INSERT INTO cs53.adminfaq (id, timestamp, questions) VALUES (1, '1997-02-13 13:37:00', 'Q: How do I make a course + link?
--------------------------------
**A:** Dashboard -> Courses -> new. And create the course there

Q: How do I make an assignment?
--------------------------------
**A:** Dashboard -> Assignments-> new. And create the assignment there

Q: How do I invite students to the course?
--------------------------------
**A:** Create a link for the course and email the students the link

Q: How do I import database?
--------------------------------
**A:** Start xampp and go to import in phpmyadmin

Q: How do I export database?
--------------------------------
**A:** Start xampp and go to export in phpmyadmin

Q: How do I sign up?
--------------------------------
**A:** You go to `/register` and register a user there

![Reddit](https://external-preview.redd.it/lzcL5WbUuBr7pI9zIM9ZbUSrETZR1UNb-g6C5DehYss.jpg?width=960&crop=smart&auto=webp&s=4b483a024ac9103bfe6df2e98599043bbed29146)');
INSERT INTO cs53.assignments (id, name, description, created, publish, deadline, course_id, submission_id, review_id, validation_id, reviewers) VALUES (2, 'Assignment 1', '# Assignment 1: in-memory IGC track viewer

## About

Develop an online service that will allow users to browse information about IGC files. IGC is an international file format for soaring track files that are used by paragliders and gliders. The program will not store anything in a persistent storage. Ie. no information will be stored on the server side on a disk or database. Instead, it will store submitted tracks in memory. Subsequent API calls will allow the user to browse and inspect stored IGC files.

For the development of the IGC processing, you will use an open source IGC library for Go: [goigc](https://github.com/marni/goigc)

The system must be deployed on either Heroku or Google App Engine, and the Go source code must be available for inspection by the teaching staff (read-only access is sufficient).

## Specification

### General rules

The **igcinfo** should be the root of the URL API. The package and the project repo name can be arbitrary, yet the name must be meaningful. If it is called assignment1 or assignment_1 or variation of this name, we will not mark it.

The server should respond with 404 when asked about the root. The API should be mounted on the **api** path. All the REST verbs will be subsequently attached to the /igcinfo/api/* root.

```
http://localhost:8080/igcinfo/               -> 404
http://localhost:8080/<rubbish>              -> 404
http://localhost:8080/igcinfo/api/<rubbish>  -> 404
```

**Note:** the use of `http://localhost:8080` serves only a demonstration purposes. You will have your own URL from the provider, such as Heroku. `<rubbish>` represents any sequence of letters and digits that are not described in this specification.


### GET /api

* What: meta information about the API
* Response type: application/json
* Response code: 200
* Body template

```
{
  "uptime": <uptime>
  "info": "Service for IGC tracks."
  "version": "v1"
}
```

* where: `<uptime>` is the current uptime of the service formatted according to [Duration format as specified by ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Durations). 




### POST /api/igc

* What: track registration
* Response type: application/json
* Response code: 200 if everything is OK, appropriate error code otherwise, eg. when provided body content, is malformed or URL does not point to a proper IGC file, etc. Handle all errors gracefully. 
* Request body template

```
{
  "url": "<url>"
}
```

* Response body template

```
{
  "id": "<id>"
}
```

* where: `<url>` represents a normal URL, that would work in a browser, eg: `http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc` and `<id>` represents an ID of the track, according to your internal management system. You can choose what format <id> should be in your system. The only restriction is that it needs to be easily used in URLs and it must be unique. It is used in subsequent API calls to uniquely identify a track, see below.


### GET /api/igc

* What: returns the array of all tracks ids
* Response type: application/json
* Response code: 200 if everything is OK, appropriate error code otherwise. 
* Response: the array of IDs, or an empty array if no tracks have been stored yet.

```
[<id1>, <id2>, ...]
```

### GET /api/igc/`<id>`

* What: returns the meta information about a given track with the provided `<id>`, or NOT FOUND response code with an empty body.
* Response type: application/json
* Response code: 200 if everything is OK, appropriate error code otherwise. 
* Response: 

```
{
"H_date": <date from File Header, H-record>,
"pilot": <pilot>,
"glider": <glider>,
"glider_id": <glider_id>,
"track_length": <calculated total track length>
}
```

### GET /api/igc/`<id>`/`<field>`

* What: returns the single detailed meta information about a given track with the provided `<id>`, or NOT FOUND response code with an empty body. The response should always be a string, with the exception of the calculated track length, that should be a number.
* Response type: text/plain
* Response code: 200 if everything is OK, appropriate error code otherwise. 
* Response
   * `<pilot>` for `pilot`
   * `<glider>` for `glider`
   * `<glider_id>` for `glider_id`
   * `<calculated total track length>` for `track_length`
   * `<H_date>` for `H_date`


## Resources

* [Go IGC library](https://github.com/marni/goigc)


## Formal aspects
Teaching staff Git usernames

| name | github | bitbucket | prod3 | gitlab.com | other |
| ---- | ------ | --------- | ----- | ----- | ---- |
| GTL  |  gtl-hig | gtl-gjovik |  gtl | gtl-hig | ask |
| Mariusz | marni | nowostawski | mariusz | marni   | ask |
| Christopher | chrfrantz  | cfrantz | frantz | frantz | ask |
| Thomas | tholok97 | tholok97 | tholok | N/A | ask |
| Bjorn  | bjornkau   | BkAune | bjornkau | N/A | ask |

This assignment is worth 10% of 100% for the course. Note, the internal portfolio is worth 40%, thus this assignment is worth 25% of the internal portfolio. 

# Submission

The submission deadline is **Sunday, October 14th, 23:59**. No extensions will be given for late submissions. 

**[Please use this form for your submission](https://goo.gl/forms/kz90umHAJd6Bru1W2)**  <br>Note - in case of repeated submissions we take only the last one into account.


# Peer Review

* All submissions must have their peer-review entry in the spreadsheet.
* All students are expected to peer-review (at least) 2 other submissions.
* Do not mess up with the spreadsheet - do not delete, modify, or manipulate the data.
* Everyone should check the integrity of the data - any violations, please report to staff.
* Deadline: Friday, 26th of October, 06:00 (6am), morning. Extended to Tuesday, November 20th, 13:00.

[The form for peer-review](https://docs.google.com/spreadsheets/d/1Iat2up_Ra1hokvkZZYE0NJJ3JeB8iTcvd3zp0VS7SCQ/edit?usp=sharing)

', '2019-03-13 10:46:07', '2019-03-13 09:45:00', '2019-03-14 23:59:00', 4, 6, 3, null, 1);
INSERT INTO cs53.course (id, hash, coursecode, coursename, teacher, description, year, semester) VALUES (4, 'bi4d2164gh0gbb7r94qg', 'IMT2681', 'Cloud Technologies', 3, '# IMT2681 Cloud Technologies 
This page is the starting point for all lecture material of the course Cloud Technologies Course IMT2681, Autumn Semester 2018, taught at NTNU, Gjøvik.

For general course and timetable information, please look [here](https://www.ntnu.edu/studies/courses/IMT2681#tab=omEmnet).

Please check this page frequently, as all updated materials will be posted here.

# Dates
* October 14th, Sunday, 23:59. Deadline for [Assignment 1](assignment-1)
* October 29th, Monday, 23:59. Deadline for [Assignment 2](assignment-2)
* November 20th, Tuesday, 13:00. Deadline for [Assignment 3](assignment-3)

15% of the internal portfolio marks (i.e., 6% of the total marks) are reserved for feedback and participation.  

# Video recordings
We will stream this course via Youtube using the on [GTL Youtube](https://www.youtube.com/channel/UCoChi3eU2ThtVkhIpq_F20Q) channel. The videos are further organized and archived in an [IMT2681 Cloud Computing playlist](https://www.youtube.com/playlist?list=PL17KQCa8hhvA57JlDGpIv0604rK5ToP0e). 

# Assignments
* [Assignment 1](assignment-1). Submission deadline: October 14th, Sunday, 23:59.
* [Assignment 2](assignment-2). Submission deadline: October 29th, Monday, 23:59.
* [Assignment 3](assignment-3). Deadline for being **peer-reviewed**: November 20th, Tuesday, 13:00.

* [Coding challenge 1](code-challenge-1) [OPTIONAL] - results: October 23rd, Tuesday. Please submit your entry to Mariusz prior to the lecture, as a git repo URL, via email or DM on Discord.
* [Coding challenge 2](code-challenge-2) [OPTIONAL] - results: November 13th, Tuesday. Please submit your entries as pull-requests to [the exchange-demo repo](http://prod3.imt.hig.no/teaching/imt2681-2018/tree/master/exchange-demo)

# Evaluation Group
* [Course feedback](course_feedback_2018)

# Lectures
* August 21, Lecture 1: Introduction
   * [Lecture video](https://youtu.be/7tMLQ0Xk7R8)
   * [Lecture slides](https://docs.google.com/presentation/d/1F9RLAVeqw0W3P12kG4-KQ6jKGMiyslO5tYMVnwmPVH0/edit?usp=sharing)
   * [Survey URL](https://goo.gl/forms/gU3R0wMNdHH0sJp43)
   * [Short video intro to cloud computing](https://youtu.be/36zducUX16w)

* August 24, Lecture 2: Iaas, PaaS, SaaS, Intro to Linux and CLI part 1/2
   * [Lecture video](https://youtu.be/OHCBwyIcc-0)
   * [Lecture slides](https://docs.google.com/presentation/d/1T0G6JNu5_A2deWaeufegffTmn03MEb1w6zf5i3Z4T60/edit?usp=sharing)
   * Environment [setup notes](https://docs.google.com/presentation/d/1xHEp0No5mE4buKd-fzMhNJjV26M8SlkzP2yPe18x6AE/edit?usp=sharing)

* August 28, Practical session
   * Setting up the environment. No stream, no lecture.

* August 31, Lecture 3: Programming with Go, Introduction, part 1/3. Motivation, toolchain, modules, and basic syntax.

   * [Setting up Vim for programming in Go](https://youtu.be/p8rFBiDzhnM)
   * [Setting up IntelliJ for Go (2017)](https://youtu.be/amaW0jJM9Us)
   * [Setting up IntelliJ for Go (updated, 2018)](https://youtu.be/vP7Fkq6l9BI)
   * [[Go Programming Tasks 1]]
   * [Lecture video 2017](https://youtu.be/6ewYPDC4HmQ)
   * [Lecture video 2018](https://youtu.be/eJXNRD6DzmQ)
   * [Lecture slides](https://docs.google.com/presentation/d/1Pbu47nD-XG4owUejkTioNNyhQOqgQcMXHDYynuVADR8/edit?usp=sharing)


* September 4th, Lecture 4: Programming with Go, Introduction, part 2 and 3/3. Data types, functions, parameters, value vs. pointer, lambda functions, higher order functions.

   * [Lecture video 2017](https://youtu.be/L9KObCRp9PU)
   * [Lecture video 2018](https://youtu.be/wnBYx1XwPaY)
   * [Lecture slides](https://docs.google.com/presentation/d/1Pbu47nD-XG4owUejkTioNNyhQOqgQcMXHDYynuVADR8/edit?usp=sharing)


* September 7th, Lecture 5: Programming with Go, Introduction, part 3/3. Interfaces, concurrency, and intro to JSON and error handling.

   * [Lecture video 2017](https://youtu.be/L9KObCRp9PU) (Same as last week)
   * [Lecture video 2018](https://youtu.be/wFRgoRcQjJk)
   * [Tutorial on Go error handling](https://youtu.be/m3imnk1ZVLY)

* September 11th, Lecture 6: Introduction to Linux, part 2/2
   
   * [Lecture slides](https://docs.google.com/presentation/d/1T0G6JNu5_A2deWaeufegffTmn03MEb1w6zf5i3Z4T60/edit?usp=sharing)
   * [Lecture Video Part 1](https://youtu.be/oZXwmHVEfCg) [Lecture Video Part 2](https://youtu.be/R44FfvnZbIo)
   * [Vim Cheat Sheet](https://github.com/chrfrantz/vim-cheatsheet/blob/master/vimCheatSheet.pdf)

* September 14th, Lecture 7: Linux ctd. (Processes, Software Installation), REST
   * [Lecture Slides Software Installation on Linux](https://drive.google.com/file/d/1R0Vd9yY_YmoLULRgtQ9g_xluDYPgE6Kx/view?usp=sharing)
   * [Lecture Slides REST](https://drive.google.com/file/d/12S_f5JRLXA_uHPynW3K5SNtwPvubh8It/view?usp=sharing)
   * [Lecture Video Part 1](https://youtu.be/TtEgYY7RPBc)
   * [Lecture Video Part 2](https://youtu.be/IE4c8AAGT4E)

* September 18th, Lecture 8: Testing, Testing strategies, TDD. Testing in Golang.
   * [Lecture video, part 1](https://youtu.be/osgufr4r29s) - simple funcs and Testing.
   * [Lecture video, part 2](https://youtu.be/eDuXMXXETeY) - simple GET request.
   * [Programming tasks 2](Go-Programming-Tasks-2)

* September 21st, NO CLASS. Video Lecture(s) 9: Programming in Go, REST, Client-Server.
   * [Lecture video, part 1](https://youtu.be/ohiuGGWVzPI)
   * [Lecture video, part 2](https://youtu.be/bL8-YNEkPR0)
   * [Lecture video, part 3](https://youtu.be/iyR9A0wImfI)

* September 25th, TUTORIAL. Programming in Go, REST, Client-Server. JSON parsing. Semantic versioning and Golang modules. 
   * [Programming tasks 3](Go-Programming-Tasks-3)
   * [Lecture video](https://www.youtube.com/watch?v=MyQdzQo380Q)
   * Note: `go mod init <name_of_your_module>`

* September 28th, No class. Video Lecture(s) 10: Regular Expressions
   * [Lecture video 1: Regular Expressions Overview](https://youtu.be/MiEqkque2bc)
   * [Slides](https://drive.google.com/file/d/1NceKJqXfN9gRNI5Sz6Sew6DRb36q_tmz/view?usp=sharing)

   * Note: This video contains some exercises to explore regular expressions.
   * Example use cases (interest-based)
     * [Video: Web Server Request Routing in Go using Regular Expressions](https://youtu.be/qumPl10DnN4)
       
      * Note: This video contains some tasks to practice the use of regular expressions in Golang.

      * [Video: Performance comparison string search vs. regex](https://youtu.be/2_sq-38v4ZM?t=11m14s)
      * Videos: Regex-based validation in web service [Part 1](https://youtu.be/2_sq-38v4ZM?t=28m12s) (server-side in Go), [Part 2](https://youtu.be/f9V15HSmcpk) (Client-side using JavaScript & HTML 5)

* October 2nd: Lecture: Networking Fundamentals (ISO OSI Model, Layers)
  * [Slides](https://drive.google.com/file/d/1aXdXPQ7ZAMdr6hFW1PKPyAHi8MIqQ0_q/view?usp=sharing)
  * Videos: [Part 1](https://youtu.be/UTPe3qcOZ5k) [Part 2](https://youtu.be/qeJYJieKaRA)
  * Network Trace [Download](https://drive.google.com/file/d/1dXr_T83yCnmTBnbHHMcRvx4cxdRxaSzC/view?usp=sharing)

* October 5th: Lecture: Networking Fundamentals ctd. (Network Addressing, Subnetting)
  * [Slides](https://drive.google.com/file/d/1N1i3nGMKMxfTV_8O6lcVBuj3dwvRpMeH/view?usp=sharing)
  * Videos: [Part 1](https://youtu.be/v3EZTZDSdis) [Part 2](https://youtu.be/spq6OdxIeaw)
  * For links to OpenStack, VPN connection, putty, key conversion, etc. see Resources, Tools and Services sections below

* October 9th: Lecture: IaaS
  * Videos: [Part 1](https://youtu.be/4qgRAeGFsRg), [Part 2](https://youtu.be/pzKmfb8XKKg)
  * [Slides](https://drive.google.com/file/d/1cLf6B6oWBv-dt1wsJyf5M_MrL364nCAT/view?usp=sharing)

* October 12th: Virtualisation
  * Videos: [Virtualisation Introduction](https://youtu.be/mt4FUaqVSfk)
   
   * Note to everyone that couldn''t make it to class. This is a recording of the lecture, since the internet connection in A255 failed during the stream. The lecture was shorter since the remainder of the class time was used for assignment support.
  
  * [Slides](https://drive.google.com/file/d/1j2MJGzaZMBdt2KVBXD2012grByUuz9dd/view?usp=sharing)

  * Additional Video (from last year): Dockerising your applications [Part 1](https://youtu.be/OB5WViLQRz4) [Part 2](https://youtu.be/yllsZf6Hkdo)


* October 16th: WebHooks, Introduction to MongoDB

  * [WebHooks slides](https://docs.google.com/presentation/d/1Tl2uk4A20wKTU9mPTrPVloiwh11epKeGt2W2ogiiX3o/edit#slide=id.p1)
  * [MongoDB intro](https://docs.google.com/presentation/d/109t8q6YzRBLv4iTBKkI-whTRlXe41wG1TR4j6_cfCuo/edit#slide=id.p)
  * [Lecture video](https://youtu.be/qXrEmBpUxKk)


* October 19th: MongoDB plus Go interfaces (video lectures). Peer review explained. Q and A session. 

   * [more MongoDB explanations](https://www.youtube.com/watch?v=WmQcIKR2s2U&index=32&list=PL17KQCa8hhvA57JlDGpIv0604rK5ToP0e&t=0s)
   * [Go interfaces](https://www.youtube.com/watch?v=FI4glTiNR6U&index=33&list=PL17KQCa8hhvA57JlDGpIv0604rK5ToP0e&t=0s)
   * [Peer Review explained]() -- coming soon
   

* October 23rd: Coding Challenge 1, discussion. Introduction to AWS.

   * [Lecture slides](https://docs.google.com/presentation/d/17ltaVfTnUy9-LdFYMx2fffmkQcJgjBdJYsplLJnTPoc/edit?usp=sharing)
   * [Lecture video](https://youtu.be/CXfUm5BCEj4)


* October 26th:  AWS

   * [AWS](https://aws.amazon.com) The actual AWS portal, describing all the available services.
   * [Introduction to Amazon Web Services by Jeff Barr](https://www.youtube.com/watch?v=CaJCmoGIW24) This is a lecture, with more general intro to Cloud computing. Worth a watch, but, limited detailed know-how on AWS.
   * [AWS services overview](https://www.sumologic.com/aws/) This provides more details on AWS architecture and services in form of videos and tutorial. 
   * There is a large number of online resources and courses on AWS. If you find something really good, please post it here or on the issue tracker.


* October 30th: AWS Lambda

   * [Lecture slides](https://docs.google.com/presentation/d/17ltaVfTnUy9-LdFYMx2fffmkQcJgjBdJYsplLJnTPoc/edit?usp=sharing)
   * [Lecture video](https://youtu.be/sSZaeT2uDMU)
   * [AWS](https://aws.amazon.com)
   * [API Gateway Proxy](https://github.com/aws/aws-lambda-go/blob/master/events/apigw.go)

* November 2nd: Docker compose
   * [Lecture Slides](https://drive.google.com/file/d/1ucGxhJg-oTI4V3GgB9FNlbcXbS_3iC1l/view?usp=sharing)
   * [Slides from previous docker lecture (related to discussion)](https://drive.google.com/file/d/1j2MJGzaZMBdt2KVBXD2012grByUuz9dd/view?usp=sharing)
   * [Lecture video, part 1](https://youtu.be/xliynkcOJOc)
   * [Lecture video, part 2](https://youtu.be/rTGGwyEizvk)

* November 6th, Docker compose and Multi-stage builds
   * [Lecture video, part 1 (docker compose example)](https://youtu.be/o7yzPnhZyoo)
   * [Lecture video, part 2 (multi-stage builds)](https://youtu.be/8Zhcqn-mpaY)

* November 9th, SLAs and Scalability
   * [Lecture video, part 1](https://youtu.be/oyxyJvqHybE)
   * [Lecture video, part 2](https://youtu.be/4bQq14BxwgA)
   * [Lecture slides](https://drive.google.com/file/d/1AlQHGuH3YpiwfGrn4HWzPagp_522Nyw-/view?usp=sharing)

# Resources
* [Code examples: client/server, DB, testing](https://github.com/marni/imt2681_cloud)
* [Conversion of .pem file to .ppk file for access using putty under Windows](https://stackoverflow.com/questions/3190667/convert-pem-to-ppk-file-format) Note: Yes, it is StackOverflow. But the description is pretty good. Start at point 3 of the instructions; the stuff before refers to generation of .pem files in vCloud (VMware''s private cloud solution, which we don''t use).

# Tools
* [putty, puttyGen, pscp, Pageant, etc.](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html)
* Docker:
  * [Documentation](https://docs.docker.com/)
  * [Dockerfile examples](https://docs.docker.com/samples/)
  * [Dockerfile Best Practices](https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices/)
  * [Docker Book](https://www.dockerbook.com/) If you are looking for a textbook on docker, this one is a great option.

# Services
* [OpenStack Login (SkyHIGh)](https://skyhigh.iik.ntnu.no/horizon/auth/login/?next=/horizon/project/)
  * [General Documentation](https://www.ntnu.no/wiki/display/skyhigh/Openstack+documentation)
  * [Starting point for setting up first OpenStack configuration](https://www.ntnu.no/wiki/display/skyhigh/Initial+setup%2C+using+the+webinterface?src=contextnavpagetreemode)
  * Note: For the use outside of the NTNU network, you will need to connect via VPN first. More information on how to do that can be found [here](https://innsida.ntnu.no/wiki/-/wiki/English/Install+VPN).
* Serverless Computing
   * AWS alternative: [Serverless.com](https://serverless.com)
   * Emulated AWS Lambda environment for local installation (based on Docker): [Docker-Lambda](https://github.com/lambci/docker-lambda)

# External talks
* [Short talk about TDD](https://www.youtube.com/watch?v=a6oP24CSdUg)
* [Good and Bad points of Go](https://www.quora.com/What-are-the-strengths-and-weaknesses-of-Golang/answer/Roman-Scharkov) - pretty detailed and accurate discussion on strengths and weaknesses of Go.

# On learning
* [Struggle means learning](https://www.kqed.org/mindshift/24944/struggle-means-learning-difference-in-eastern-and-western-cultures)
* [Growth vs. Fixed mindset](https://www.youtube.com/watch?v=KUWn_TJTrnU)
* [On Challenging Learning](https://www.youtube.com/watch?v=KUWn_TJTrnU) - stepping outside the comfort zone, throwing oneself into challenging tasks, learning from errors

# Q & A
* How to setup Go for Visual Studio Code (on windows) [answer](http://prod3.imt.hig.no/teaching/imt2681-2018/issues/2)
* What are the benefits of using Go versus other more established languages? [answer](https://www.quora.com/What-are-the-benefits-of-using-GoLang-versus-other-more-established-languages)
* Is Go better than C? [answer](https://www.quora.com/Is-Golang-better-than-C/answer/Richard-Kenneth-Eng)', 2019, 'fall');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (42, 8, 'url', 'git_repository_w_comment_url_35', 'Git Repository', 1, 'Make sure the repository is public before delivering!', 0, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (43, 8, 'textarea', 'git_repository_w_comment_textarea_0', 'textarea', 0, '', 1, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (44, 8, 'url', 'git_repository_w_comment_url_1', 'url', 0, '', 2, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (45, 8, 'number', 'git_repository_w_comment_number_2', 'number', 0, '', 3, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (46, 8, 'checkbox', 'git_repository_w_comment_checkbox_3', 'checkbox', 0, '', 4, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (47, 8, 'radio', 'git_repository_w_comment_radio_4', 'radio', 0, '', 5, 0, 'a,b,c');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (48, 9, 'text', 'test_all_fields_text_36', 'Text', 1, 'This is a text field', 0, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (49, 9, 'textarea', 'test_all_fields_textarea_37', 'Textarea', 0, 'This is a long text field', 1, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (50, 9, 'url', 'test_all_fields_url_38', 'URL', 0, 'This is a URL field', 2, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (51, 9, 'number', 'test_all_fields_number_39', 'Number', 0, 'This is a number field', 3, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (52, 9, 'checkbox', 'test_all_fields_checkbox_40', 'Checkbox', 0, 'This is a checkbox field', 4, 0, '');
INSERT INTO cs53.fields (id, form_id, type, name, label, hasComment, description, priority, weight, choices) VALUES (53, 9, 'radio', 'test_all_fields_radio_41', 'Radio', 0, 'This is a radio field', 5, 0, 'a,b,c');
INSERT INTO cs53.forms (id, prefix, name, created) VALUES (8, 'git_repository_w_comment', 'Git Repository w/Comment', '2019-03-13 10:52:52');
INSERT INTO cs53.forms (id, prefix, name, created) VALUES (9, 'test_all_fields', 'TEST: All fields', '2019-03-13 10:54:02');
INSERT INTO cs53.logs (userid, timestamp, activity, assignmentid, courseid, submissionid, oldvalue, newValue) VALUES (3, '2019-03-13 10:44:36', 'COURSE-CREATED', null, 4, null, null, null);
INSERT INTO cs53.logs (userid, timestamp, activity, assignmentid, courseid, submissionid, oldvalue, newValue) VALUES (3, '2019-03-13 10:44:36', 'JOINED-COURSE', null, 4, null, null, null);
INSERT INTO cs53.peer_reviews (id, submission_id, assignment_id, user_id, review_user_id) VALUES (4, 6, 2, 4, 5);
INSERT INTO cs53.peer_reviews (id, submission_id, assignment_id, user_id, review_user_id) VALUES (5, 6, 2, 4, 6);
INSERT INTO cs53.peer_reviews (id, submission_id, assignment_id, user_id, review_user_id) VALUES (6, 6, 2, 5, 4);
INSERT INTO cs53.peer_reviews (id, submission_id, assignment_id, user_id, review_user_id) VALUES (7, 6, 2, 5, 6);
INSERT INTO cs53.peer_reviews (id, submission_id, assignment_id, user_id, review_user_id) VALUES (8, 6, 2, 6, 4);
INSERT INTO cs53.peer_reviews (id, submission_id, assignment_id, user_id, review_user_id) VALUES (9, 6, 2, 6, 5);
INSERT INTO cs53.reviews (id, form_id) VALUES (3, 9);
INSERT INTO cs53.submissions (id, form_id) VALUES (6, 8);
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (5, 4, 5, 3, 2, 'text', 'test_all_fields_text_0', 'Text', 'input text', null, '2019-03-13 12:25:49');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (6, 4, 5, 3, 2, 'textarea', 'test_all_fields_textarea_1', 'Textarea', 'input textarea', null, '2019-03-13 12:25:49');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (7, 4, 5, 3, 2, 'url', 'test_all_fields_url_2', 'URL', 'http://vg.no', null, '2019-03-13 12:25:49');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (8, 4, 5, 3, 2, 'number', 'test_all_fields_number_3', 'Number', '1337', null, '2019-03-13 12:25:49');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (9, 4, 5, 3, 2, 'checkbox', 'test_all_fields_checkbox_4', 'Checkbox', 'on', null, '2019-03-13 12:25:49');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (10, 4, 5, 3, 2, 'radio', 'test_all_fields_radio_5', 'Radio', '2', null, '2019-03-13 12:25:49');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (11, 4, 6, 3, 2, 'text', 'test_all_fields_text_36', 'Text', 'a', null, '2019-03-14 13:44:34');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (12, 4, 6, 3, 2, 'textarea', 'test_all_fields_textarea_37', 'Textarea', 'b', null, '2019-03-14 13:44:34');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (13, 4, 6, 3, 2, 'url', 'test_all_fields_url_38', 'URL', 'http://vg.no', null, '2019-03-14 13:44:34');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (14, 4, 6, 3, 2, 'number', 'test_all_fields_number_39', 'Number', '1', null, '2019-03-14 13:44:34');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (15, 4, 6, 3, 2, 'checkbox', 'test_all_fields_checkbox_40', 'Checkbox', '', null, '2019-03-14 13:44:34');
INSERT INTO cs53.user_reviews (id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (16, 4, 6, 3, 2, 'radio', 'test_all_fields_radio_41', 'Radio', '', null, '2019-03-14 13:44:34');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (13, 5, 2, 6, 'url', 'http://www.go.no', 'asdf', '2019-03-13 16:32:43');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (14, 5, 2, 6, 'textarea', 'asdf', '', '2019-03-13 16:32:43');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (15, 5, 2, 6, 'url', 'http://www.go.no', '', '2019-03-13 16:32:43');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (16, 5, 2, 6, 'number', '1332', '', '2019-03-13 16:32:43');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (17, 5, 2, 6, 'checkbox', '', '', '2019-03-13 16:32:43');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (18, 5, 2, 6, 'radio', '1', '', '2019-03-13 16:32:43');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (19, 4, 2, 6, 'url', 'http://github.com', '', '2019-03-14 13:46:12');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (20, 4, 2, 6, 'textarea', 'lorem upsum', '', '2019-03-14 13:46:12');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (21, 4, 2, 6, 'url', 'http://vg.no', '', '2019-03-14 13:46:12');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (22, 4, 2, 6, 'number', '42', '', '2019-03-14 13:46:12');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (23, 4, 2, 6, 'checkbox', 'on', '', '2019-03-14 13:46:12');
INSERT INTO cs53.user_submissions (id, user_id, assignment_id, submission_id, type, answer, comment, submitted) VALUES (24, 4, 2, 6, 'radio', '1', '', '2019-03-14 13:46:12');
INSERT INTO cs53.usercourse (userid, courseid) VALUES (3, 4);
INSERT INTO cs53.usercourse (userid, courseid) VALUES (4, 4);
INSERT INTO cs53.usercourse (userid, courseid) VALUES (5, 4);
INSERT INTO cs53.usercourse (userid, courseid) VALUES (6, 4);
INSERT INTO cs53.usercourse (userid, courseid) VALUES (7, 4);
INSERT INTO cs53.usercourse (userid, courseid) VALUES (8, 4);
INSERT INTO cs53.usercourse (userid, courseid) VALUES (9, 4);
INSERT INTO cs53.usercourse (userid, courseid) VALUES (10, 4);
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (1, 'Ken Thompson', 'hei@gmail.com', 1, 'mannen@harmannenfalt.no', '$2a$14$MZj24p41j2NNGn6JDsQi0OsDb56.0LcfrIdgjE6WmZzp58O6V/VhK');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (2, 'Frode Haug', 'frodehg@teach.ntnu.no', 1, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (3, 'Ola Nordmann', 'olanor@stud.ntnu.no', 1, 'swag-meister69@ggmail.com', '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (4, 'Johan Klausen', 'johkl@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (5, 'Stian Fjerdingstad', 'stianfj@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (6, 'Svein Nilsen', 'sveini@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (7, 'Kjell Are-Kjelterud', 'kjellak@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (8, 'Marius Lillevik', 'mariuslil@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (9, 'Jorun Skaalnes', 'jorunska@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
INSERT INTO cs53.users (id, name, email_student, teacher, email_private, password) VALUES (10, 'Klaus Aanesen', 'klausaa@stu.ntnu.no', 0, null, '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');
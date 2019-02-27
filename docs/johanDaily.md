Johan's logbook :)
==================

Tuesday 29/1
------------
* I created the database model through phpmyadmin, worked fine I guess although the final specs are not complete
* Distributed tasks for the first week or two on trello :)
* Sick af, no really

Wednesday 30/1
------------
* I started early because I am a straight up boss, actually I caught a cold and I'm sick and got 
nothing better to do
* Started with card #3 Router
* Made all the basic endpoints and tests, ran into some issued where the test would run as if it 
was inside the /handlers folder, and thus the template couldn't find /web/test.html. Got fixed by
a small init function which ships the test suite back.
```
func init() {
 	if err := os.Chdir("../"); err != nil { //go out of /handlers folder
 	    panic(err)
 	}
 }
```
* Found out a designated error page would be nice so I added that to the UML and router too

Monday 4/2
------------
* Started doing the card #14 database connection class
* Decided on a global approach to DB access, aka all the files that import the db package can access
the global DB (dbcon*) variable :)
* https://www.alexedwards.net/blog/organising-database-access
* Svein reminded me with need some sort of session management, although GO doesn't include it's own
session functionality
* Could implement the functionality myself by creating unique values on cookies, and storing that
value in the database, although this would cost IO operations
* Decided on going with the Sessions package from Gorilla
* https://www.gorillatoolkit.org/pkg/sessions
* Together with the more powerful mux router package and secureCookies package, all part of the Gorilla
Web Toolkit
* Seems like the go-to toolkit for most people on stack overflow
* Password encryption added with bcrypt https://godoc.org/golang.org/x/crypto/bcrypt

Tuesday 5/2
-----------
* Link to execute go commands on scripts: https://github.com/gojp/goreportcard/blob/master/check/utils.go
* Changed a bit on the login register things, works nicely

Wednesday 6/2
-----------
* Started the day setting up the project on my laptop, ugh
* Changed from local imports to package imports
* Wrapped up the login/register functionality, works nicely afaik
* Started on #11 Create Course
* Svein already made most of the frontend so I dont have to, score

Thursday 7/2
-----------
* 11 Create Course cont.
* Added url link fields, working on DB query to 'save' the course permanently
* Implemented input fields for 3x url's, after talking with Svein we decided
on implementing markdown for the description field, so the teacher
can just add links there.
* https://simplemde.com/
* Christopher said it was 'sensible' :)
* Kinda back on tests, working on making some now
* Changed DB functions to handle User objects instead of 3-4 fucking variables
* Now checks if user is a teacher before serving /admin pages

Friday 8/2
-----------
* 11 Create Course cont/testing
* Finished up 11, found out that the tests that needed some kind of session could be hacked because
i have the cookiestore/sessionstore stuff available in the test suite, sweet
* Helped Brede with some bugs on his 8-userpage branch
* didn't have too much time to work this day :(

Sunday 10/2
-----------
* Refactored 8
* Made more session function and moved them into the util package
* Moved DB functions out of db.go and into feks. coursedb.go and userdb.go
* Had to rewrite a lot of UserHandler code because it was messy and unreadable/unefficient

Monday 11/2
-----------
* Meeting day
* Found out it was time to do a major restructure of the code
* Restructure took the whole day :(

Tuesdag 12/2
-----------
* Restructure cont.
* Svein knew what was supposed to be done with the restructure, I had to find out by trying
* This took longer than it should have, next time we should have a better plan for what goes where and how it should work
* BEFORE starting the restructuring.. ffs
* Lots of time went into getting the god damn tests to work again

Wednesday 13/2
-----------
* Cont. on #7
* Got the md -> html functionality working :)
* This gonna look nice
* Also some more db schtuf

Thursday 14/2
-----------
* Cont #7 jesus fuck will it ever end
* Made front end look nice :)
* Changed markdown processor so it would read 'github flavored markdown' better
* yeah!

Friday 15/2
----------
* Wrapping up #7
* Added comment tab

Monday 18/2
----------
* Wrapped up the coursepage
* Meetings n stuff <3
* Started on participant list

Tuesday 19/2
----------
* Wrapped up the participant list thing, wasn't very hard to do
* Started doing research on docker containers and docker-compose kind of setups to "knit" together the different
services we are creating

Wednesday 20/2
----------
* Started writing the peer review service
* The peer review service is supposed to retrieve a request whenever, the request should include some sort of
authentication, haven't decided what yet. And what submission is supposed to be reviewed as well as how many
submissions every person is supposed to review.
* I distribute the who-reviews-what by fetching every submission and their user from database
* Randomly shuffle the array
* If everyone is reviewing 2 tasks, then every person get's the 2 next in the array.

Thursday 21/2
----------
* skiday!

Friday 22/2
----------
* Hungover day, no work was done except reviewing a pull request.

Saturday 23/2
----------
* 28 cont.
* Trying to setup Docker on my pc, had to enable hypervisor in powershell to make it work :(
```powershell
bcdedit /set hypervisorlaunchtype auto
```
* Sitting for 8 hours now trying to make the docker thingies work, but without luck. Will try again tomorrow.
* https://www.melvinvivas.com/my-first-go-microservice/
* Multi staged docker builds learned through: https://levelup.gitconnected.com/multi-stage-docker-builds-with-go-modules-df23b7f91a67
* Shows how to make small docker images for GO applications

Monday 25/2
----------
* Dockerizing cont
* IT WORKS
* I can finally say I'm starting to understand Docker, and to this useage I mean no specific article helped out a lot
* PO wanted the database to be persistent, implemented this through a Dockerfile for the dbservice which will add the 
database.sql file to the initialization folder in the container, but now it will only generate the db if it doesn't exist.
Then in the docker-compose file I set the persistent volume, works like a charm.

Tuesday 26/2
---------
* Setup OpenStack
* Deployed application with docker, a lot easier than expected
* Added Auth, POST and more functionality to the peer service
* Need a Scheduler Service to run tasks at a given time
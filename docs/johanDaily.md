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

Wednesday 27/2
---------
* Finished off the peer review service, although lacking of a duplicate request failure
* Helped Brede setup the project on both openstack and on his computer with docker installed <3
* There is a lot of duplicate files across the services, tried to move this out of the services and into a /internal folder
* No luck in making the services use those files tho
* Creating the SchedulerService
* Making it so it doesn't neccessarily only take peer scheduling request
* First bit of the request has to be the same, but from there I can add different execution
methods and functionality :)
* Using Goroutines to keep track of when to execute what services
* Switched to use the AfterFunc(duration, func) which triggers func after duration has passed.
* Perfect for this usecase

Thursday 28/2
----------
* Scheduler service is now working wonderfully, both registering through a http post request
and it actually triggering the peerservice correctly works.
* Still missing RUD out of CRUD, but that should be pretty easy to do tomorrow.
* Scheduler will not accept requests with the same submission id as a current timers submission id
* The peer service and Scheduler service both needs some polishing, but are functional for the alpha
planned on monday

Friday 1/3
----------
* Implemented GET, PUT, DELETE to schedulerservice

Saturday 2/3
-----------
* Added tests to schedulerservice
* Issues with computer, piece of crap 
* Made edit course functionality
* Changed course model to courseRepository to access db functionality of model

Sunday 3/3
----------
* Created 'see active assignments' functionality for frontpage
* Fixed course_code to show in the badges on assignments
* Fixed insert to use transaction in submissionRepo
* Had to redo shcedulerservice to also use assignmentID because apparently one submission might have several rows of input
* Made update and delete functions to schedule peer tasks in webservice
* Setting and updating reviewer count/nr functionality
* Bug fixes from hell

Monday 4/3
----------
* Fixed course create button

Tuesday 5/3
----------
* Added XSS Sanitization with the bluemonday package
    * Login
    * Register
    * Assignment submission page
* Made tabs more clearer and design neater on course page
* Added css library for Go syntax highlighting from markdown code tags

Wednesday 6/3
----------
* Added time submitted and a count of nr of submissions to admin assignment submissions page
* Fixed width issues with course and assignment boxes
* Updated design on assignment page


Thursday 7/3
-----------
* Assignment page frontend finished :)
* Added 'edit assignment' button to assignment boxes in admin dashboard

Friday 8/3
-----------
* Added invalid login message
* Changed the message to be stored in session
* Added e-mail already in use message
* Various error messages added
* Added CourseID and AssignmentID to admin page overview boxes on frontend

Monday 11/3
-----------
* Changed to using showdown.js with highlight.js to convert markdown into html code and syntax highlighting for code snippets
* Removed set header to update flash message/session message
* Added AssignmentID to assignment details

Tuesday 12/3
-----------
* Added scheduler files to see scheduled tasks in admin view
* Refactored database.sql to not include everything
* Separated out the insert to it's own database_insert.sql file
* Updated insert file with dummy data for testing and debugging
* Added a frontend for seeing data from the scheduler, taken from the api provided by schedulerservice

Wednesday 13/3
------------
* Linting errors fixed in session files
* Updated update function for assignments
* Updated scheduler to update reviewers count
* Imported sveins changes
* Updated the frontend panels for scheduler in admin view
* Added the functionality to delete scheduled tasks from the frontend

Thursday 14/3
------------
* Fixed schedule delete form
* Fixed lint errors

*VACATION until 25/3*

Tuesday 26/3
------------
* Fixed bug related to database not initiating timers in schedulerservice
* Removed submissionID from peerservice and schedulerservice as it was not needed
* Updated database table to exlude the submissionID for the same reason

Thursday 28/3
-----------
* Added reviews overview for assignments
* Refactored a large spelling mistake done to the repository folder, it was named repositroy, fixed it.
* fixed count issue for reviews done

Friday 29/3
-----------
* Moved back button to a better location in admin/assignment/review
* Added email service template for brede
* Fixed issue when seeing two reviews done on an assignment, only one of the reviews's radiobuttons would be selected.
* Solved by wrapping them in its own forms.

Tuesday 2/4
-----------
* Changed title on reviewers overview page
* Removed redundant div row tag

Thursday 4/4
-----------
* Merging done

Monday 8/4
-----------
* Mailservice merge and once again updating the repositroy folder typo

Tuesday 9/4
-----------
* Added privacy policy
* Added it as a term of registering
* During the monday meeting we decided that the automated way of the reviewing system was far too complex to adress minor issues like:
    * Removing an user from the review cycle
    * Adding a late submission to the review cycle
    * etc
* These scenarios would need it's own algorithms that I had been mocking up the past weeks, but it was a lot of work for code that would rarely be run
* Decided to use a review-pool instead, so by pushing a button you would receive a review, instead of automatically assigning all the reviews at the deadline
* Removed schedulerservice and peerservice
* Updated docker-compose
* Removed more traces of the services
* Removed scheduler overview from admin page
* Updated assignment setting with a 0 value for reviews
* Disable reviews button added to edit assignment
* Updated create assignment page with the same disable button

Wednesday 10/4
------------
* Added review-pooling functionality backend
    * This would query all submissions from the database and count all reviews done on them, then randomly distribute on of the least reviewed submissions to the user requesting to do a review.

Thursday 11/4
-----------
* Added the slice shuffle function as a util
* Removed traces of the auto validation service as we decided there was not enough time to do it.
* Peer review button active when peer review is active
* Removed penile object from comment
* Review pooling updated, tested and bug checked, working.

Friday 12/4
-----------
* Tried to figure out why assignments were no longer updating
* Rolled back some files to figure out the issue
* Was database related as well as some computer issues on my part.

Wednesday 17/4
------------
* Fixed a 'go backs' typo on a button
* Assignment page dynamically shows the users how many reviews they have left to do.

Thursday 18/4
-------------
* Buttons to do reviews removed after review deadline is over
* Reviewers count changed 0 value to 'any' so any number of reviews can be done

Sunday 21/4
--------------
* Fixed issue with simpleMDE where it wouldn't render because of a template comparison issue

Monday 22/4
------------
* Fixed a bug where 'any' reivewers would crash the submissions overview page for admins
* Removed schedulerservice link
* fixed updating review_id on assignment by deleting a foreign key
* Changed naming in assignment tabs for reviews received
* Added a error for reviews limit reached
* Made that message better.

Tuesday 23/4
------------
* Added error message when changing email to an existing email
* Added a bootstrap badge to received reviews column on assignment page, to see how many reviews the user has received on their assignment without entering the tab.
* Updated submissions forms so they can be deleted if they have no reviews done on the form yet
* Same changes done to review forms.

Wednesday 24/4
------------
* removed scheduled tasks table from database
* fixed deleting submissions forms properly
* updated function names for deleting review forms

Thursday 25/4
------------
* added review_enable column to database table for assignments
* cleared out unused functions from model/assignment
* Removed unused structs
* cleared out unused functions from model/course
* cleared out unused functions from model/review
* cleared out unused functions from model/submission
* cleared out unused functions from model/user
* cleared out unused functions from model/userSubmission
* Also fixed functions still using these to use the ones in /repository instead.
* Changed review_enabled column to normal bool, not nullable
* added enable/disable functionality to assignment, so it's really disabled even if a review_id is specified.
* added enable/disable funcitonality to update assignment page

Friday 26/4
------------
* enabling/disabling peer review functionality working for assignments creating and updating.
* If updating an assignment and the review_deadline is not set, it will be set one month after the assignment deadline
    * This won't have any impact as long as the peer reviewing is disabled
* changed the hasReview function to check the reviewEnabled variable instead.

Tuesday 30/4
------------
* Fixed changing passwords on users in admin/manage students
* Changed session messages to use the gorilla sessions.flashes instead
* Merged the brach into master.
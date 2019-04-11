# The Log of Bread
<!-- By BredeFK -->
## Week One
### Monday - 04/02/19
 *Started late today because of Super Bowl airing until 04.00 at the night.*
* Finished setting up the environment.
* Started working/preparing for task 8(User Profile).
* Watched a tutorial about bootstrap.

### Tuesday - 05/02/19
* Continued working on the design for the user profile page.
* Had a bug about not getting the data from go file to html file. I fixed it with changing the struct variables to have capital letters :)))
* Almost finished with #8, just need to get actual user information and actually change the information.
* if the user wants to change information, it happens on the same side.
* information is displayed in input element for making the changes easier.

### Wednesday - 06/02/19
*I overslept 45min today, but that's fine, I'm just working 1h 45min longer :)*
* I found a bug yesterday with adding a secondary email, that I need to fix today
* A new bug about the form in `user.html` is sending unvalidated input to the handler. I/we choose to not fix this not, but added an TODO to fix it later :) 
* We were a bit unsure how to update user user information. Since you can choose to only change one thing or all the information, should we request a change to the db for each change or all together every time?
    * Solution: Different queries to different changes
* Another bug with getting the password-hash from the DB, it only works when the secondary-email is not NULL (*I think*).
    * Solution: Ask [Johan](https://github.com/JohanAanesen) for help.
    * He kinda fixed it with `sql.Nullstring`, it still doesn't work right :/
* Today was just bug after bug, not a fun day, I also worked an hour longer so I can leave earlier on Friday.

### Thursday - 07/02/19
*I overslept 30min today, I was supposed to start one hour earlier.*
* The real courses show up on the profile now :D 
* I had to make an own function for getting the password hash because of an annoying bug >:(
* User can now change, secondary email, name and password :D
* Only need to refactor some code, write some tests and lint and add some confirmation that the information is changed
* Fixed error where the page went blank after submitting form. It was solved by requesting the view function to start again.
* Didn't work a full day today because not good :/

### Friday - 08/02/19
*I'm going home today to celebrate my birthday, so I'm leaving 2hours earlier*
* Fixed the hash bug, [Johan](https://github.com/JohanAanesen) saw that i 'fixed' it on the wrong variable so now the hash appeared too!
* Fixed the input validation bug, I had to switch from `onclick` to `onsubmit` that started the javascript script in the button-save. 
* Merged with master and got some problems, but [Johan](https://github.com/JohanAanesen) helped me fix it.
* Almost done with #8 now, only need some refactoring and I have to fix one test.

### Sunday - 10/02/19
*I'm 5h too short for this week, so I have to work some today too*
* Started up with setting up the environment on my laptop since I'm home and not in Gjøvik.
* I'm going to fix one test today and [Johan](https://github.com/JohanAanesen) is refactoring the code, so after that, I can make a pull request.
* Merged my branch to master

## Week Two
### Monday - 11/02/19
* Started on [#16 - Logging to database](https://trello.com/c/CwIxfhpk) - Log stuff that the user and/or admin does
* Everything went fine for now at least, that's nice 😍
* Had two meetings and plan to discuss new project structure
* I made a powershell script that runs go fmt,vet,lint,cyclo and test. I did it to make the linting/testing go faster and more "clean"

### Tuesday - 12/02/19
* I started on looking how the logs table will be.
* I created an struct for keeping the log data for easier use and fewer parameters.
* Logs when the user change name/email/password now.
* bug: Added foreign keys to logs, if course,submission or assignment id is blank, it doesnt work.
So I have to figure out how to send nil instead of a number to the db.
* I know one way to fix the bug, but it's to much and messy code :/ (hella many if-else), and that would be to awful.
* This day went more to thinking about how to solve something than actually do it

### Wednesday - *Birthday edition* 13/02/19
*Woooo birthday boi, halfway to 44*
* I still had a bug with adding nil instead of int to the db with #16. I solved it with creating nasty
if-else statements, I hope to refactor it later.
* I also commented a lot for the function so it's more clear how to use it :D
* Pushed #16 to master and started on #23. Need to fix homepage first tho :/
* I'm unsure of how to go further as the home page seem to just be empty on purpose and other stuff.
* I think I can finish tomorrow tho <3
* Nothing much happened today :/

### Thursday - 14/02/19
*I am 22 now, also the power went out home, so I went to the School of NTNU*
* Started on restoring home page, currently stuck in a bug where the post function in index.go won't start.
* I can now check if with a given course id, if it exists.
* I need to figure out how to confirm that the user joined a class. I want to start an js script from go :/
    * Solution: Global variable and send to template
* Can now add user to course through home page :D
* Course have an unique ID now, got inspiration form [here](https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html), 
I chose `github.com/rs/xid` because it was the length i wanted and quick and easy to use.
* Added logging for create/join course

### Friday - 15/02/19
* Added functionality for adding user to course through link when the user is logged in and when the user logs in.
* Added functionality for adding a new user to course through link.
* Added a test for RegisterGET handler :)
* Short day I guess, I'm going to write the report today, also -> no bugs today woo.

## Week Three
### Monday - 18/02/19
* \#23 Changed courseID back to int and auto increment so it's similar to the other tables. But i added a column for hash instead.
* \#20 Started on home page, but could only finish one of two tasks since assignments isn't done yet.
* We almost didn't have anything to talk about today with the [supervisor](https://www.ntnu.no/ansatte/ivar.farup).
* I'm going to start on [#22 - Admin FAQ Page](https://trello.com/c/0trVQS8x) now :D
* \#22 Looked at [this page](https://www.codeply.com/go/syFXJL6m5p/bootstrap-4-faq-accordion) for inspiration to the faq page, I liked the animations and stuff, but... If I use markdown instead
it would be way easier to just add a new faq in frontend and is over all less code and easier to implement. Agreed with [Project Owner](https://www.ntnu.no/ansatte/christopher.frantz)
to use md, it's also more consistence with this solution.
* Have to find out if we just store a hardcoded md file for faqs or make it possible to edit in the front-end by any teacher tomorrow.

### Tuesday - 19/02/19
* \#22 Decided on storing the md in db and let any teacher edit it, but also log every update.
* I chose to copy some of the design [Johan](https://github.com/JohanAanesen) used for course page to keep it more constance all over <3 
* I have some bugs on the faq site, but the main functionality is all done soon.
* I also moved db functions from shared/db to model and gave a temp fix to the extremely annoying go lint errors...
* All functions for faq is now done, each time a teacher updates the faq, the time is added **IN NORWEGIAN TIME**, this has to be written somewhere
as we talked about with the [Project Owner](https://www.ntnu.no/ansatte/christopher.frantz) yesterday.

### Wednesday - 20/02/19
* \#22 Tried to switch to a new package to get the text in the textarea editor to show, it still didn't work >:(
    * I think the problem is that id doesn't load again when the tab is shown.
    * It was bootstrap all along :'( 
    * I fixed the problem by removing the tabs and having the edit page on it's on page.
    * I also changed the design so it's more consistence to the admin dashboard.
    * Not really sure how to write tests for \#22 yet so I'll leave that for later.
    * Finished for now
* Started on [#21 - Admin Page Dynamic](https://trello.com/c/J8GQvTCt)
* Used some time looking at a pull request
* Updated the getCoursesToUser to get the courses sorted by the year and then semester in descending order.
It's nicer this way and more effective to have current classes at the top where they are easy reachable.
### ~~Thursday - 21/02/19~~ Ski day!
### Friday - 22/02/19
*Ski day was rough*
* Looked at pull request
* \#21 Merged main to branch
* Displaying courses and assignments in dashboard, course and assignment for admin now. I need to make assignments in sorted order by deadline
and display only active courses and assignments on dashboard, and then I'm done.

## Week Four - One month woo
### Monday - 25/02/19
* Continuing the work on \#21
* \#21 I created a new function in `model/assignment.go` called `GetAllToUserSorted()` which gets all the assignment to the user
sorted by the deadlines desc. 
* Course and assignment is displayed in correct order now
* Had some meetings, played some Minecraft, back now
* Made a pull request for [#21 - Admin Page Dynamic](https://trello.com/c/J8GQvTCt)
* Starting later on [#18 - Assignment Delivery page](https://trello.com/c/zyQpCo4K)


### Tuesday - 26/02/19
* Made an instance `BredeVM` on Openstack
    * Tried to do stuff there but I didn't understand shit
* Going back to [#18 - Assignment Delivery page](https://trello.com/c/zyQpCo4K) again
    * Started on front-end draft for delivery page
    * Stuck on how to get the forms value out of the sql statement \#bug
    * It was just stupid errors that shouldn't happened! >:( [Johan](https://github.com/JohanAanesen) helped me see them again <3
    * I'm going to die now, bye! <3
    * I'm back for another 1h, Made a shitty draft of the assignment delivery page, I think I can be done within two days actually.
    * Added countdown on delivery page, inspiration from https://www.w3schools.com/howto/howto_js_countdown.asp 

### Wednesday - 27/02/19
*I overslept 1h today*
* Continuing on [#18 - Assignment Delivery page](https://trello.com/c/zyQpCo4K)
    * Made a function for getting name and code from course
    * I used 2h one one pull request because I don't like Openstack...
    * And now I'm going to merge with master -_-'
    * I'm just trying to run the project now... shit day
    * Almost all day has gone to Openstack and trying to run the project local
    * [Johan](https://github.com/JohanAanesen) fixed docker in Intellij <3 can now run it local from intellij
    * Environment is up-ish again and I'm gone

### Thursday - 28/02/19
* After 1h, I can finally start coding again, since the env is completely up!
* Added functionality to get form values in an array
* Added functions for inserting and getting answers from db
* Merging with master
* Added test values in db, but it took some while since I didn't get errors about what was wrong :/


### Friday - 01/03/19
* Added test data to .sql file
* Users that has delivered can see what they have delivered.
* Users can update what they have delivered
* I made the design much better
* Fixed a bug with deadline
* Added real link to `see assignment` button
* Done with [#18 - Assignment Delivery page](https://trello.com/c/zyQpCo4K), thought I would use 21h, actually used 24h 17min,
this was mostly because I use ca. 7h all together to get the project starting.
* worked more on 18 do improve consistency and slicker design

### Sunday - 03/03/19
* started on [#32 - Admin Show User Submissions](https://trello.com/c/OR60Pi7l).
    * Can now show list of all students in course for each assignment
    * The teacher can see the submission outside admin/ because other students that are going to peer review it should 
    also see the submission, and this way I am saved from duplicate code :D
    * Merged with main
    * Checks if review user is verified now
    * Gets the links to the submissions user is going to review

## Week Five
### Monday - 04/03/19
* Had some meetings
* Started on [Alpha fix functionallity Req](https://trello.com/c/kv8MfMcv) with Svein and Johan

### Tuesday - 05/03/19
* Fixed joined course bug

### Wednesday - 06/03/19
* Course hash is now displayed to Admin
* It's only one link for joining course now
* Admin can copy the one join course link
* Changed name to Full name and made it unchangeable
* Switched to PRETTYTIME in faq and updated the faq 

## Thursday - 07/03/19 
* **Hafjell**

### Friday - 08/03/19
* Added confirm to user if admin wants to delete submission form or not.
* Updates submission_id in assignment even if it's nil now (That's good)
* Former submissions will now be deleted if submission form is changed

## Week Six
### Monday  - 11/03/19
_Merged following to master_ 
```
* Fixed join class bug
* Fixed submission/review bug
* Updated faq to be more professional
* Fixed it so it's only one join course link now
* When a user updates an assignment, the submitted time is updated too
* Also added TODO : time on places that used time, for later refactoring
```
* Starting on `Hardcode in norwegian timezone- alpha` and `Display PRETTYTIME in Norwegian format`
* Prettytime is fixed
* Go function for norwegian time is done
* Bug on js function for norwegian time

### Tuesday - 12/03/19
* deleted unnecessary .sql files
* fixed schedulerservice error, *I used an function from webservice before* 
* I changed some data to be better testdata in .sql
* Added two new functions
    * one for padding numbers with 0
    * and one for replacing [`toISOString()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString) so it doesn't alter the time before formatting.
* Also changed go and js `getTimeNorwegian()` to use locale `"Europe/Oslo"` 

### Wednesday - 13/03/19
* Removed timezone from sql and use golang instead, this is so it's easier to change time later
* Tried to remove go.mod/.sum but we needed them after all
* Fixing assignment card
* Added time zone on date formating
* ALso added envar for the timezone so it's easier to change later
* Changed to transaction some places in SQL queries
* Changed logtodb to return error instead of boolean 

### Thursday - 14/03/19
* I'm waiting for 3 pull request to be passed *wohoo*
* Starting on forgotten password
* Finished on forgotten password
* Not the best code, but it works for the alpha :D
* Also finished `dont'show-review-name`

### Friday 15/03/19
* Merged shit together

## Week seven
### Monday 18/03/19
* Fixed some alpha functionality

### Tuesday 19/03/19
* Fixed some alpha functionality

### Wednesday 20/03/19
* Fixed some alpha functionality

### Thursday 21/03/19
* Went to Hafjell

### Friday 22/03/19
* Checked for bugs
* Fixed dups email in register
* Fixing `stay-on-same-page-after-managed-student`

## Week Eight
### Monday 25/03/19
* Had some meetings
* Ran through the cycle of reviewing assignments

### Tuesday 26/03/19
* Fixed join course through hash error and enhanced it a little by searching for the hash in the given string

### Wednesday 27/03/19
* Sorted participant list by teacher and then name
* Looked at prettytime bug, it was an user error
* Updated go to 1.12.1, I have to fix some errors because of that now...
* Active assignments stay on home page until review deadline goes out
* Active assignments display `Submitted` or `Not submitted` for students only now
* Fixed deadline passed design on home and in course + the rest in admin without delivered status

### Thursday 28/03/19
* Peer Review is on every assignment card now, admins also
* Review deadline can not be before Assignment deadline
* Also fixed a bug in create assignment that removes submission/review form when an error occurred
* And smaller stuff. 

### Friday 29/03/19
* Fixed some bugs on pullrequest and merged to master

## Week Nine
### Monday 01/04/19
* Created a new gmail and started on mailservice
* Can already send an email since it's only 3 lines of code, too easy :/
* Function in mailservice for resetting password is ready
* Changed from https to ssh!!!

### Tuesday 02/04/19
* Added check to see if the email exists in the database
* Added feedback to the user
* Fixed `EmailExists` function
* Added mailservice to webservice, it works to get email from front-end now <3 :D 
* Added authentication to the mailservice, so that only the system can trigger mailservice
* Added new table in db for forgotten passwords OBS! TODO brede : ADD THIS TO PULLREQUEST!!!!!!!!!!!!!!!!!!!!!!! <3
```Mysql
CREATE TABLE `validation`
(
    `id`        int(11)     NOT NULL AUTO_INCREMENT,
    `hash`      varchar(64) NOT NULL,
    `user_id`   int(11)     NULL,
    `valid`     tinyint(1)  NOT NULL DEFAULT '1',
    `timestamp` datetime    NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users (`id`)
);
```
* Now adds to the table if the email is correct
* Implemented check for valid hash
* Changes password and added a new column in forgottenPass
* Also logs password has been changed
* Added some tests for the handlers
* Implemented two more functions, one for sending email to single recipient and one for multiple

### Wednesday 03/04/19
* Fixed forgotten password email message text
* Fixed `goFixShit.ps1` script and added colours <3
* Fixed golint errors
* Merged emailservice to master
* Changed name from `NTNU-Bachelor-Management-System-For-CS-Assignments` to `CSAMS` short for `Computer Science Assignment Management System` in project and repo
* Possible solution to `confirm email` issue [#97](https://github.com/JohanAanesen/CSAMS/issues/97) below
***
#### ~~forgotten_password~~ validation table
| id | hash | user | valid | time |
| ------ | ------ | ------ | ---- | ----- |
| 1 | jhdsgjshgfjsdgfsdfgusdgfyusdgfuysgeuyfesg |1| 1 |2019-03-24 14:42:00 |
| 2 | jhdsgjshgfjsdsfsddgfsdfgusdgfyusdgfuysgeu |32| 0 |2019-03-24 14:42:00 |
| 3 | jhdsgjshgfjfdgdfgfggusdgfyusdgfuysgeuyfes |11| 1 | 2019-03-24 14:42:00 |

#### pending_user table
| id | email | name | password | validationID|
| ----- | ----- | ----- | ----- | ---- |
| 1 | dummy@hotmail.com | Dummy dummy |adhjguhguygedwuyGYDGWEIUYG | 2 |
| 2 | YEET@mymeat.com | Svein Klausen | hjfsuyefgsuyegfuiygef | 3 |
| 3 | swag@yolo.com | Johan Nilsen |kjewgfhiuefwoig | 1 |

##### How to solve `confirm email` function
* Create user
* Create new entry in `validation`
* Add user to `pending_user` and link to `validation`
* Get email with one time use link (including `hash` from `validation`)
* Check if validation hash exists, not expired and is valid
* Get row from `pending_user` matching `validationID` to `validation` field `id`
* Move user from `user_pending` to `user` table
***

### Thursday 04/04/19
* Don't have time to work a full day today, but i have 6h to spear from mon-wed.
* I added repository and service for the table `users_pending` also struct in model
* Users can now create user through email confirmation


## Week Ten
### Tuesday 09/04/19
* Fixed some smaller stuff
* Added send email to whole course functionality
* Have to fix it more tomorrow, but it works!

### Wednesday 10/04/19
* Had some long meetings
* We have established the core functionality we need to be done now and added them to trello
* Added list of all student or private emails the email is going to be sent to

### Thursday 11/04/19
**I don't have too much time to work today, I will have to work some in the easter**
* Had to set up environment on my laptop
* Checked out an pullrequest and fixed merge conflicts when it was merged to master
* Created a new email `csams.noreply@gmail.com` and replaced the old with this
* Added count of how many student, private and total emails there is
* STarted on logs converting to service/repository

### Friday 12/04/19

### Saturday 13/04/19
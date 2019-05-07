# Meeting Notes 
<!-- Written by BredeFK - Moral supporter -->
## Thursday - 17/01/19
* We have the possibility to switch the report language to Norwegian anytime (*but we won't*)
* Make a good plan, so it's easier to know the progression!
* Weekly meetings with [Project Owner](https://www.ntnu.no/ansatte/christopher.frantz) at 12:15 and [Supervisor](https://www.ntnu.no/ansatte/ivar.farup) 13:15 on mondays.
* Find model (*we chose [Kanban](https://en.wikipedia.org/wiki/Kanban_(development\))*), and assign roles.
* Make Gantt scheme.
* Report and comment while working on project.
    * How/why did you solve the task like that?
* First week in June is presentation time!

## Monday - 28/01/19
* We presented the Gantt scheme to [Supervisor](https://www.ntnu.no/ansatte/ivar.farup), we needed to make it better.
* **Very important:** comment and write log while programming.
* Make a prototype for peer review and get good feedback, make before march pls.
* Milestones: can be seen on the Gantt scheme, also talk to [Supervisor](https://www.ntnu.no/ansatte/ivar.farup) after each milestone.
* Maybe make the project test driven?
* ~~Trello export comments~~
* Write why we chose Kanban in plan report.
* What license do we use? (*[GNU General Public License v3.0](https://github.com/JohanAanesen/CSAMS/blob/master/LICENSE)*)
* Add possibility to add a second email (*BredeFK did it 😎*)
* Add possibility to have stats (% of reviews done and grades stuff) and make it possible to export (json?).
* Have a checkbox for permission to share users email.
* Info for courses
* Save version build - ex for mobile development
* Log for delivery
* Log everything
* User can request to be removed.

## Tuesday - 05/02/19
**404**

## Monday - 11/02/19
* Make possible to upload course information for later use, have it in front-end.
* Add weight to assignments, but not max at 100, just fix it when all the weight is written and then calculate percentage and show to admin.
    * Also add option to show user the percentage
* Make it open for adding other tools than Go tools.
* Change `Priority` to `Order` in assignment creationX
* Make it so the user can write an description explaining why he/she rated X.
* Have an FAQ for admins
    * also add comment section for each assignment for the students to ask question to admin or other students.
* Look up Go modules.
* We are fast bois **gg**

## Monday - 18/02/19
* Assignment review 
    * Mandatory and voluntary questions
    * Talked about [#12 - Create Assignment form/page](https://trello.com/c/QpvcbVb6) for 25min
* Add deadline to assignments and gray out former assignments
* We informed [Project Owner](https://www.ntnu.no/ansatte/christopher.frantz) that it's [skiday](https://www.facebook.com/events/2070975702972081/) on thursday :)
* Document all smaller things that's vital for others later, ex. timezone in the sql file.

## Monday - 04/03/19
**Project Owner**
* Check url in back-end
* Sanitize input
* If a user has answerd submission form, don't let admin change the form
* Tell why stuff be like it be to the end-user
* Get hash from course and display to admin in front-end
* Give admin login/register link for joining course
* Add convert to MD plugin
* Set time uploaded on submissions admin
* Get the total number of submissions admin
* Maybe track view for submissions admin, to see whats new
* Remove joined course bug
* Fix submit new course btn
* Add order by course.semester back again >:(
* Make real name uneditable, maybe add username instead
* Forgot password functionality?

**Supervisor**
* Tabs on course -> more visible
* Maybe don't let admin submit assignment
* Add equation possibility in MD ([mathjax](https://www.mathjax.org/))

## Monday - 11/03/19(?)
* Deploy with different settings
* Demo of project
* Fix timezone
* Add which timezone it is
* Submission forms
* Remove name from reviewing
* Convert weight to percent
* Change from `update` to `re-upload`
* Change users password (admin feature)

## Thursday - 21/03/19
* Test Protocols - QA
    * Use cases
* Fix JS time bug
* Use Githubs issue tracker
* Implement function to override scheduler, maybe delay too
* Manage students
    * Persistence  
    
## Monday - 01/04/19
* Add buttons on submissions
* Edit review
* Show reviews done by user
* Don't implement auto-validation after all :)
* BUG: If user not in review table after deadline, it could go wrong on the assignment card.
* Confirm submission is delivered
* Only use one `!` after messages, instead of multiple
* Sc in review/submission (I don't remember what this was supposed to mean)
* Let admin change users name (We didn't implement this)
* Separate first and last name (Not this either)
* Maybe implement student number (nah)
* Create an noreply email
* Admin view, but user can't to admin stuff. Ex for student assistant
* Separate deadline. Add 5 min or something to deadline that is displayed
* Create new in assignment? (No idea what this is supposed to mean)
* Add tags for submission/review
* Marks for giving review and receive review

**Supervisor**
* Ask user before using data
    * GDPR
 * Wednesday 11.15 (?)

## Wednesday - 11/04/19
**Supervisor**
* ITICSE
* SIGCSE

**Project owner**
* Implement functionality to email all students in course
* Split first and last name
* BUG: if there is 0 peer review on assignment cards
* Fix update form to none
* Log everything. Filter by system, course, all, admin.
* Add functionality to have group
    * Create
    * Equal rights
    * Teacher can kick users

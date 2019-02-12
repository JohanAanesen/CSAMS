# The Log of Bread

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
    * Solution: Ask Johan for help.
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
* Fixed the hash bug, Johan saw that i 'fixed' it on the wrong variable so now the hash appeared too!
* Fixed the input validation bug, I had to switch from `onclick` to `onsubmit` that started the javascript script in the button-save. 
* Merged with master and got some problems, but Johan helped me fix it.
* Almost done with #8 now, only need some refactoring and I have to fix one test.

### Sunday - 10/02/19
*I'm 5h too short for this week, so I have to work some today too*
* Started up with setting up the environment on my laptop since I'm home and not in Gjøvik.
* I'm going to fix one test today and Johan is refactoring the code, so after that, I can make a pull request.
* Merged my branch to master

### Report for Week One
#### What I Have Done
In the first week I have worked on only one story/card, [#8 - User Page](https://trello.com/c/FZk85I6L). I had to learn about Bootstrap and in general 
web development before I could start properly. [#8 - User Page](https://trello.com/c/FZk85I6L) was about making the user profile. It had to
Show the users information as his name, student email, private email and his courses. I also had to implement a way for the user
to change information as his name, private email and password, but it also had to look nice (see Bootstrap). As most stories/cards 
I had to make tests for some functions also.

#### Design/Architecture Decisions
I got my inspiration from [this page](https://getbootstrap.com/docs/4.2/examples/offcanvas/#) and i Liked it because it had a 
"clean" look and the card design feels very modern. The page consists of two cards, one for the user information and one for the courses
the user is assigned to. This was to separate the user with the courses and it just overall looked better this way than
having all the information in one card. It's also easier for when the user needs to edit something to have all the editable 
information in one card. Most user information (everything except `NoOfCourses`) is displayed in an input field so it's easier
to change information. When teh user presses the edit button, the input fields is made editable by javascript. I also liked this design 
better so the information displayed does not change that much when the edit button is pressed.

#### What Went Good
* The user profile page ended up looking better than expected.
* I learned more about bootstrap and web development.
* I got to work with front-end, back-end and the database. In other words, I understand more of the 
projects design and are more prepared for the next card now. 

#### What Went Bad
* Loooots of bugs that ended up taking more time than it should. 
* I was also new to bootstrap and web development in general.
* I wrote messy code and didn't see it before Johan refactored it. 
* I overslept two times and stopped earlier three times, so I had to work on a sunday.

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
So I have to figure out how to send nil instead of anumber to the db.
* I know one way to fix the bug, but it's to much and messy code :/ (hella many if-else), and that would be to awful.
* This day went more to thinking about how to solve something than actually do it

### Wednesday - *Birthday edition* 13/02/19

### Thursday - 14/02/19

### Friday - 15/02/19


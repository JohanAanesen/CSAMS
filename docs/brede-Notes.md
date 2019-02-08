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

## Week Two
### Monday - 11/02/19

### Tuesday - 12/02/19

### Wednesday - *Birthday edition* 13/02/19

### Thursday - 14/02/19

### Friday - 15/02/19


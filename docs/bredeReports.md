# Report of Bread
<!-- By BredeFK -->
## Week One
### What I Have Done
In the first week I have worked on only one story/card, [#8 - User Page](https://trello.com/c/FZk85I6L). I had to learn about Bootstrap and in general 
web development before I could start properly. [#8 - User Page](https://trello.com/c/FZk85I6L) was about making the user profile. It had to
Show the users information as his name, student email, private email and his courses. I also had to implement a way for the user
to change information as his name, private email and password, but it also had to look nice (see Bootstrap). As most stories/cards 
I had to make tests for some functions also.

### Design/Architecture Decisions
I got my inspiration from [this page](https://getbootstrap.com/docs/4.2/examples/offcanvas/#) and i Liked it because it had a 
"clean" look and the card design feels very modern. The page consists of two cards, one for the user information and one for the courses
the user is assigned to. This was to separate the user with the courses and it just overall looked better this way than
having all the information in one card. It's also easier for when the user needs to edit something to have all the editable 
information in one card. Most user information (everything except `NoOfCourses`) is displayed in an input field so it's easier
to change information. When teh user presses the edit button, the input fields is made editable by javascript. I also liked this design 
better so the information displayed does not change that much when the edit button is pressed.

### What Went Good
* The user profile page ended up looking better than expected.
* I learned more about bootstrap and web development.
* I got to work with front-end, back-end and the database. In other words, I understand more of the 
projects design and are more prepared for the next card now. 

### What Went Bad
* Loooots of bugs that ended up taking more time than it should. 
* I was also new to bootstrap and web development in general.
* I wrote messy code and didn't see it before Johan refactored it. 
* I overslept two times and stopped earlier three times, so I had to work on a sunday.

## Week Two
### What I Have Done
* This week, I have worked on [#16 - Logging to database](https://trello.com/c/CwIxfhpk) and #[#23 - Join Class Functionality](https://trello.com/c/lGOGylxO).
* I also made a powershell script for go fmt, vet, lint, cyclo and test. *It went faster this way*
* With the logging I logged the stuff the user does to the database in an own table.
* Join class was quick to fix as soon as I understood how to do it and it wasn't so much thinking about how to do it either.

### Design/Architecture Decisions
* I choose the unique id function for the courses from this site [here](https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html), 
and chose `github.com/rs/xid` because it was the correct length and the least hassle.
* I had every log in one table with many nulls instead of one table to each type to have it more oversikkelig. 
<!-- TODO Write more here -->

### What Went Good
* Way less bugs this week, and overall higher productivity for each card!
<!-- TODO Write more here -->

### What Went Bad
* Too few bugs I'm too good now :/
<!-- TODO Write more here -->

## Week Three
### What I Have Done
* This week I have worked on \#23, \#20, \#22 and \#21 in that order.

### Design/Architecture Decisions

### What Went Good

### What Went Bad

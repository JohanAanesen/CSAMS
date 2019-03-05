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
* I had every log in one table with many nulls instead of one table to each type to have it more 'oversikkelig'. 


### What Went Good
* Way less bugs this week, and overall higher productivity for each card!
* I had more than one card!

### What Went Bad
* Too few bugs I'm too good now :/
* Maybe too few cards I guess

## Week Three
### What I Have Done
* This week I have worked on [#23 - Join Class Functionality](https://trello.com/c/lGOGylxO), [#20 - Front Page Dynamic](https://trello.com/c/AxyDWjuP), 
[#22 - Admin FAQ Page](https://trello.com/c/0trVQS8x) and [#21 - Admin Page Dynamic](https://trello.com/c/J8GQvTCt) in that order.
* First I had to fix a bug on \#23 and change the primary key to course: courseid back to int + auto_increment and add another column: hash to have in the link.
* I then started on \#20, but couldn't finish the formerly two tasks because assignments wasn't done yet then.
* I then started looking at \#22 and was a bit unsure how to solve it, but after talking to the [Project Owner](https://www.ntnu.no/ansatte/christopher.frantz)
we decided to go with an editable markdown page.
* After that I started on \#21 and I'm still working on it, I have completed task one which is dynamically list courses sorted by year/semester. So almost done!

### Design/Architecture Decisions
* \#23, I changed back to int as ID since it was the same as the other tables, also more consistent this way. It's also usual to have hash and id separate.
* \#20, was just straight forward, it was just to display courses and I did.
* \#22, This was not straight forward at all! <!-- ¯\_(ツ)_/¯ --> 
    * We first talked about having a hardcoded file in the project, like a json file with the questions and answers or something.
    * Then I talked to the [Project Owner](https://www.ntnu.no/ansatte/christopher.frantz), and he recommended using markdown since we already had that in another place.
    * And finally we ended up using markdown that every teacher can edit on the front-end, but also tracking the changes in the log table in the db, in case anyone fucks up :)
        * This was a very simple solution and more consistent all over the page, since we have md in assignment and course page :D
* \#21 No big design/architecture here, the courses and assignment are displayed in cards to make the more modern and easier see they are separate.


### What Went Good
* I did more cards this week, I completed three cards and started on one more :D 
* I work faster now, It's also less bugs
* We agreed on how to display the FAQ and it was a simple solution but also kinda nice.
* Had a day off, but could complete at least 30h anyway <3 

### What Went Bad
* I couldn't complete \#20 :/
* I didn't have time to finish my last card (\#21) this week :/
* Used some time talking about how to solve the faq solution, (but that could also be a good thing since we discussed it before implementing )
* I wasn't that productive this week (see: hangover after ski day)
* Used some time with \#22 because of various bugs, found out it was bootstrap that fucked me over at the end thx to Johan <3
* I din't have birthday this week :/

## Week Four
### What I Have Done
Worked on \#21, \#18 and \#32.
Also deployed VM on openstack, that took some time...
Finished with \#21 on monday, started on \#18 on tuesday - friday (24h).
On saturday, I started on \#32 and finished it at 2-3 night on monday.

### Design/Architecture Decisions
* \#21 I chose to sort the assignments by earliest deadline and course by year/semester desc.
This is so you can see the earliest assignment the user is going to deliver and with the courses, the
most relevant course the user is taking that year.
* \#18 I added a countdown on the assignment delivery page so the user can see more properly
how much time is left until the deadline. 
* \#32 I have to fix some stuff from here next week. Admin can see everybody in assignment, even if student hasn't delivered.
This was because its a alpha prototype. Student and teacher share the same page for when they are reviewing an assignment
this was to save me from duplicate code.
### What Went Good
* Got a lot done because we had a prototype deadline on monday.
* Actually liked the design in my code now
most stuff worked

### What Went Bad
* Docker and Openstack slowed me down by 7ish h in total this week, since I couldn't get it working
* Had some really easy bugs, but couldn't find them before after 1-2hours...

Week 1
===
This first week I have worked on:
* '#2 Initial database tables'
* '#3 Go Router / Web Server'
* '#14 Database connection class'
* '#5 Login/Register page'
* '#11 Create Course form/page'

As well as a major refactor on '#8  User page'

Sessions - reflection
---
I was working on the Login/Register page, but somewhere along the process Svein reminded me we needed 
a session management system of sorts. I already had an idea in the back of my head to just generate
a cookie with a random value, store it in a database table and check that everytime. Though this
approach would include a lot more I/O operations than needed on the poor database. After some rounds 
with Google Search I figured that the go-to approach was to use some sort of 'web toolkit'. For Go 
there is a lot of them, but Gorilla Web Toolkit kind of stood out because it's apparently the simplest
of the lot, aka it has more bare-bone functionality and nothing really special. The toolkit consists of
several packages and you import the ones you need. I essentially started with the Sessions package, as
this was something I needed, but it needed the 'mux router' package as well. The mux router is very 
similar to the included router functionality in net/http, but it has some more functionality and is
very easy to use. It also seems like it is the router package most people prefer on StackOverflow.
I also added the Gorilla SecureCookies package to encrypt the data in the cookies.

I like the sessions system in place now a lot, works like it should and it is very efficient to have
user data stored in the session variables.

The good, the bad, the rest
---
I also worked on the Create Course part of the admin panel, which Svein already did most of the front
end on. The only part we really had a discussion on here was wether or not to have static link fields
so the teacher could ass links, or set the description field as a markdown editor, and let the teacher
add as many links as he wanted there. We chatted a little bit with the Product Owner who wondered about
his semantic approach to this, although we ended up selling him the markdown version and will discuss 
this further on Monday 11/2.

Positives:
+ Gotten quite far in one week
+ Using 3. party packages in Go is easy, fun and beneficiary
+ Brede learned Bootstrap!
+ Sufficient tasks for one week
+ Everybody had something to work on the whole time

Negatives:
- Had some problems mocking a session in testing, not obvious in documentation this could be done
- Sometimes I would notice a lot of duplicate code through out the functions, and i did nothing with it
    - Fixed this in the major refactor i did on #8
- Tasks could probably have been better defined from my end as the project leader

Week 2
===
I have worked on
* Refactoring the project
* '#7 - Course Page'

Decisions
---
I made the Course Page where I decided on a standard container width, with tabbed information pages. One tab for the
course description, one for assignment links, one for participant list and one for questions. I am waiting on
Christopher if I should implement the questions tab or not :/

The good, the bad, the rest
---
I felt like I didn't get much done this week, the refactoring took a lot longer than I anticipated, and a lot of time
was wasted just waiting for the refactoring to get anywhere because Svein had the main idea behind how the structure 
should become. When he couldn't work too much the monday we ended up spending the whole tuesday aswell. Then I had birthday
and some other stuff came up, Here's to hoping next week we're getting loads done :cheers:

Positives:
+ New project structure is more tidy
+ Course page looks nice
+ Found a MD to HTML package that is perfect for our usage
+ Got time to do a lot of research

Negatives:
- Didn't get a lot done
- Refactoring took a lot more time than anticipated

Week 3
===
I have worked on
* '#10 Participant list'
* '#28 Peer Review Service'
* Dockerizing the application

Decisions
---
In terms of decisions I have been making, then what I have been thinking the most about is whether or not to make this 
application microservice oriented, because in a lot of different cases, it might not be the best idea to do. Let's
take the Peer Service as an example, it doesn't really take much computing capacity to do, even if it where doing thousands
of submissions at once. But again we want to separate functionality that isn't directly included in a webservice from the webservice.
This is because at any given deadline, that will be the time with the most traffic at the site, and if the webservice also will start
calculating who is reviewing who at that given deadline, then the service might fail to service users. This is why we want to branch
the peer service into it's own container/service so it won't affect the webservice in terms of hogging cpu time/resources. 

Another aspect is that the application will have 2 main components to the delivery system, and those are the peer review part, and
the auto validation part. And the auto validation really need to be its own service because it will be really resource heavy.

The good, the bad, the rest
---

Positives:
+ Peer Review Service fun to make and very functional
+ List I made looks good
+ Docker is fantastic when it works
+ Learning a lot around the functionality of microservices and how to create services

Negatives:
- Struggled with learning and understanding Docker, mostly getting it to work properly with out special use case
- git thought i deleted 177 files and then created 177 files so i got like 7000 lines deleted and added in 2 fast commits ay
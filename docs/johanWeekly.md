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


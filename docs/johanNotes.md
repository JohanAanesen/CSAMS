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
* 
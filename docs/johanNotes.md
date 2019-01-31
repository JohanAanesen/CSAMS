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
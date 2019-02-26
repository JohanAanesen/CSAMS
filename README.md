# NTNU-Bachelor-Management-System-For-CS-Assignments

## Links
* [Github](https://https://github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments)
* [Trello](https://trello.com/bachelor531)
* [ShareLatex](https://www.overleaf.com/project/5c3491a162ba3128fda8c11d)
* [Discord](https://discord.gg/rZ4zg2R)
* [Google Drive](https://drive.google.com/drive/folders/1kiQiBj12zrn45q6QOfXefrzgNb4fZhyW?usp=sharing)
* [Toggl](https://toggl.com)
* [UML](https://www.lucidchart.com/invitations/accept/421b3f38-581e-4790-80f7-3d43604a717c)

## Members
| Role | Name | E-mail |
| -------- | -------- | ------- |
| Project Leader | Johan Aanesen | johanaan@stud.ntnu.no |
| Member | Brede Fritjof Klausen | bredefk@stud.ntnu.no |
| Member | Svein Are Danielsen | sveiad@stud.ntnu.no |

## Project Owner
* [Christopher Frantz](https://www.ntnu.no/ansatte/christopher.frantz)

## Supervisor
* [Ivar Farup](https://www.ntnu.no/ansatte/ivar.farup)

## OBS!
This project is in the norwegian timezone! At least the:
* [`config/database.sql`](https://github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/blob/b74344a4a1673c4473442db99f965c17643d83c1/config/database.sql#L19) file at line `19`: `SET time_zone = "+01:00";`
* [`model/adminfaq.go`](https://github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/blob/d4a2c5d1bce5a6f22c6d179a80351f1980fc55e6/model/adminfaq.go#L50) file at line `50`: `date := time.Now().UTC().Add(time.Hour)`
* [`template/assignment/upload.tmpl`](https://github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/blob/18-Assignment-Delivery-Page/template/assignment/upload.tmpl) <!-- TODO change link when in master --> file at line `39`: `let now = new Date().getTime() + (1000 * 60 * 60);`

## Commit messages guideline
Start with the numger `#42` where the number is the number to the `card` you're working on Trello followed by:
* update	- updated dunction/file
* add	- added function/file
* fix	- fixed a bug
* remove	- removed a function/file
* test	- added a test for a function
* refactor - refactored the codestructure, moved a file or changed the visual representation of code.

Aka #CARDNR ACTION BRIEF_EXPLANATION

Examples:
* "#2 add database.sql initial table file"
* "#8 fix changing password bug"
* "#14 update weekly report for week 2"

## Internal Peer Review Checklist
- [ ] Compiles
- [ ] Testing is sufficient and passing
- [ ] Branch completes the assigned tasks (from trello card)
- [ ] Code is well commented and tidy
- [ ] Is there something that could be done better? propose update
- [ ] Leave feedback as comment on the pull request :)

## Go Module
* See [quick start](https://github.com/golang/go/wiki/Modules#quick-start) for Modules

## Go Get's
* go get -u github.com/go-sql-driver/mysql
* go get github.com/gorilla/sessions
* go get -u github.com/gorilla/mux
* go get github.com/gorilla/handlers
* go get github.com/gorilla/securecookie
* go get golang.org/x/crypto/bcrypt
* go get -u github.com/go-chi/chi
* go get -u github.com/shurcooL/github_flavored_markdown
* go get github.com/rs/xid

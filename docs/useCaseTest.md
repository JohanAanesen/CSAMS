# Use Cases For Testing Manually
## User
* Create user
    * Email exists
    * Passwords doesn't match
* Log in user
    * Wrong email
    * Wrong password
    * Wrong email and password
    * Forgotten password (only for students)

## Teacher
1. Startup
    1. Create Course
    2. Create Assignment
    3. Create Submission Form
    4. Create Review Form
    5. Copy link and send to students
    
2. Change startup
    1. Course
        1. Delete course(?)
        2. Edit course
        3. See assignment to course
        4. See hash and ID
        5. Correct `join-course-link`
    2. Assignment
        1. Delete assignment(?)
        2. Edit assignment
        3. See submissions
        4. See details
        5. See deadline and ID
    3. Submission Form
        1. Delete submission form
        2. Add new field
        3. Enable/disable weights
        4. Rearrange fields
        5. Delete field
        6. Edit field
            7. Comment/no comment
    4. Review Form
            1. Delete submission form
            2. Add new field
            3. Enable/disable weights
            4. Rearrange fields
            5. Delete field
            6. Edit field
                7. Comment/no comment
3. Scheduler
    1.  Review students group is created after assignment deadline
4. Manage Student
    1. Change password to student
        1. Cancel change password
        2. Confirm Change password
    2. Remove student from course
        1. Cancel remove student
        2. Confirm remove student
5. FAQ
    1. Faq exists
    2. Can edit faq in md
    3. Correct new date and everything works



## Student
* Join course
    * Through link while logged in
    * Through link while logging in
    * Through link new user
    * Through logged in with hash in `join course` bar
* Deliver assignment
    * No submission form
    * Before deadline
    * After deadline
* Review assignment(s)
    * After deadline the review pair should be there
    * Click on on review and submit
* See reviews on own submission after deadline

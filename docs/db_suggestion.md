# Database Design Suggestion

## Tables

### Users
Name | Type | Explained
--- | --- | ---
ID | int | Primary Key
user_login | text | Username
user_password | text | Password
user_email | text | School email
user_registered | datetime | Datetime on account creation

### User Meta
Name | Type | Explained
--- | --- | ---
meta_id | int | Primary Key
user_ID | int | Foreign Key (Users.ID)
meta_key | text | Key
meta_value | text | Value

#### Meta Keys & Values
Key | Values | Explained
--- | --- | ---
user_role | ["teacher", "student"] | Authentication level
user_email_forward | email | Email used if user want to forward emails
user_phone | number | Phone number

### Courses
Name | Type | Explained
--- | --- | ---
ID | int | Primary Key
course_code | text | Course Code
course_name | text | Course Name
course_description | text | Course Description

### Course Meta
Name | Type | Explained
--- | --- | ---
meta_id | int | Primary Key
user_ID | int | Foreign Key (Course.ID)
meta_key | text | Key
meta_value | text | Value

#### Meta Keys & Values
Key | Values | Explained
--- | --- | ---
course_year | number | What year is the course
course_semester | ["fall", "spring"] | What semester is the course


### Assignments
Name | Type | Explained
--- | --- | ---
ID | int | Primary Key
assignment_name | text | Name of assignment
assignment_description | text | Description for the assignment
assignment_course_ID | int | Foreign Key (Course.ID) 

### Assignment Meta
Name | Type | Explained
--- | --- | ---
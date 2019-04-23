# Log functions checklist
## Implemented and works
**Users**
- [X] NewUser           
- [X] ChangeEmail         
- [X] ChangePassword    
- [X] ChangePasswordEmail
- [X] CreateSubmission 
- [X] UpdateSubmission
- [X] DeleteSubmission   
- [X] FinishedOnePeerReview 
- [X] UpdateOnePeerReview   
- [X] JoinedCourse

**Admin**
- [X] AdminCreateAssignment  
- [X] AdminUpdateAssignment  
- [X] AdminCreateSubmissionForm 
- [X] AdminUpdateSubmissionForm 
- [X] AdminDeleteSubmissionForm
- [X] AdminCreateReviewForm
- [X] AdminUpdateReviewForm
- [X] AdminDeleteReviewForm  
- [X] AdminCreatedCourse 
- [X] AdminUpdateCourse
- [X] AdminEmailCourseStudents
- [X] AdminRemoveUserFromCourse
- [X] AdminChangeStudentPassword  
- [X] AdminCreateSubmissionForUser 
- [X] AdminUpdateSubmissionForUser 
- [X] AdminDeleteSubmissionForUser

**Bug**
- [ ] AdminUpdateFAQ : `http: panic serving [::1]:8348: runtime error: invalid memory address or nil pointer dereference`


**No functions for yet:**
- [ ] AdminCreateFAQ

**No triggers for yet**
- [ ] LeftCourse 
- [ ] AdminDeleteAssignment
- [ ] AdminDeleteCourse
# go-dutch
Expense sharing application
Functional requirements:
1. A user can create an account with a username, email and password.
2. A user can login with his username and password.
3. Only a logged in user can a create a group with/without registered members.
4. Only the creator of a group can add/remove members from that group.
5. A user can add/remove/update only their expenses in their groups.
6. A user should be able to fetch only his/her groups.
7. A user should be able to see only his/her expenses in their groups.
8. A user should be able to see only his/her total dues with other members in a group.

API design:
Create user /users
Login user /users/login

Create group /groups
Get groups by user ID /groups/:user_id

Add group member /groups/:group_id/members
Get group members /groups/:group_id/members

Add group expense /groups/:group_id/expenses
Get group expenses /groups/:group_id/expenses

CREATE GROUP FLOW:
    - check if user is logged in
    - create group
    - check if member names are sent
    - get user ID(s) by user name(s)
    - add members

ADD GROUP MEMBER FLOW:
    - check if user is logged in
    - check if group exists
    - check if user is creator of group
    - get used ID(s) by user name(s)
    - add member

ADD EXPENSE FLOW:
    - check if user is logged in
    - check if group exists
    - check if user is member of that group
    - add expense
    - update balance
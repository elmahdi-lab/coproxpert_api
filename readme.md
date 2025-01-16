Snippets:
```bash
go get -u ./... 
go mod tidy

atlas migrate diff --env gorm 
atlas migrate apply --env gorm --url "postgres://postgres:postgres@localhost:5432/coproxpert_db?sslmode=disable&search_path=public" --revisions-schema true
atlas migrate validate --env gorm 
```

authentication is already set by checking each request by the auth.go middleware.
Levels of Permission: Admin, Manager, User.
Separate admin/managers from users

units : only users but any admin or manager in the org can manage.
unit group : only org manager or admin can manage.
org : only org admin (create unit groups...)

super admin anything.

V1 Core Features:
- X User Authentication
- X Organization, Building and properties (or units)
-  Common Spaces (pool, gym, garden ...)
-  Complaints
-  Document sharing
-  Notifications
- X Maintenance schedule
- X Resolutions and votes
-  Admin Subscriptions with licenses

V2 Features:
- Chat
- Invoice reminders
- Polls
- Surveys
- Events
- Calendar

V3 Features:
- Payments
- Budgets
- Reports
- Analytics
- Integrations

V4 Features:
- AI

## INFRA & DEVTOOLS
Neon DB, MailTrap

HTZ 2x Compute Instances for API & DB
GCP instance will host insta karma
GCP Functions // TBD
GCP PubSub // TBD
GCP Cloud Logs // Yes
GCP Secrets // Probably
AWS SES Emails // Using EC2 small instance for sending emails
Push Notifications TBD
Hashicorp Vault


https://github.com/zincsearch/zincsearch
https://github.com/gofiber/recipes?tab=readme-ov-file

report, invoice, payment, assembly, budget, chat

"CoProXpert, pour une gestion de qualité, en toute simplicité!"

User:
Join with invitation code.
Login, manage own profile.
View unit, view unit group. view resolution, vote, view documents, view budgets.
View Inspections, view maintenance schedule, view own complaints. view reports.
Create a complaint.


# User Actions
id, user_id, action, is_viewed, created_at, updated_at
action types: login, sign up, create something, receive something
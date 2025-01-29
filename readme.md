Snippets:
```bash
go get -u ./... 
go mod tidy

atlas migrate diff --env gorm 
atlas migrate apply --env gorm --url "postgres://postgres:postgres@localhost:5432/coproxpert_db?sslmode=disable&search_path=public" --revisions-schema true
atlas migrate validate --env gorm 
```

Levels of Permission: Admin, Manager, User.
Separate admin/managers from users

units : only users but any admin or manager in the org can manage.
unit group : only org manager or admin can manage.

super admin anything.

V1 Core Features:
- X User Authentication
- X , Building and properties (or units)
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
GCP Functions // YES
GCP PubSub // YES
GCP Cloud Logs // Yes

GCP Secrets // YES
AWS DynamoDB // YES
AWS SES Emails // with AWS LAMBDA

Push Notifications TBD Probably AWS SNS
Hashicorp Vault / GCP Secret Manager


https://github.com/zincsearch/zincsearch
https://github.com/gofiber/recipes?tab=readme-ov-file

report, invoice, payment, assembly, budget, chat

"CoProXpert, pour une gestion de qualité, en toute simplicité!"
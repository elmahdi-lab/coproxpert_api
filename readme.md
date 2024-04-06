authentication is already set by checking each request by the auth.go middleware.
Levels of Permission: Admin, Manager, User.
Separate admin/managers from users

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

SW HTZ 2x Compute Instances for API & DB
GCP Compute for a MicroService
GCP Functions
GCP PubSub
GCP Cloud Logs
GCP Secrets
AWS SMS
AWS SES Emails / Using EC2 small instance for sending emails
Push Notifications TBD





https://supertokens.com/docs/thirdpartyemailpassword/introduction
https://github.com/zincsearch/zincsearch
https://instatus.com/
https://github.com/gofiber/recipes?tab=readme-ov-file

report, invoice, payment, assembly, budget, chat, 
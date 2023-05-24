# User

1. POST /authenticate {username, password, remember_me: true} : token
2. POST /authenticate with social media.
3. POST /password_forget {username}
4. GET /disconnect

# Organization

1. POST /organization/add {organization} : organization
2. PATCH /organization/id {organization} : organization
3. PUT /organization/id {organization} : organization
4. DELETE /organization/id : bool

# Community

1. POST, PATCH, PUT, /community
2. DELETE /community

# Residence

1. POST, PATCH, PUT /residence
2. DELETE /residence

# Files

1. POST /files/add
2. DELETE /files/id

# Activity

1. POST, PATCH, PUT, DELETE /activity

# Events

1. POST, PATCH, PUT, DELETE /events

# Maintenance

1. POST, PATCH, PUT, DELETE /maintenance

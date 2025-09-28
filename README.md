# ginie

Web Application powered by Gin

## Features
### Backend
- Blog
    - Post
    - Category
    - Tag
- Administration Panel
- User Management (Register\Login)
    - Roles:
        - Admin (has access to dashboard)
        - User (has access to some features but not to the dashboard)
- Localization
    - Two languages for now (English and Persian)
- Websocket Server
    - Publishes realtime statistics about the OS/System

### Frontend
- Scroll-to-top
- Dark-mode Switcher
- Minimal Data-table component (using MaterializeCSS table and API
  pagination/filtering instead of 3rd-party plugins/libraries

## Dependencies
### Backend
- **Framework:** Gin
- **Templating:** `html/template`
- **Database:** SQLite, GORM

### Frontend
- **UI:** Materialize CSS, Alpine.js


## TODO

- [ ] Restructure the project into a clean architecture

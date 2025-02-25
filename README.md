# Tudu - Task Management Application

A full-stack task management application built with:

- **Frontend**: React (Next.js) - planned
- **Task API**: Go
- **Authentication API**: Python (Flask)
- **Database**: PostgreSQL

## Project Overview

Tudu is a task management application designed to demonstrate a microservices architecture using different programming languages. The project serves as a learning platform for Go development while strengthening Python full-stack skills.

## Project Structure 
tudu/
├── services/
│ ├── task-service/ # Go API service for task management
│ └── auth-service/ # Python Flask service for authentication
├── database/ # Database schemas and migrations
└── docker-compose.yml # Docker configuration for local development

## Technology Stack

- **Task Service**: Go - handles task CRUD operations
- **Auth Service**: Python/Flask - manages user authentication and authorization
- **Database**: PostgreSQL - stores user and task data
- **Frontend**: React (Next.js) - planned for future implementation
- **Containerization**: Docker - ensures consistent development and deployment environments

## Development Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/tudu.git
   cd tudu
   ```

2. Start the services using Docker Compose:
   ```
   docker-compose up
   ```

3. Access the services:
   - Task API: http://localhost:8080
   - Auth API: http://localhost:5000
   - PostgreSQL: localhost:5432

## Current Status

- Basic Docker infrastructure is set up
- Minimal Go and Python services are running
- PostgreSQL database is configured with initial schema
- Frontend implementation is planned for future development

## Next Steps

- Implement complete task management API in Go
- Develop authentication flows in the Python service
- Create React frontend with user authentication
- Add testing and CI/CD pipeline

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
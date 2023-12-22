## Project Description

A Go-powered "Event Booking" REST API

## Routes

| Method | Route                 | Description                 | Auth Required |
| ------ | --------------------- | --------------------------- | ------------- |
| GET    | /events               | Get all available events    | No            |
| GET    | /events/{id}          | Get an event                | No            |
| POST   | /events               | Create a new bookable event | Yes           |
| PUT    | /events/{id}          | Update an event             | Yes           |
| DELETE | /events/{id}          | Delete an event             | Yes           |
| POST   | /signup               | Create a new user           | No            |
| POST   | /login                | Login a user                | No            |
| POST   | /events/{id}/register | Register user for event     | Yes           |
| DELETE | /events/{id}/register | Unregister user from event  | Yes           |

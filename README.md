# Bureaucr.at

This repository holds a simple management system for the employees of Bureaucr.at.
The management system is available by means of a simple REST API.

## Scenarios

### Add a New Employee

This use case refers to adding new employees. A new employee can be added by providing
the name of their direct manager.

### Remove an Employee

This use case refers to removing employees. All the employees managed by the employee
to be removed can be transferred to a different manager, if a name for such a manager is
provided.

### Find Common Manager

This use case refers to finding the least senior manager that is common to two given
employees.

### Get Hierarchy

This use case refers to presenting the whole hierarchy in a human-readable form.

## Process View

### Interface

To implement the above mentioned scenarios, a few HTTP REST APIs are exposed.

| Path       |  Method  | Description                                      |
| ---------- | :------: | ------------------------------------------------ |
| /add       |   POST   | Add a new employee                               |
| /remove    |   POST   | Remove an employee                               |
| /manager   |   GET    | Get least senior common manager of two employees |
| /hierarchy |   GET    | Get whole hierarchy                              |

#### POST /add

##### Body Format

##### curl Example

```bash
$ curl # TODO
```

#### POST /remove

##### Body Format

##### curl Example

```bash
$ curl # TODO
```

#### GET /manager

##### Path Format

##### curl Example

```bash
$ curl # TODO
```

#### GET /hierarchy

##### curl Example

```bash
$ curl # TODO
```

### Storage

TODO

## Logical View

TODO: class diagram

## Development View

### Implementation Constraints

* the employee directory is an in-memory structure
* in the hierarchy, nodes point to children and not vice-versa
* implementation in golang

### Implementation Assumptions

* if employee E0 manages employee E1, their least senior common manager is the employee immediately above E0, if they exist, otherwise E0

### Building

```bash
$ # TODO
```

### Unit Testing

```bash
$ # TODO
```

### Integration Testing

```bash
$ # TODO
```

### Performance Testing

```bash
$ # TODO
```

### Kubernetes deployment

```bash
$ # TODO
```

### Further Improvements

TODO

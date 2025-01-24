# TemplateGolangServer

This template golang server app, will serves as a template for us to use as base in all project developed in the Midnight_8k livestream.

## Instalation

### Clone the template project
```
git clone <TemplateGolangServer>
```

### Download and install the dependencies golang
Ensure that you open the same folder where you clone the project.

```
go mod tidy
```

## REST API Demo Template Contract
This demo application contains two domains, the User domain and the Message domain.

On Message domain we are able to create update and delete a message from the database.

## Database
Currently we are using the SQLite as database engine. But we've used GORM as obeject relational mapping, that means that we can upgrade this database to another database engine sutch as Postgre SQL, MYSql or Oracle SQL.

## Tests
We don't have any test at the moment. That an future endeveour.
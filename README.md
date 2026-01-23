# GATOR: A CLI RSS Feed tool

Gator is a multi-user tool to follow and subscribe to RSS feeds all from within your
terminal.

## Requirements

1. Postgresql: Need to install postgresql. This is used to locally store the 
feeds and their posts. In a Debian-based system, the following could be used:

```
sudo apt update
sudo apt install postgresql postgresql-contrib

```

2. Config: Gator requires a JSON config file with a couple of properties:

```
{
    "db_url":"postgres://<username>:<password>@<host>:<port>/gator?sslmode=disable",
    "current_user_name":"<a user name>"
} 

```

# url-shortener

#### *In Development*
 

## Database Setup

Postgres DB setup

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start postgres for you.

### Create Your Databases

Once you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.


### *Setup Script Coming Soon*

## References

[http://gobuffalo.io](http://gobuffalo.io) 


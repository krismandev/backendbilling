#### Go Basic API billingdashboard

##### Installation
- Delete go.mod & go.sum if exists
- By default, the appname is "billingdashboard", if you want to rename the appname, please replace all "billingdashboard" text in all files with your appname
- Run "go mod init {appname}" command (default appname = billingdashboard). It will return error if the appname is not same as declared in all imported files 
- Setup environtment for database, redis, or JWT in "config/config.yml"

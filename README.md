#### Go Basic API backendbillingdashboard

##### Installation
- Delete go.mod & go.sum if exists
- By default, the appname is "backendbillingdashboard", if you want to rename the appname, please replace all "backendbillingdashboard" text in all files with your appname
- Run "go mod init {appname}" command (default appname = backendbillingdashboard). It will return error if the appname is not same as declared in all imported files 
- Setup environtment for database, redis, or JWT in "config/config.yml"


# Invoice Type
If yout want  new format of invoice or new source data of invoice, add new record to invoice type table.


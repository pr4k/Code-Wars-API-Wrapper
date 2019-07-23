# About this project

This project is created to periodically update code wars solution repo by directly scraping user solved solution and pushing it to the repo all just by a single click.
All it needs is your code wars Api secret for fetching user info and kata info using the API given by code wars.
You can get your API Secret by going into `Account setting ` from the code wars dashboard.
It also needs username/email and password of code wars account and github account to scrap solution from codewars and push it to 
the repo. We need to scrap the solution because code wars API has no end point to fetch user solved solutions.

# How To Setup
### Required credentials

* Code Wars API Secret
* Login credentials for code wars
* Login credentials for Github

Go to the  `return_auth.go` and enter your credentials to the indicated variables .

## To use this script to auto update a Github Repo
Change the repo path in `main.go` to the local git repo where you want to push the solution.
Done thats all ;)

# How to Run

You can either create a binary by using `go build ` or simply run by cloning the repo and using `go run main.go return_auth.go`.
It will automatically fetch the solved kata which can be later used to fetch solution by kata ID .

#### Created By ~Prakhar

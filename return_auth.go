package main

import(
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"fmt"
	"os/exec"
)


func returnAuth() string{
	apiSecret:="Your_API_Secret"
	return apiSecret
}
func returnIdPassword()(string,string){
	username:="Your_email_id"
	password:="Password_at_codewars"
	return username,password
}
func returnGitRepo()(string,string){
	username:="Github_username"
	password:="Github_Password"
	return username,password
}
func updateGitRepo(commitMsg string,pathToRepo string){
	username,password:=returnGitRepo()
	cmd:=exec.Command("git","add",".")
	cmd.Dir=pathToRepo
	out,err:=cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
	cmd=exec.Command("git","commit","-m",commitMsg)
	cmd.Dir=pathToRepo
	out,err=cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
	path := pathToRepo

	r, err := git.PlainOpen(path)
	CheckIfError(err)

	Info("git push")
	// push using default options
	err = r.Push(&git.PushOptions{Auth:&http.BasicAuth{Username:username,Password:password} })
	CheckIfError(err)
}
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"io"

	"path"
	"errors"

	"strings"
	"context"
	"log"
	"encoding/json"
	"encoding/csv"
	"github.com/chromedp/chromedp"
	"time"


	
)
type solved struct {
	name string
	code string
	lang string
	id string
}

func main(){
	/*
	
	
		
	
	solutions,err:=fetchSolutionByID(username,password,"541c8630095125aba6000c00")
	if err==nil{
		fmt.Println(solutions)
		writeToFile(solutions,"",baseURL,apiSecret,"pr4k")
	}
	*/
	apiSecret:=returnAuth()
	baseURL:="https://www.codewars.com/api/v1"
	username,password:=returnIdPassword()
	pathToRepo:="/home/pr4k/go_proj/src/codeWarsSolution"
	err:=uploadToRepo(username,password,apiSecret,baseURL,pathToRepo,"pr4k")
	if err != nil {
		fmt.Println(err)
	}
	

}

func fetchUserInfo(apiSecret string , baseURL string,username string) (string,error) {
	url:=fmt.Sprintf("users/%s",username)
	req,_:=http.NewRequest("GET",fmt.Sprintf("%s/%s",baseURL,url),nil)
	req.Header.Add("Authorization" ,apiSecret)
	res,err:=http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!=nil{
		return "",err
	}
	body, _ := ioutil.ReadAll(res.Body)
	s:=string(body)
	return s,nil

}

func fetchUserCompletedChallenge(apiSecret string , baseURL string,username string,page int) (string,error){

	url:=fmt.Sprintf("users/%s/code-challenges/completed?page=%d",username,page)
	req,_:=http.NewRequest("GET",fmt.Sprintf("%s/%s",baseURL,url),nil)
	req.Header.Add("Authorization" ,apiSecret)
	res,err:=http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!=nil{
		return "",err
	}
	body, _ := ioutil.ReadAll(res.Body)
	s:=string(body)
	return s,nil
}

func fetchChallengeInfoByID(apiSecret string , baseURL string,username string,id string) (string,error){

	url:=fmt.Sprintf("code-challenges/%s",id)
	req,_:=http.NewRequest("GET",fmt.Sprintf("%s/%s",baseURL,url),nil)
	req.Header.Add("Authorization" ,apiSecret)
	res,err:=http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!=nil{
		return "",err
	}
	body, _ := ioutil.ReadAll(res.Body)
	s:=string(body)
	return s,nil
}

func fetchChallengeInfoBySlug(apiSecret string , baseURL string,username string,slug string) (string,error){

	url:=fmt.Sprintf("code-challenges/%s",slug)
	req,_:=http.NewRequest("GET",fmt.Sprintf("%s/%s",baseURL,url),nil)
	req.Header.Add("Authorization" ,apiSecret)
	res,err:=http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!=nil{
		return "",err
	}
	body, _ := ioutil.ReadAll(res.Body)
	s:=string(body)
	return s,nil
}
// Returns a Struct of all solved problems solution id name and code
// It scrapes data from the solved page 
func fetchAllSolution(username string,password string)([]solved,error){
	var solutions []solved
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture screenshot of an element
	

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, chromedp.Navigate(`https://www.codewars.com/users/sign_in`)); err != nil {
		log.Fatal(err)
		return solutions,err
	}

	if err := chromedp.Run(ctx,
		chromedp.SetValue(`#user_email`, username, chromedp.ByID),
		chromedp.SetValue(`#user_password`, password, chromedp.ByID),
		chromedp.Submit(`#new_user`),
		chromedp.Sleep(10 * time.Second),
		chromedp.Navigate(`https://www.codewars.com/users/pr4k/completed_solutions`),
		//chromedp.ScrollIntoView(`document.querySelector(".js-infinite-marker")`,chromedp.ByJSPath),
		//chromedp.WaitNotPresent(`document.querySelector(".js-infinite-marker")`,chromedp.ByJSPath),

		chromedp.Sleep(10 * time.Second),
				); err != nil {
			log.Fatal(err)
			return solutions,err
		}

	
		var html string
		if err := chromedp.Run(ctx,
			chromedp.InnerHTML(`document.querySelector(".items-list")`,&html,chromedp.ByJSPath) ); err != nil {
					log.Fatal(err)
					return solutions,err
				}

		var text string
		
		if err := chromedp.Run(ctx,
			chromedp.Text(`document.querySelector(".items-list")`,&text,chromedp.ByJSPath) ); err != nil {
					log.Fatal(err)
					return solutions,err
				}
		
		

		
		var code string
		var name string
		var lang string
		var id string
		//log.Println(len(strings.Split(text,"months")),strings.Split(strings.Split(text,"months")[1],"kyu")[1])
		initial:=strings.Split(text,"months")[0]
		//fmt.Println(initial)
		//fmt.Println(html)
		for _,solution:=range strings.Split(initial,"month"){
	
			name=strings.Split(strings.Split(solution,"kyu")[1],":")[0]
			if len(strings.Split(name,"Go"))==2{
				name=strings.Split(name,"Go")[0]
				lang="Go"
				id=strings.Split(strings.Split(html,name)[0],"/kata/")[len(strings.Split(strings.Split(html,name)[0],"/kata/"))-1]
				id=id[:len(id)-2]
				html=strings.Split(html,name)[1]
			

			
			}
			if len(strings.Split(name,"Python"))==2{
				name=strings.Split(name,"Python")[0]
				lang="Python"
				id=strings.Split(strings.Split(html,name)[0],"/kata/")[len(strings.Split(strings.Split(html,name)[0],"/kata/"))-1]
				id=id[:len(id)-2]
				html=strings.Split(html,name)[1]
		
			
			}
			//fmt.Println(html)
			code=strings.TrimSpace(strings.Join(strings.Split(strings.Split(solution,"kyu")[1],":")[1:],":"))
			code=code[:len(code)-1]
			solutions=append(solutions,solved{name,code,lang,id})
		}
		for _,solution:=range strings.Split(text,"months")[1:len(strings.Split(text,"months"))-1]{
			
			name=strings.Split(strings.Split(solution,"kyu")[1],":")[0]
			if len(strings.Split(name,"Go"))==2{
				name=strings.Split(name,"Go")[0]
				lang="Go"
				id=strings.Split(strings.Split(html,name)[0],"/kata/")[len(strings.Split(strings.Split(html,name)[0],"/kata/"))-1]
				id=id[:len(id)-2]
				html=strings.Split(html,name)[1]
			
			
			}
			if len(strings.Split(name,"Python"))==2{
				name=strings.Split(name,"Python")[0]
				lang="Python"
				id=strings.Split(strings.Split(html,name)[0],"/kata/")[len(strings.Split(strings.Split(html,name)[0],"/kata/"))-1]
				id=id[:len(id)-2]
				html=strings.Split(html,name)[1]
				
		
			}
			code=strings.TrimSpace(strings.Join(strings.Split(strings.Split(solution,"kyu")[1],":")[1:],":"))
			code=code[:len(code)-1]
			solutions=append(solutions,solved{name,code,lang,id})

		}
		
		return solutions,nil
}
//returns your solution by kata id 
func fetchSolutionByID(username string,password string,id string)(solved,error) {
	res,err:=fetchAllSolution(username,password)
	if err!=nil{
		return solved{"","","",""},err
	}
	for _,solution:=range res{
		if solution.id==id{
			fmt.Println("Yep")
			return solution,nil
		}
	}
	return solved{"","","",""},errors.New("Can't Find Id in solutions")
}

// writes kata solution to a file with app description
func writeToFile(kata solved,folderPath string,baseURL string,apiSecret string,user string){
	var description map[string]interface{}
	
	res,err:=fetchChallengeInfoByID(apiSecret,baseURL,user,kata.id)
	if err==nil{
		json.Unmarshal([]byte(res),&description)
	}
	fileName:=strings.Join(strings.Split(kata.name,"/")," ")
	
	fileName=path.Join(folderPath,fileName)+extensions(strings.ToUpper(kata.lang))
	
	os.Create(fileName)
	file,err:=os.OpenFile(fileName,os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666,)
	if err!=nil{
		log.Fatal(err)

	}
	defer file.Close()
	byteSlice:=[]byte(addComment(description["description"].(string),strings.ToUpper(kata.lang), kata.code))
	_,err=file.Write(byteSlice)
	if err!=nil{
		log.Fatal(err)
	}
	
}
func addComment(description string,lang string,code string)string{
	if lang=="PYTHON"{
		return "'''"+description+"'''\n"+code
	}
	if lang=="GO"{
		return "/*"+description+"*/\n"+code
	}
	return code
}

func uploadToRepo(username string,password string,apiSecret string,baseURL string,pathToRepo string,user string)error{
	res,err:=fetchAllSolution(username,password)
	if err != nil {
		return err
	}
	csv_file, _ := os.Open(path.Join(pathToRepo,"log.csv"))
	r := csv.NewReader(csv_file)
	
	var database [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
		log.Fatal(err)
		}
 
		database=append(database,record)
	}
	file, err := os.Create(path.Join(pathToRepo,"log.csv"))
    if err != nil {
		return err
	}
    defer file.Close()
	
	
    writer := csv.NewWriter(file)
	defer writer.Flush()
	
	fmt.Println("Fetched solution")
	writer.Write([]string{"Name","Id","Language"})
	for _,solution := range res{
		fmt.Println([]string{solution.name,solution.id,solution.lang})
		if doesNotContains(database,[]string{solution.name,solution.id,solution.lang}){
		
		writeToFile(solution,pathToRepo,baseURL,apiSecret,user)
		}
		writer.Write([]string{solution.name,solution.id,solution.lang})
		
	}
	writer.Flush()
	
	fmt.Println("Updated the folder")
	updateGitRepo("Added Code",pathToRepo)

	return nil
	
}

func doesNotContains(s [][]string, e []string) bool {
    for _, a := range s {
        if a[2] == e[2] && a[1]==e[1] {
            return false
        }
    }
    return true
}

func extensions(lang string)string{
	if lang=="PYTHON"{
		return ".py"
	}
	if lang=="GO"{
		return ".go"
	}
	return ""
}


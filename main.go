package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"errors"

	"strings"
	"context"
	"log"
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
	apiSecret:=returnAuth()
	baseURL:="https://www.codewars.com/api/v1"

	
	res,err:= fetchChallengeInfoByID(apiSecret,baseURL,"pr4k","526dbd6c8c0eb53254000110")
	if err==nil{
			fmt.Println(res)
	}
	username,password:=returnIdPassword()
	solutions,err:=fetchSolutionByID(username,password,"541c8630095125aba6000c00")
	if err==nil{
		fmt.Println(solutions)
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

		log.Println("Crossed 3")
		var html string
		if err := chromedp.Run(ctx,
			chromedp.InnerHTML(`document.querySelector(".items-list")`,&html,chromedp.ByJSPath) ); err != nil {
					log.Fatal(err)
					return solutions,err
				}
		log.Println(len(strings.Split(html,"/kata/")))
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
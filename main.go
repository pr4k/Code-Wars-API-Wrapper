package main

import (
	"fmt"
	"net/http"
	"io/ioutil"


	
)

func main(){
	apiSecret:=returnAuth()
	baseURL:="https://www.codewars.com/api/v1"

	
	res,err:= fetchChallengeInfoByID(apiSecret,baseURL,"pr4k","52efefcbcdf57161d4000091")
	if err==nil{
			fmt.Println(res)
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


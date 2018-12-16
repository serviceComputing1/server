package model
import (
	"github.com/873421427/server/swapi"
	"encoding/json"
	//"bytes"
	"github.com/boltdb/bolt"
	"fmt"
	"strings"
)


//at first i want to return []byte directly,
//but i found the json data received in the client will
//have \"\",so i have to change data format that is returned
func GetPeople(begin int, len int) []swapi.People{
	if !open{
		return nil;
	}
	//var ps []byte =nil
	var people []swapi.People 
	db.View(func(tx *bolt.Tx) error{
		b :=tx.Bucket([]byte("people"))
		c := b.Cursor()
		count :=1
		for k,v:=c.First();k!= nil; k,v =c.Next(){
			if count >= begin && count < begin+len{
				var tmpPeople swapi.People
				json.Unmarshal(v,&tmpPeople)
				people = append(people,tmpPeople)

			}
			count++
		}
		return nil
	})

	return people
}

func Search(bucketName string, searchName string, begin int, len int)[]swapi.People{
	var outcome []swapi.People
	db.View(func(tx * bolt.Tx) error{
		b :=tx.Bucket([]byte(bucketName))

		//prefix := []byte(searchName)
		count :=1
		b.ForEach(func(k,v []byte)error{
			if strings.Contains(strings.ToLower(string(k)),strings.ToLower(searchName)){
				fmt.Printf("find one : %s\n",k)
				if count>=begin && count<begin+len{
					var tmp swapi.People
					json.Unmarshal(v,&tmp)
					//fmt.Println(tmp)
					outcome= append(outcome,tmp)
				}	
				count++
			}
		
			return nil
		})
		return nil
	})
	return outcome
}


func GetTotalNumOfSearch(bucketName string, searchName string) int{
	count :=0
	db.View(func(tx *bolt.Tx)error{
		b:=tx.Bucket([]byte(bucketName))
		//prefix := []byte(searchName)
		//count =:1
		b.ForEach(func(k,v []byte)error{
			if strings.Contains(strings.ToLower(string(k)),strings.ToLower(searchName)){
				count++
			}
			return nil
		})
		return nil
	})
	return count
}

func GetTotalNumOfPeople() int{
	count :=0
	db.View(func(tx *bolt.Tx) error{
		b:=tx.Bucket([]byte("people"))
		b.ForEach(func(k,v []byte)error{
			count++
			return nil
		})
		return nil
	})
	return count
}

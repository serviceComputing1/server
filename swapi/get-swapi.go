package swapi

import (
	"net/http"
	"fmt"
	//"io"
	"crypto/tls"
	"github.com/boltdb/bolt"
	"log"
	"encoding/json"
	//"reflect"
	"runtime"
	"path"
	"strconv"
	//"time"
	//simplejson "github.com/bitly/go-simplejson"
)

var db *bolt.DB
var open bool

func Open() error {
    var err error
    _, filename, _, _ := runtime.Caller(0)  // get full path of this file
    dbfile := path.Join(path.Dir(filename), "my.db")
    //config := &bolt.Options{Timeout: 1 * time.Second}
    db, err = bolt.Open(dbfile, 0600, nil)
    if err != nil {
        log.Fatal(err)
	}
    open = true
    return nil
}

func Close() {
    open = false
    db.Close()
}




type People struct {
	Name         string        `json:"name"`
	Height       string        `json:"height"`
	Mass         string        `json:"mass"`
	HairColor    string        `json:"hair_color"`
	SkinColor    string        `json:"skin_color"`
	EyeColor     string        `json:"eye_color"`
	BirthYear    string        `json:"birth_year"`
	Gender       string        `json:"gender"`
	Homeworld    string        `json:"homeworld"`
	FilmURLs     []filmURL     `json:"films"`
	SpeciesURLs  []speciesURL  `json:"species"`
	VehicleURLs  []vehicleURL  `json:"vehicles"`
	StarshipURLs []starshipURL `json:"starships"`
	Created      string        `json:"created"`
	Edited       string        `json:"edited"`
	URL          string        `json:"url"`
}


type Film struct {
	Title         string         `json:"title"`
	EpisodeID     int            `json:"episode_id"`
	OpeningCrawl  string         `json:"opening_crawl"`
	Director      string         `json:"director"`
	Producer      string         `json:"producer"`
	CharacterURLs []characterURL `json:"characters"`
	PlanetURLs    []planetURL    `json:"planets"`
	StarshipURLs  []starshipURL  `json:"starships"`
	VehicleURLs   []vehicleURL   `json:"vehicles"`
	SpeciesURLs   []speciesURL   `json:"species"`
	Created       string         `json:"created"`
	Edited        string         `json:"edited"`
	URL           string         `json:"url"`
}

type Planet struct {
	Name           string        `json:"name"`
	RotationPeriod string        `json:"rotation_period"`
	OrbitalPeriod  string        `json:"orbital_period"`
	Diameter       string        `json:"diameter"`
	Climate        string        `json:"climate"`
	Gravity        string        `json:"gravity"`
	Terrain        string        `json:"terrain"`
	SurfaceWater   string        `json:"surface_water"`
	Population     string        `json:"population"`
	ResidentURLs   []residentURL `json:"residents"`
	FilmURLs       []filmURL     `json:"films"`
	Created        string        `json:"created"`
	Edited         string        `json:"edited"`
	URL            string        `json:"url"`
}

type Species struct {
	Name            string      `json:"name"`
	Classification  string      `json:"classification"`
	Designation     string      `json:"designation"`
	AverageHeight   string      `json:"average_height"`
	SkinColors      string      `json:"skin_colors"`
	HairColors      string      `json:"hair_colors"`
	EyeColors       string      `json:"eye_colors"`
	AverageLifespan string      `json:"average_lifespan"`
	Homeworld       string      `json:"homeworld"`
	Language        string      `json:"language"`
	PeopleURLs      []personURL `json:"people"`
	FilmURLs        []filmURL   `json:"films"`
	Created         string      `json:"created"`
	Edited          string      `json:"edited"`
	URL             string      `json:"url"`
}


type Starship struct {
	Name                 string      `json:"name"`
	Model                string      `json:"model"`
	Manufacturer         string      `json:"manufacturer"`
	CostInCredits        string      `json:"cost_in_credits"`
	Length               string      `json:"length"`
	MaxAtmospheringSpeed string      `json:"max_atmosphering_speed"`
	Crew                 string      `json:"crew"`
	Passengers           string      `json:"passengers"`
	CargoCapacity        string      `json:"cargo_capacity"`
	Consumables          string      `json:"consumables"`
	HyperdriveRating     string      `json:"hyperdrive_rating"`
	MGLT                 string      `json:"MGLT"`
	StarshipClass        string      `json:"starship_class"`
	PilotURLs            []personURL `json:"pilots"`
	FilmURLs             []filmURL   `json:"films"`
	Created              string      `json:"created"`
	Edited               string      `json:"edited"`
	URL                  string      `json:"url"`
}

type Vehicle struct {
	Name                 string      `json:"name"`
	Model                string      `json:"model"`
	Manufacturer         string      `json:"manufacturer"`
	CostInCredits        string      `json:"cost_in_credits"`
	Length               string      `json:"length"`
	MaxAtmospheringSpeed string      `json:"max_atmosphering_speed"`
	Crew                 string      `json:"crew"`
	Passengers           string      `json:"passengers"`
	CargoCapacity        string      `json:"cargo_capacity"`
	Consumables          string      `json:"consumables"`
	VehicleClass         string      `json:"vehicle_class"`
	PilotURLs            []personURL `json:"pilots"`
	FilmURLs             []filmURL   `json:"films"`
	Created              string      `json:"created"`
	Edited               string      `json:"edited"`
	URL                  string      `json:"url"`
}




type personURL string


type residentURL string

type filmURL string
type speciesURL string
type vehicleURL string
type starshipURL string

type characterURL string
type planetURL string



var url  = [6]string{"https://swapi.co/api/people/", "https://swapi.co/api/planets/", "https://swapi.co/api/films/", "https://swapi.co/api/species/", "https://swapi.co/api/vehicles/", "https://swapi.co/api/starships/"}

var count = [6]int{87,61,7,37,39,37}

var p People;
var f  Film;
var pl Planet;
var s  Species;
var ve  Vehicle;
var ss  Starship;



func Get_People(){


	tr := &http.Transport{
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport:tr}

	

	for j:=1;j<=87;j++ {
		resp, err := client.Get("https://swapi.co/api/people/" + strconv.Itoa(j))
		fmt.Printf("https://swapi.co/api/people/" + strconv.Itoa(j) + "\n")
		if err!= nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var dec = json.NewDecoder(resp.Body)
		err = dec.Decode(&p)
		
		fmt.Println(p)
	
		if !open{
			fmt.Println("db not open")
			return
		}
		fmt.Println(db)

		db.Update(func(tx *bolt.Tx) error{
			b, err:= tx.CreateBucketIfNotExists([]byte("people"))
			buf, err := json.Marshal(p)
			if err !=nil{
				fmt.Println(err)
			}
			err = b.Put([]byte(p.Name),buf)
			if err !=nil{
				fmt.Println("sadasf")
				fmt.Println(err)
			}
			return err
		})		
		
	}
				
}


func Get_Films(){


	tr := &http.Transport{
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport:tr}

	

	for j:=1;j<=6;j++ {
		resp, err := client.Get("https://swapi.co/api/films/" + strconv.Itoa(j))
		fmt.Printf("https://swapi.co/api/films/" + strconv.Itoa(j) + "\n")
		if err!= nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var dec = json.NewDecoder(resp.Body)
		err = dec.Decode(&f)
		
		fmt.Println(f)
	
		if !open{
			fmt.Println("db not open")
			return
		}
		fmt.Println(db)

		db.Update(func(tx *bolt.Tx) error{
			b, err:= tx.CreateBucketIfNotExists([]byte("films"))
			buf, err := json.Marshal(f)
			if err !=nil{
				fmt.Println(err)
			}
			err = b.Put([]byte(f.Title),buf)
			if err !=nil{
				fmt.Println("sadasf")
				fmt.Println(err)
			}
			return err
		})		
		
	}
				
}



func Get_Planets(){


	tr := &http.Transport{
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport:tr}

	

	for j:=1;j<=61;j++ {
		resp, err := client.Get("https://swapi.co/api/planets/" + strconv.Itoa(j))
		fmt.Printf("https://swapi.co/api/planets/" + strconv.Itoa(j) + "\n")
		if err!= nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var dec = json.NewDecoder(resp.Body)
		err = dec.Decode(&pl)
		
		fmt.Println(pl)
	
		if !open{
			fmt.Println("db not open")
			return
		}
		fmt.Println(db)

		db.Update(func(tx *bolt.Tx) error{
			b, err:= tx.CreateBucketIfNotExists([]byte("planets"))
			buf, err := json.Marshal(pl)
			if err !=nil{
				fmt.Println(err)
			}
			err = b.Put([]byte(pl.Name),buf)
			if err !=nil{
				fmt.Println("sadasf")
				fmt.Println(err)
			}
			return err
		})		
		
	}
				
}


func Get_Species(){


	tr := &http.Transport{
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport:tr}

	

	for j:=1;j<=37;j++ {
		resp, err := client.Get("https://swapi.co/api/species/" + strconv.Itoa(j))
		fmt.Printf("https://swapi.co/api/species/" + strconv.Itoa(j) + "\n")
		if err!= nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var dec = json.NewDecoder(resp.Body)
		err = dec.Decode(&s)
		
		fmt.Println(s)
	
		if !open{
			fmt.Println("db not open")
			return
		}
		fmt.Println(db)

		db.Update(func(tx *bolt.Tx) error{
			b, err:= tx.CreateBucketIfNotExists([]byte("species"))
			buf, err := json.Marshal(s)
			if err !=nil{
				fmt.Println(err)
			}
			err = b.Put([]byte(s.Name),buf)
			if err !=nil{
				fmt.Println("sadasf")
				fmt.Println(err)
			}
			return err
		})		
		
	}
				
}


func Get_Vehicles(){


	tr := &http.Transport{
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport:tr}

	

	for j:=1;j<=39;j++ {
		resp, err := client.Get("https://swapi.co/api/vehicles/" + strconv.Itoa(j))
		fmt.Printf("https://swapi.co/api/vehicles/" + strconv.Itoa(j) + "\n")
		if err!= nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var dec = json.NewDecoder(resp.Body)
		err = dec.Decode(&ve)
		
		fmt.Println(ve)
	
		if !open{
			fmt.Println("db not open")
			return
		}
		fmt.Println(db)

		db.Update(func(tx *bolt.Tx) error{
			b, err:= tx.CreateBucketIfNotExists([]byte("vehicles"))
			buf, err := json.Marshal(ve)
			if err !=nil{
				fmt.Println(err)
			}
			err = b.Put([]byte(ve.Name),buf)
			if err !=nil{
				fmt.Println("sadasf")
				fmt.Println(err)
			}
			return err
		})		
		
	}
				
}


func Get_StarShips(){


	tr := &http.Transport{
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport:tr}

	

	for j:=1;j<=37;j++ {
		resp, err := client.Get("https://swapi.co/api/starships/" + strconv.Itoa(j))
		fmt.Printf("https://swapi.co/api/starships/" + strconv.Itoa(j) + "\n")
		if err!= nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var dec = json.NewDecoder(resp.Body)
		err = dec.Decode(&ss)
		
		fmt.Println(ss)
	
		if !open{
			fmt.Println("db not open")
			return
		}
		fmt.Println(db)

		db.Update(func(tx *bolt.Tx) error{
			b, err:= tx.CreateBucketIfNotExists([]byte("starships"))
			buf, err := json.Marshal(ss)
			if err !=nil{
				fmt.Println(err)
			}
			err = b.Put([]byte(ss.Name),buf)
			if err !=nil{
				fmt.Println("sadasf")
				fmt.Println(err)
			}
			return err
		})		
		
	}
				
}


func Get_db_info(db_name string) {
	if !open{
		return 
	}
	db.View(func(tx * bolt.Tx) error{
		b := tx.Bucket([]byte(db_name))
		b.ForEach(func(k,v []byte) error{
			fmt.Printf("%T %T\n",k,v)
			return nil
		})
		return nil
	})
}

func GetPeople() error{
	if !open{
		return nil
	}
	fmt.Println(db)
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("testBucket1")).Cursor()
		for k,v:= c.First();k!=nil;k,v=c.Next(){
			fmt.Printf("%s %s \n",k,v)
		}

		return nil
	})

	return err
}

	//fmt.Println(values)
	/*
	t, err := dec.Token()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T %v\n",t,t)

	for {
		t,err = dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil{
			fmt.Println(err)
		}
		fmt.Printf("%T %v\n",t,t)
	}
	*/
	
	//err = json.Unmarshal([]byte(body),&v)
	/*
	js,err := simplejson.NewJson([]byte(body))
	arr,_ := js.Get("name").String()
	fmt.Println(arr)
	*/

/*
func TestUpdate() {
	db, err := bolt.Open("my.db",0600,nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b,err := tx.CreateBucket([]byte("testBucket"))
		if err !=nil{
			return err
		}
		err = b.Put([]byte("answer"),[]byte("42"))

		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("testBucket"))
		v:=b.Get([]byte("answe"))
		if v ==nil {
			fmt.Println("empty")
		}
		fmt.Printf("answer is %s\n",v)
		return nil
	})
}


*/



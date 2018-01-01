package main

import(
	"math/rand"
	"strconv"
	"time"
	"strings"
)

type fields [][]cases

type cases struct{
	Name string
	Type string
	Hidden bool
	Bomb int
	Neighboor []string
}


var (
	src = rand.NewSource(int64(time.Now().Local().Nanosecond()))
)


func NameInt(a string)(c,d int){
	b := strings.Split(a,"-")
	c,_ = strconv.Atoi(b[0])
	d,_ = strconv.Atoi(b[1])
	return
}



func makemin(b int)(a int){
	a = b - 1
	if a == -1{
		a = 0
	}
	return
}

func makemax(a,b int)(c int){
	c = a + 2;
	if c > b {
		c = b
	}
	return
}


func (me fields)makeNeig(nbr int){
	for a,_ := range me{
		for b,_ := range me[a]{
			for min := makemin(a); min < makemax(a,nbr);min = min +1{
				for emin := makemin(b);emin < makemax(b,nbr);emin = emin + 1{
					if me[a][b].Name != me[min][emin].Name{ 
						me[a][b].Neighboor=append(me[a][b].Neighboor,me[min][emin].Name)
						if me[min][emin].Type == "bombe"{
							me[a][b].Type = "Neig"
							me[a][b].Bomb = me[a][b].Bomb +1
						}
					}
				}
			}
			
		}		
	}
}


func (me fields)makeBomb(mine int){
	aa := rand.New(src)
	for {
		for a,_ := range me{
			for b,_ := range me[a]{
				ar := aa.Intn(100)
				if ar > 49 {
					mine = mine - 1;
					me[a][b].Type = "bombe"
					if mine == 0  {
						return
					}
				}
			}
		}
	}
}

func nameMe(a,b int)(s string){
	s = strconv.Itoa(a)+"-"+strconv.Itoa(b)
	return
}

func intMe(a string)(b,c int){
	g := strings.Split(a,"-")
	b,_ = strconv.Atoi(g[0])
	c,_ = strconv.Atoi(g[1])
	return
}

func (me fields)revelme(a string){
	un,deux := intMe(a)
	if me[un][deux].Type == "empty" && me[un][deux].Hidden == true {
		me[un][deux].Hidden = false
		for _,ok := range me[un][deux].Neighboor{
			me.revelme(ok)
		}
	}
	if me[un][deux].Type == "neig"{
		me[un][deux].Hidden = false
	} 
	
}

func makeFields(nbr,mine int)(ar fields){
	are := [][]cases{}
	for i := 0 ; i < nbr;i++{
		t := []cases{}
		for a := 0;a<nbr;a++{
			t = append(t,cases{nameMe(i,a),"empty",true,0,nil})
		}
		are = append(are,t)
	}
	ar = are
	ar.makeBomb(mine)
	ar.makeNeig(nbr)
	return
}

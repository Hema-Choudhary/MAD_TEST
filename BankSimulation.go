package main

//imports 
import ("fmt"
		"time"
		"os"	
		"strconv"
		"sync"	
		"strings"
		)

//create a type Cashier
type Cashier struct{
	id 		int64
	state	string
}

//create a type Customer
type Customer struct{
	id		int64
}

//Method of Cashier that will serve a customer for a given time 
func (ch Cashier) HandleCust(queue <-chan Customer, cTime int64){
	defer wg.Done()
	ch.state = "unoccupied"
	
	for{
			cust := <-queue
			if ch.state == "unoccupied" && cust.id != 0 {
			ch.state = "occupied"
			fmt.Println(time.Now().Format("2006-01-02 3:4:5")," --> Cashier ", ch.id,": Customer ",cust.id," Started")
			time.Sleep(time.Duration(cTime) * time.Second)
			fmt.Println(time.Now().Format("2006-01-02 3:4:5")," --> Cashier ", ch.id,": Customer ",cust.id," Completed")
			ch.state = "unoccupied"
		}else {
			break

		}	
	}
	
}


//This function will assign ids to cashiers and spawn HandleCust for each cashier
func cashiersStartWork(nCashr int64, queue <-chan Customer, cTime int64){
	var i int64
	var ch Cashier
	fmt.Println(time.Now().Format("2006-01-02 3:4:5")," --> Bank Simulation Started")
	for i=1; i<=nCashr; i++{
		wg.Add(1)
		ch.id = i
		go ch.HandleCust(queue, cTime)		
	}
	wg.Wait()
	fmt.Println(time.Now().Format("2006-01-02 3:4:5")," --> Bank Simulated Ended")
}


//Customer method that adds Customers to the queue(channel)
func (c Customer) queueUp(queue chan<- Customer){
	queue<-c	
}

//this function will assign ids to the customers and provide them to queueUp to be added to the queue
func arrangeCust(nCusts int64, queue chan<- Customer){
	var c Customer
	var i int64 

	for i=1; i<=nCusts  ;i++{
		c.id = i
		c.queueUp(queue)
	}
	close(queue)
 
}

//Declare a waitgroup to wait for all goroutines to finish
var wg sync.WaitGroup

//declare variables
var noOfCashiers, noOfCustomers, timePerCust int64

//assign commantline args to variable as per thier flags
func getArguments(arg1, arg2, arg3 string){
	if a1 := strings.Split(arg1,"="); a1[0]=="--numCashiers" {
		noOfCashiers,_ = strconv.ParseInt(a1[1],10,64)
	}else if a1[0] == "--numCustomers"{
		noOfCustomers,_ = strconv.ParseInt(a1[1],10,64)
	}else{
		timePerCust,_= strconv.ParseInt(a1[1],10,64)
	}

	if a2 := strings.Split(arg2,"="); a2[0]=="--numCashiers"{
		noOfCashiers,_ = strconv.ParseInt(a2[1],10,64)
	}else if a2[0] == "--numCustomers"{
		noOfCustomers,_ = strconv.ParseInt(a2[1],10,64)
	}else{
		timePerCust,_	= strconv.ParseInt(a2[1],10,64)
	}

	if a3 := strings.Split(arg3,"="); a3[0]=="--numCashiers"{
		noOfCashiers,_ = strconv.ParseInt(a3[1],10,64)
	}else if a3[0] == "--numCustomers"{
		noOfCustomers,_ = strconv.ParseInt(a3[1],10,64)
	}else{
		timePerCust,_	= strconv.ParseInt(a3[1],10,64)	
	}

}


func main(){
	//collect commandline args 
	arg1 :=	os.Args[1]
	arg2 :=	os.Args[2]
	arg3 :=	os.Args[3]

	//assign commandline values to varibles
	getArguments(arg1, arg2,arg3)
	
	//create a buffered channel (queue) for the customers
	var queue = make(chan Customer,noOfCustomers)

	//add all customers in the queue	
	arrangeCust(noOfCustomers,queue)

	//Run the goroutines (cashiers) to serve customers
	cashiersStartWork(noOfCashiers, queue, timePerCust)
}
package main
import "fmt"

//our struct representing a person 
type person struct{
	name string
	age int
}
//Return true,and the older person in a group of persons
//or false,and nil if the group is empty.

func Older(people ...person)(bool,person){//variadic function.
	if(len(people) == 0){return false, person{}}// the group is empty.
	older := people[0]//The first one is the older For now.
	//loop through the slice people.
	
	for _, value := range people{//we dont need the index.
	//compare the current person's age with the oldest one so far
		if value.age > older.age {
			older = value //is value is older, replace older.
		}
	}
	return true, older
}


func main(){
	//two variables to be used by our program.
	var(
		ok bool
		older person
	)
	
	// declare some persons.
	paul := person{"Paul", 23};
	jim := person{"Jim", 24};
	sam := person{"Sam", 84};
	rob := person{"Rob", 54};
	karl := person{"Karl", 19};
	//who is older? PAUL OR JIM?
	_, older = Older(paul, jim)// notice how we used the black identifier
	fmt.Println("The older of Paul and Jim is:",older.name,older.age)
	_, older = Older(paul, jim, sam)
	fmt.Println("the older of Paul, Jim, Sam is:", older.name,older.age)
	_, older = Older(paul, jim, sam, rob)
	fmt.Println("The older of Pal, Jim,Sam and Rob is:", older.name,older.age)
	_, older = Older(karl)
	fmt.Println("when Karl is alone in a group ,the older is:", older.name,older.age)
	//is there an older person in an empty group?
	ok, older = Older()//this time we use the boolean variable ok
	if !ok {
		fmt.Println("In an empty group there is no older person")
	}
}
	

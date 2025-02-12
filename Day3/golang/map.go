package main
import "fmt"

//Map will store elements in key,val format
//"contact": "mobilenumber"
//"name": "123-456-8901"
//you can store env variables
//JAVA_HOME=/usr/lib/jvm/jdk-11

func main() {
	//string within square bracket is key
	//string outside square bracket is value
	toolsPath := map[string]string {
		"java_home": "/usr/lib/jdk-11",
		"mvn_home" : "/usr/share/maven",
	}

	//Given a key, map will be able to retrieve the value 
	fmt.Println ( "Java home directory : ", toolsPath["java_home"] )

	//add a key,value pair into an existing map
	toolsPath["go_home"] = "/usr/go"

	//iterating a map and printing its values
	for key,value := range toolsPath {
		fmt.Println ( key, value )
	}

	//delete a key-value pair
	delete(toolsPath, "go_home")
	fmt.Println(toolsPath)
}

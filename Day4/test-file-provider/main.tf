terraform {
   required_providers {
       localfile = {
          source = "tektutor/file"
       }
   }
}

resource "localfile" "myfile" {
    file_name = "./test.txt"
    file_content = "Testing"

}

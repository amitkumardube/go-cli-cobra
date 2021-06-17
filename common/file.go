package common

import (
	"os"
	"encoding/json"
	"log"
)


func Read_file(file_name string) conf{
	file , err := os.Open(file_name)

	if err != nil {
		if os.IsNotExist(err) {
			return conf{}
		}else {
			log.Fatal(err)
			return conf{}
		}
	}

	defer file.Close()

	decode_file := json.NewDecoder(file)
	decode_file.UseNumber()

	var config conf

	err = decode_file.Decode(&config)
	if err != nil {
		log.Fatal(err)
		return conf{}
	}

	return config


}

func Write_file(file_name string, key string, value string){
	conf := Read_file(file_name)

	conf.Cli.Project_id = value
	file, _ := json.MarshalIndent(conf, "", " ")
	_ = os.WriteFile(file_name, file, 0644)
	//file, _ := os.Create(file_name)
	//defer file.Close()
	//encode_file := json.NewEncoder(file)
	//encode_file.Encode(conf)
}


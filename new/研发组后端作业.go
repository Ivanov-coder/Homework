package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"
)


func choose() string {
	var choose string
	fmt.Println("请输入1、2、3中的任意一个数字进入对应模块")
	fmt.Print("1.命令行\n2.使用说明\n3.退出程序\n请输入指令：\n")
	fmt.Scanln(&choose)
	return choose
}



func writer(data map[string]string) {
	file, err := os.Create("info.json")
	if err != nil{
		fmt.Println("创建文件失败",err.Error())
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil{
		fmt.Println("写入文件失败",err)
	} else {
		fmt.Println("写入文件成功")
	}
}


func use() {
	fmt.Print("请输入指令: ")
	content := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	function, _, _ := reader.ReadLine()
	str := string(function)
	result := strings.Split(str," ")

	enterlog1 := (len(result) == 3) 
	enterlog2 := (len(result) == 2) 

	if (enterlog1) {

		command := result[0]
		Key := result[1]
		Value := result[2]

		switch command {

			case "Set","set","SET":
				fmt.Println("Now executing \"Set\"")
				content[Key] = Value
				writer(content)

			case "Setnx","SETNX","setnx":
				fmt.Println("Now executing \"Setnx\"")

				var data map[string]string
				file, err := os.ReadFile("info.json")

				if err != nil {
					fmt.Println("读取文件失败")

				} else {
					err := json.Unmarshal(file, &data)

					if err == nil{
						value_for_check := data[Key]
						fmt.Println(value_for_check)

						if (value_for_check == Value){
							fmt.Println(0)

						} else {
							fmt.Println(1)
							content[Key] = Value
							writer(content)
						}

					} else {
						fmt.Println(err.Error())	
					}
				}	

			case "LPUSH","Lpush","lpush":
				fmt.Println("Now executing \"Lpush\"")
				Key := []string{}
				res := append(Key,Value)
				fmt.Println(res)

			default:
				fmt.Println("请输入有效的指令")

			}

	} else if (enterlog2) {
		command := result[0]
		Key := result[1]
		switch command{
			case "Get","get","GET":
				fmt.Println("Now executing \"Get\"")

				var data map[string]string
					file, err := os.ReadFile("info.json")

					if err != nil {
						fmt.Println("读取文件失败")

					} else {
						err := json.Unmarshal(file, &data)
						value := data[Key]

						if (err == nil && value != ""){
							fmt.Printf("%s 对应的值是 %s\n",Key,value)
						} else {
							fmt.Println("无法获取对应值")
						}
					}

				case "DEL","del","Del":
					fmt.Println("Now executing \"DEL\"")

					var data map[string]string
					file, err := os.ReadFile("info.json")
					
					if err != nil {
						fmt.Println("读取文件失败")
					} else {
						json.Unmarshal(file, &data)
						// str := string(file)
						delete(data,Key)
						ctn, Err := json.Marshal(data)

						if Err != nil {
							fmt.Println("删除失败")
						} else {
							os.WriteFile("info.json",ctn,fs.ModeDevice)
							fmt.Println("删除成功")
						}
					}


				default:
					fmt.Println("请输入有效的指令")
		}

	} else {
		fmt.Println("请输入有效的指令")
		use()
	}
}



func explaination() {
	fmt.Println("\n这是使用说明")
	fmt.Println("进入各分支的按钮为1、2、3")
	fmt.Println("而在【1.命令行模块】中有多个指令")
	fmt.Println("比如：")
	fmt.Println("1、set Key value 将键值存储起来")
	fmt.Println("2、setnx Key value 如果键存在 返回0 如果键不存在 返回1并存储值")
	fmt.Println("3、get Key 获得对应键所对应的值")
	fmt.Println("4、del Key 删除所对应的键和值")
	fmt.Println("下面为具体例子：")
	fmt.Println("使用命令set aaa 123之后，会出现一个json文件")
	fmt.Println("文件内容如下：")
	fmt.Println("{\"aaa\":\"123\"}")
	fmt.Println("希望这则说明对你有所帮助")
	fmt.Println("\n现在进入命令行模块")
	use()
}



func exit() {
	fmt.Println("已退出程序")
	os.Exit(-1)
}


func main() {
	choice := choose()
	switch choice{
		case "1":
			use()
		case "2":
			explaination()
		case "3":
			exit()
		default:
			fmt.Println("输入有误，请输入1、2、3中的任意一个数字")
	}
}
package main

import (
	"log"

	"gitea.bee.anarckk.me/anarckk/go_util/go_encrypt/go_rsa"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test1() {
	var message = "hello world! I love fuck mm"

	priKey, pubKey, err := go_rsa.GenerateKeyPairStr()
	if err != nil {
		panic(err)
	}
	log.Println("priKey: " + priKey)
	log.Println("pubKey: " + pubKey)
	encrypted, err := go_rsa.EncryptByPubKeyStr(message, pubKey)
	if err != nil {
		panic(err)
	}
	log.Println("encrypted: " + encrypted)
	decrypted, err := go_rsa.DecryptByPriKeyStr(encrypted, priKey)
	if err != nil {
		panic(err)
	}
	log.Println("decrypted: " + decrypted)
}

func test2() {
	priKey := "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCSWSU05sQkPU6JD3SZP7J0Cc1NriFkkpfMtJhYdYDZ4gA9KSqbqn378UWsDIe36gYXJc6deyLkgYGAOtsGkbvfo3mWguk1awIVOHkYuj494C5NgTVdCTEfUs6HJ3CQXe9zl3YoiA1Jemp5BqFW60miQXayWifz7zkZDCvm95ynW6IaqDuPxtAA2DN201HmnNrvB5LA3ysc49raupp6gIil9W3XANRWOgbIpL8FC65Rzd+K2LT7yPH6co5ct7/4XKfPBHQ+vrWLzb6G45/RqiMrw0DUIBz7kACFspr7G0RjrkgCwDt1e71DO23zoM03CGvCkSTRQ+Z3nyaI65BkIjeJAgMBAAECggEAeV1eQqHBNUB2OanMiy5MwnhCftISNUJwir2VvX4sjgQjKJVUFWXtNpteqRB3GKkFxfp/fw/X3uIbUAj/DFKdGBiMw6nq1nbYclqz6jLBXTTlkTa+11nBF/Xm+iRV8BNGeXi472HsiuvvElDSSa+0D8/0LHIhweS4WDJE0jS0AAD0XvGMc+MbI232uzpClvk5wBCCcyNn8UM5PIgvmTc0IydSMidf0E8mih3FnpGXhjXXRWx4EFIxAFuBA165oBqBjhz01FxLgxZWcwLbgd7IuSlaaxtVYtFoLP5uxnWi/IkDLVQ8zq5B7rkIYJ2DA9BVS+QeHzo+nGbAzfGFz2ZkBQKBgQDf0RdT/SrXcxxyyUhFEeE2ew9rttFEtkw/KYv/XBTfn4eFYeciu7WrZN/DTiMlXbBMm8P6ZqY2EmKbWvcaQqkQc30JldYZczxK/5HO1gOwW9HZiayo40s9vQpVd/sE7qSx3gh4AHiv31tkBBE36TU4DOGEa6SjUSgaFsJSOfNgXwKBgQCnZF0BMKCx+gqha2HiCqJAH04i9jmAuwOc6v6IMhd4r7SzwaFvb09DYoFxi7Mkr7cvFqYRSANEh+PHbrZelJLJHnGDlwSxcinYgtMUH/hNHufzrw4BZy0uX+NTXGowI3NL3q2uzYSjntfyMCiOfcsFCf0vrr9uAGKg5apeDcnRFwKBgHGXH7zLlyujS0PibeBIE8HfsNLdBNZXotjHkDq4lAtuXoxORM028RucZYgspt+27dvjjhIOeLqmmA76msBkJoOn6UStG+zstCPoEysjKNofr6A1JEDOoogh4hXAf9BgAYwYALpOmvG/bRWUjtyOaikZOHdJXlRYwv6CoHq02JUHAoGAWMRuVa016mvQq11IoRhGhn6TbxLn145VEifEJvF5ZPS4fQLX20JJ5FAemNoee/v6xqvaERwBL5xofGAHsgxT8veD9uZlBLyn2Ds4OFnj0PHsy1svsCrI2Ojcol5FqZWDFN7Xd/Vgu2wG1FYZi8bFnLx5WYnv1iO6KdzhBdOGDK0CgYAOun0l1UkyovW9AEqD6Fn3sG23EnVf2P0WuDsnMvNVigWdktLsajwn50CRjkqTK9/tCPA0YsdD2NGyGPvM/nADZVTakOP/++Gm7TzYA4WZ+ehLiQBUMGtFdijEdkQSYVrRXpgxEiYKgdnHe1mrAD2EgwiIWBGL2xySUZ8+LHUe0w=="
	encrypted := "C3616vDiOW3t/yJ7mv8f43VdGhx9aPcQwmOcZFf35laE5coaePcEmudZpAa67Sd+QC+OXsx4WexkzX2mTTh9rQ4VpTwPjOEQLZMhE+DVhFZOYRQEhTSPIim6POmqh8IcjRQZLzurCcrLsWiBUfAz2mRfN1EHtGsY+sIL/2NzwA7CPz+d9TZC7BmfB5K+O0GTqx73Ge3+c94Rg2whr2Zk6uWSN7r8HmzEhptVzYsqwE0lAUhG+ilszy46h0M5SlkdsWqog8Jhi3q3VI++G79idwS0TQ7QlFUHkuGAWyKH8AAOWFKE8uIC0cOB2vxK67m0CZcGEWkoGceLWGknD7bypQ=="
	decrypted, err := go_rsa.DecryptByPriKeyStr(encrypted, priKey)
	if err != nil {
		panic(err)
	}
	log.Println(decrypted)
}

func main() {
	test2()
}

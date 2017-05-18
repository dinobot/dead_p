package main

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage:\n  %s dump_file db_file\n", os.Args[0])
		os.Exit(0)
	}

	fname := os.Args[1]
	db_name := os.Args[2]

	reader, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Filed to open dump file: %s\n", string(fname))
		return
	}

	db, err := bolt.Open(db_name, 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(reader)
	target := make([]interface{}, 0)
	err = decoder.Decode(&target)
	if err != nil {
		fmt.Printf("Appending raw data\n")
		stamp := strconv.FormatInt(time.Now().Unix(), 10)
		fmt.Printf("%s\n", stamp)
		dump_data, _ := ioutil.ReadFile(fname)
                db.Update(func(tx *bolt.Tx) error {
                        tx.CreateBucketIfNotExists([]byte("Snapshots"))
                        b := tx.Bucket([]byte("Snapshots"))
                        err := b.Put([]byte(stamp), dump_data)
                        return err
                })
	os.Exit(0)
	}

	fmt.Printf("Building from scratch\n")
	for _, item := range target {
		top := item.(map[string]interface{})
		stamp := top["timestamp"].(string)
		fmt.Printf("%s\n", stamp)
		data, _ := json.Marshal(top["snapshot"])
		db.Update(func(tx *bolt.Tx) error {
			tx.CreateBucketIfNotExists([]byte("Snapshots"))
			b := tx.Bucket([]byte("Snapshots"))
			err := b.Put([]byte(stamp), data)
			return err
		})
	}
}

package main

import (
	"fmt"
        "os"
	"io/ioutil"
	"github.com/boltdb/bolt"
)

func main() {
        timestamp := os.Args[1]
        dump_file := os.Args[2]
        db_name := os.Args[3]

	dump_data, err := ioutil.ReadFile(dump_file)
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := bolt.Open(db_name, 0600, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("timestamp: %s json_snapshot_file: %s\n", string(timestamp), string(dump_file))

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("Snapshots"))
		b := tx.Bucket([]byte("Snapshots"))
                err := b.Put([]byte(timestamp), dump_data)
		return err
	})
	defer db.Close()
}

package main

//func main() {
//	// Open the my.db data file in your current directory.
//	// It will be created if it doesn't exist.
//	db, err := bolt.Open("my.db", 0600, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	db.Update(func(tx *bolt.Tx) error {
//		bucket := tx.Bucket([]byte("block"))
//		if bucket == nil {
//			bucket, err = tx.CreateBucket([]byte("block"))
//			if err != nil {
//				log.Panic("创建bucket()失败")
//			}
//		}
//		bucket.Put([]byte("a"), []byte("b"))
//		return nil
//	})
//
//	db.View(func(tx *bolt.Tx) error {
//		bucket := tx.Bucket([]byte("block"))
//		if bucket == nil {
//			log.Panic("bucket 为空")
//		} else {
//			value := bucket.Get([]byte("a"))
//			fmt.Printf("%s", value)
//		}
//		return nil
//	})
//	defer db.Close()
//}

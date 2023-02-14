package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func initImage() {
	if _, err := os.Stat("../public/image"); os.IsNotExist(err) {
		os.Mkdir("../public/image", 0777)
	}

	_imageIds := []int{}
	for i := 0; i < 100; i++ {
		err := db.Select(&_imageIds, "SELECT `id` FROM `posts`")
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	// get file list
	imageIds := []int64{}
	files, err := os.ReadDir("../public/image")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range _imageIds {
		alreadyExists := false
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			filename := file.Name()
			id, err := strconv.Atoi(strings.Split(filename, ".")[0])
			if err != nil {
				log.Fatal(err)
			}
			if v == id {
				alreadyExists = true
			}
		}
		if !alreadyExists {
			imageIds = append(imageIds, int64(v))
		}
	}

	if len(imageIds) == 0 {
		fmt.Println("No image to download")
		return
	}

	imageDLParallel(imageIds)
}

func imageDLParallel(imageIds []int64) {
	fmt.Println("Image DL: ", len(imageIds), "images to download")

	// 非同期的にする、同時は10個まで
	var wg sync.WaitGroup
	var s = semaphore.NewWeighted(10)
	cnt := 0
	for _, u := range imageIds {
		wg.Add(1)
		go func(postId int64) {
			defer wg.Done()
			if err := s.Acquire(context.Background(), 1); err != nil {
				log.Fatal(err)
				return
			}
			defer s.Release(1)

			post := Post{}

			err := db.Get(&post, "SELECT * FROM `posts` WHERE `id` = ?", postId)
			if err != nil {
				log.Fatal(err)
			}

			ext := strings.Split(post.Mime, "/")[1]
			if ext == "jpeg" {
				ext = "jpg"
			}
			filename := fmt.Sprintf("../public/image/%d.%s", postId, ext)

			file, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
			_, err = file.Write(post.Imgdata)
			if err != nil {
				log.Fatal(err)
			}
			cnt++
			fmt.Printf("Image DL: %d/%d done\n", cnt, len(imageIds))
			defer file.Close()
		}(u)
	}
	wg.Wait()
}

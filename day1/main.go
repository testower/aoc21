package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func Simple() {
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  
  var prev int64 = -1
  count := 0
  for scanner.Scan() {
    s := scanner.Text()
    if i, err := strconv.ParseInt(s, 10, 64); err == nil {
      if prev != -1 {
        if i > prev {
          count++
        }
      }
      prev = i
    }
  }
  
  fmt.Println(count)
  
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func SlidingWindow() {
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  
  var window []int64
  var prev int64 = -1
  var count = 0
  
  for scanner.Scan() {
    s := scanner.Text()
    if x, err := strconv.ParseInt(s, 10, 64); err == nil {
      
      // push to end
      window = append(window, x)
      
      // wait until window is large enough
      if len(window) == 3 {
        var sum int64 = 0
        
        // find current sliding window sum
        for i := 0; i < len(window); i++ {
          sum += window[i]
        }
        
        // check if current sum is bigger than previous sum
        if prev != -1 {
          if sum > prev {
            
            // add to counter
            count++
          }
        }
        
        // update previous sum to last
        prev = sum
        
        // pop from beginning
        window = window[1:]
      }
      
    }
  }
  
  // Print result
  fmt.Println(count)
  
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func main() {
  SlidingWindow()
}
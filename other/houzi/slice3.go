package main

import ("fmt")

func main()  {
    fmt.Println("开始")
    var a []int
    for i := 0; i < 10; i++ {
        a=append(a,i+1)
    }
    fmt.Println("a")
    fmt.Println(a)   
    fmt.Println(A(a))
}

func A(a []map[int]int)[]int{
     aa :=make([]map[int]int,0)
     var k int
     if len(a)==10{
  k=1
     }
  for _,v:=range a[len(a)-1]{
      k=v
  }
     for _,v :=range a{
         amap:=make(map[int]int)
         k+=1
         amap[v]=k
         aa=append(aa,amap)
     }
    
    for{

    }
    if len(aa)>1{
        fmt.Println("aa",aa)
        A(aa)
    }
    return aa
}
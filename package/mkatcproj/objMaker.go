package mkatcproj

import (
	"fmt"
	"os"
)

const CPP_CODE = `#include <bits/stdc++.h>
using namespace std;
#define REP(i,n) for(int i=0, i##_len=(n); i<i##_len; ++i)
#define all(x) (x).begin(),(x).end()
using ll = long long;
string char_to_string(char val) {
  return string(1, val);
}
int char_to_int(char val) {
  return val - '0';
}
template<class T> inline bool chmin(T& a, T b) {
  if (a > b) {
    a = b;
    return true;
  }
  return false;
}
template<class T> inline bool chmax(T& a, T b) {
  if (a < b) {
    a = b;
    return true;
  }
  return false;
}
int vector_finder(std::vector<ll> vec, int number) {
  auto itr = std::find(vec.begin(), vec.end(), number);
  size_t index = std::distance( vec.begin(), itr );
  if (index != vec.size()) { // 発見できたとき
    return 1;
  }
  else { // 発見できなかったとき
    return 0;
  }
}

int main() {

}
`

const GO_CODE = `package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	
	sc.Scan()
	a := sc.Text()
	sc.Scan()
	b := sc.Text()

	fmt.Println(a + b)
}`

type Code interface {
	ObjMaker()
}

type GoCode struct {
	FileName   string
	ObjectType string
}

type CppCode struct {
	FileName   string
	ObjectType string
}

func (c *CppCode) ObjMaker() {
	if c.ObjectType == "project" {

	} else if c.ObjectType == "file" {
		file, err := os.Create(c.FileName + ".cpp")
		if err != nil {
			fmt.Println("ファイルが作成できませんでした")
			os.Exit(1)
		}
		if _, error := file.Write([]byte(CPP_CODE)); error != nil {
			fmt.Println("予期せぬエラーです")
			os.Exit(1)
		}
		fmt.Printf("ファイルを作成しました! : %v\n", file.Name())
	}
}

func (g *GoCode) ObjMaker() {
	if g.ObjectType == "project" {

	} else if g.ObjectType == "file" {
		file, err := os.Create(g.FileName + ".go")
		if err != nil {
			fmt.Println("ファイルが作成できませんでした")
			os.Exit(1)
		}
		if _, error := file.Write([]byte(GO_CODE)); error != nil {
			fmt.Println("予期せぬエラーです")
			os.Exit(1)
		}
		fmt.Printf("ファイルを作成しました! : %v\n", file.Name())
	}
}

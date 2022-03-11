package main

import "fmt"

const(
	num1 = 0x55555555  //0101010101010101 0101010101010101
	num2 = 0x33333333  //0011001100110011 0011001100110011
	num3 = 0x0f0f0f0f  //0000111100001111 0000111100001111
	num4 = 0x00ff00ff  //0000000011111111 1111111100000000
)


func reverseBits1(num uint32) uint32 {
	// 常规方式
	//ans:=uint32(0)
	//for i:=0;i<32;i++{
	//	ans |= (1 & (num >> i)) << (31 - i)
	//}

	// 神奇方式
	num = ((num >> 1) & num1) | ((num & num1) << 1)
	num = ((num >> 2) & num2) | ((num & num2) << 2)
	num = ((num >> 4) & num3) | ((num & num3) << 4)
	num = ((num >> 8) & num4) | ((num & num4) << 8)
	num = (num >> 16) | (num << 16)
	return num
}

func main(){
	var num uint32 = 0
	s:="00000010100101000001111010011100"
	fmt.Printf(s)
	fmt.Println("")
	for i:=0; i<len(s); i++{
		num |= (uint32(s[len(s) -i - 1] - '0')) << i
	}
	ans:=reverseBits1(num)
	fmt.Println(ans)
	fmt.Printf("%032b",ans)


}

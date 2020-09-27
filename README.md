# go-rand-str
random string generator for golang

# example
``` go
s := randstr.Gen(10)
fmt.Print(s) // output: mUNERA9rI2


// Letters is source of generated random string.(Default: A~Z & a~z & 0~9)
randstr.Letters = []rune("グレープフルーツ最高！！あとマンゴーも！！")
s = randstr.Gen(6)
fmt.Print(s) // output: あ高ツフマレ

randstr.Letters = []rune(randstr.StrLettersNum + randstr.StrLettersUpper)
s = randstr.Gen(8)
fmt.Print(s) // output: IIDAZG7H


// setting seed
randstr.GetRandModule().Seed(47)


// new RandStr
rs := randstr.New(
	rand.NewSource(2003),
	[]rune(randstr.StrLettersNum+"+-*/"),
)
s = rs.Gen(16)
fmt.Print(s) // output: 587466+2343008+8
```

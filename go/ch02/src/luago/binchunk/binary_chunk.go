package binchunk

type binaryChunk struct {
	header		//头部
	sizeUpValue byte    //主函数upvalue值
	mainFun *Protottype //主函数原型
}

type header struct {
	signature		[4]byte
	version  		byte
	format			byte
	luacData		[6]byte
	cintSize		byte
	sizetSize		byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt			int64
	luacNum			float64
}


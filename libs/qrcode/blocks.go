package qrcode

type block struct {
	data []byte
	ecc  []byte
}
type blockList []*block

func splitToBlocks(data <-chan byte, vi *Version) blockList {
	result := make(blockList, int(vi.NumberOfErrorCorrectionBlocksInGroup1)+int(vi.NumberOfErrorCorrectionBlocksInGroup2))
	//fmt.Println("result:",len(result)," ,group1:",vi.NumberOfErrorCorrectionBlocksInGroup1," ,group2",vi.NumberOfErrorCorrectionBlocksInGroup2)

	for b := 0; b < int(vi.NumberOfErrorCorrectionBlocksInGroup1); b++ {

		blk := new(block)
		blk.data = make([]byte, int(vi.NumberOfDataCodewordsPerBlockInGroup1))
		//fmt.Println("len blk.data:",len(blk.data))
		for cw := 0; cw < int(vi.NumberOfDataCodewordsPerBlockInGroup1); cw++ {
			//fmt.Println("cw:",cw)
			blk.data[cw] = <-data
		}
		blk.ecc = ec.calcECC(blk.data, vi.NumberOfErrorCorrectionCodewordsPreBlock)
		result[b] = blk
		//fmt.Println("b:",b,",blk:",blk)
	}

	for b := 0; b < int(vi.NumberOfErrorCorrectionBlocksInGroup2); b++ {
		blk := new(block)
		blk.data = make([]byte, int(vi.NumberOfDataCodewordsPerBlockInGroup2))
		//fmt.Println("len blk2.data:",len(blk.data),",b:",b)
		for cw := 0; cw < int(vi.NumberOfDataCodewordsPerBlockInGroup2); cw++ {
			//fmt.Println("cw2:",cw)
			blk.data[cw] = <-data
			//fmt.Println("cw2OK:",cw)
		}
		blk.ecc = ec.calcECC(blk.data, vi.NumberOfErrorCorrectionCodewordsPreBlock)
		//fmt.Println("b2:",int(vi.NumberOfErrorCorrectionBlocksInGroup1)+b,",blk:",blk)
		result[int(vi.NumberOfErrorCorrectionBlocksInGroup1)+b] = blk
	}

	return result
}

func (bl blockList) interleave(vi *Version) []byte {
	var maxCodewordCount int
	if vi.NumberOfDataCodewordsPerBlockInGroup1 > vi.NumberOfDataCodewordsPerBlockInGroup2 {
		maxCodewordCount = int(vi.NumberOfDataCodewordsPerBlockInGroup1)
	} else {
		maxCodewordCount = int(vi.NumberOfDataCodewordsPerBlockInGroup2)
	}
	resultLen := int((vi.NumberOfDataCodewordsPerBlockInGroup1+vi.NumberOfErrorCorrectionCodewordsPreBlock)*vi.NumberOfErrorCorrectionBlocksInGroup1) +
		int((vi.NumberOfDataCodewordsPerBlockInGroup2+vi.NumberOfErrorCorrectionCodewordsPreBlock)*vi.NumberOfErrorCorrectionBlocksInGroup2)

	result := make([]byte, 0, resultLen)
	for i := 0; i < maxCodewordCount; i++ {
		for b := 0; b < len(bl); b++ {
			if len(bl[b].data) > i {
				result = append(result, bl[b].data[i])
			}
		}
	}
	for i := 0; i < int(vi.NumberOfErrorCorrectionCodewordsPreBlock); i++ {
		for b := 0; b < len(bl); b++ {
			result = append(result, bl[b].ecc[i])
		}
	}
	return result
}

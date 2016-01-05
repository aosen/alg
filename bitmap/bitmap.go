/*
Author: aosen
Data: 2016-01-05
Contact: QQ 316052486
Desc:
来自于《编程珠玑》。所谓的Bit-map就是用一个bit位来标记某个元素对应的Value， 而Key即是该元素。由于采用了Bit为单位来存储数据，因此在存储空间方面，可以大大节省。
如果说了这么多还没明白什么是Bit-map，那么我们来看一个具体的例子，假设我们要对0-7内的5个元素(4,7,2,5,3)排序（这里假设这些元素没有重复）。
那么我们就可以采用Bit-map的方法来达到排序的目的。要表示8个数，我们就只需要8个Bit（1Bytes），首先我们开辟1Byte的空间，将这些空间的所有Bit位都置为0
然后遍历这5个元素，首先第一个元素是4，那么就把4对应的位置为1（可以这样操作 p+(i/8)|(0×01<<(i%8)) 当然了这里的操作涉及到Big-ending和Little-ending的情况，
这里默认为Big-ending）,因为是从零开始的，所以要把第五位置为1。
然后再处理第二个元素7，将第八位置为1,，接着再处理第三个元素，一直到最后处理完所有的元素，将相应的位置为1。
然后我们现在遍历一遍Bit区域，将该位是一的位的编号输出（2，3，4，5，7），这样就达到了排序的目的。
其实就是把计数排序用的统计数组的每个单位缩小成bit级别的布尔数组
*/
package bitmap

type Bitmap struct {
	// 保存实际的 bit 数据
	data []byte
	// 指示该 Bitmap 的 bit 容量
	bitsize uint64
	// 该 Bitmap 被设置为 1 的最大位置（方便遍历）
	maxpos uint64
}

// SetBit 将 offset 位置的 bit 置为 value (0/1)
func (self *Bitmap) SetBit(offset uint64, value uint8) bool {
	index, pos := offset/8, offset%8

	if self.bitsize < offset {
		return false
	}

	if value == 0 {
		// &^ 清位
		self.data[index] &^= 0x01 << pos
	} else {
		self.data[index] |= 0x01 << pos

		// 记录曾经设置为 1 的最大位置
		if self.maxpos < offset {
			self.maxpos = offset
		}
	}

	return true
}

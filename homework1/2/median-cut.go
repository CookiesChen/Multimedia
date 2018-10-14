package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"imgo"
	"math"
	"os"
	"sort"
)

/* variable */
var(
	colorTable []RGBcolor
)

/* type */
type RGBcolor struct {
	R int
	G int
	B int
}

type RGBcolors []RGBcolor

type sortR struct {
	RGBcolors
}

type sortG struct {
	 RGBcolors
}

type sortB struct {
	 RGBcolors
}
/* sort */
func (rgblist RGBcolors) Len() int {
	return len(rgblist)
}

func (rgblist RGBcolors) Swap(i, j int) {
	rgblist[i], rgblist[j] = rgblist[j], rgblist[i]
}

func (data sortR) Less(i, j int) bool {
	return data.RGBcolors[i].R > data.RGBcolors[j].R
}

func (data sortG) Less(i, j int) bool {
	return data.RGBcolors[i].G > data.RGBcolors[j].G
}

func (data sortB) Less(i, j int) bool {
	return data.RGBcolors[i].B > data.RGBcolors[j].B
}

func main() {
	/*file, err := os.Open("redapple.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}
	bound := img.Bounds()*/
	img := imgo.MustRead("redapple.jpg")
	y := len(img)
	x := len(img[0])
	colorlist := make(RGBcolors, x*y)
	/* 获取像素RGB值 */
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			r, g, b := img[j][i][0], img[j][i][1], img[j][i][2]
			colorlist = append(colorlist, RGBcolor{R:int(r),G:int(g),B:int(b)})
		}
	}

	/* 递归处理 */
	medianCut(colorlist[:], 0)
	newImg := image.NewRGBA(image.Rect(0,0,x,y))
	/* 计算欧式距离 */
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			var min float64
			var cor RGBcolor
			r, g, b := img[j][i][0], img[j][i][1], img[j][i][2]
			for index, v := range colorTable {
				sum := math.Pow(float64(int(r)-v.R),2)
				sum += math.Pow(float64(int(g)-v.G),2)
				sum += math.Pow(float64(int(b)-v.B),2)
				dis := math.Sqrt(sum)
				if index == 0 || dis < min{
					min = dis
					cor = v
				}
			}
			newImg.SetRGBA(i, j, color.RGBA{R:uint8(cor.R), G:uint8(cor.G), B:uint8(cor.B)})
		}
	}
	outputfile, _ := os.Create("newapple.jpg")
	jpeg.Encode(outputfile, newImg, &jpeg.Options{100})
}

func medianCut(data RGBcolors, colortype int){
	if(colortype == 9){
		// 已分成256个区间,计算平均值
		sumR, sumG, sumB := 0, 0, 0
		for _, v := range data {
			sumR = sumR + v.R
			sumG = sumG + v.G
			sumB = sumB + v.B
		}
		colorTable = append(colorTable, RGBcolor{
			R: sumR/data.Len(),
			G: sumG/data.Len(),
			B: sumB/data.Len() })
		return
	}
	switch colortype % 3 {
	case 0:
		sort.Sort(sortR{data})
	case 1:
		sort.Sort(sortG{data})
	case 2:
		sort.Sort(sortB{data})
	}
	length := data.Len()/2
	medianCut(data[:length], colortype+1)
	medianCut(data[length:], colortype+1)
}
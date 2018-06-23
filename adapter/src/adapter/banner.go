package adapter

import (
	"fmt"
)

type PrintIntf interface {
	PrintWeak()
	PrintStrong()
}

func NewPrintBannerStu(name string) *printBannerStu {
	var printbanner printBannerStu
	printbanner.Init(name)
	return &printbanner
}

type printBannerStu struct {
	banner *bannerStu
}

func (p *printBannerStu) Init(name string) {
	p.banner = newbannerStu(name)
}

func (p *printBannerStu) PrintWeak() {
	p.banner.showWithParen()
}

func (p *printBannerStu) PrintStrong() {
	p.banner.showWithAster()
}

func newbannerStu(name string) *bannerStu {
	var banner bannerStu
	banner.Init(name)
	return &banner
}

type bannerStu struct {
	Name string
}

func (b *bannerStu) Init(name string) {
	b.Name = name
}

func (b *bannerStu) showWithParen() {
	fmt.Println("(" + b.Name + ")")
}

func (b *bannerStu) showWithAster() {
	fmt.Println("*" + b.Name + "*")
}

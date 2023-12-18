package task21

import "fmt"

// Пустой класс MAIUniversity
type MAIUniversity struct{}

// Список специальностей класса MAIUniversity
func (mai *MAIUniversity) MAIMajors() {
	fmt.Println("MAI Majors: {Aerospace engineering, Computer Science}")
}

// Пустой класс MSUUniversity
type MSUUniversity struct{}

// Список специальностей класса MSUUniversity
func (msu *MSUUniversity) MSUMajors() {
	fmt.Println("MSU Majors: {Biology, Philosophy}")
}

type HSEUniversity struct{}

func (hse *HSEUniversity) HSEMajors() {
	fmt.Println("HSE Majors: {Economics, Computer Science}")
}

type UniverstyAdapter interface {
	Majors()
}

type MAIAdapter struct {
	*MAIUniversity
}

func NewMAIAdapter(mai *MAIUniversity) UniverstyAdapter {
	return &MAIAdapter{
		MAIUniversity: mai,
	}
}

func (adapter *MAIAdapter) Majors() {
	adapter.MAIMajors()
}

type MSUAdapter struct {
	*MSUUniversity
}

func NewMSUAdapter(msu *MSUUniversity) UniverstyAdapter {
	return &MSUAdapter{
		MSUUniversity: msu,
	}
}

func (adapter *MSUAdapter) Majors() {
	adapter.MSUMajors()
}

type HSEAdapter struct {
	*HSEUniversity
}

func NewHSEAdapter(hse *HSEUniversity) UniverstyAdapter {
	return &HSEAdapter{
		HSEUniversity: hse,
	}
}

func (adapter *HSEAdapter) Majors() {
	adapter.HSEMajors()
}

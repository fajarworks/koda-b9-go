package checkout

import "fmt"

type Payment interface {
	Pay(list []uint) (string, error)
}

func PayBill(pay Payment, list []uint) (string, error) {
	return pay.Pay(list)
}

type Bank struct{}

func (b Bank) Pay(list []uint) (string, error) {

	var total uint
	for _, v := range list {
		if v <= 0 {
			return "", fmt.Errorf("input tidak boleh kurang dari sama dengan kosong")
		}
		total += v
	}
	return fmt.Sprintf("total: %d, Sukses dibayarkan lewat Bank", total), nil
}

type Online struct{}

func (o Online) Pay(list []uint) (string, error) {
	var total uint
	for _, v := range list {
		if v <= 0 {
			return "", fmt.Errorf("input tidak boleh kurang dari sama dengan kosong")
		}
		total += v
	}
	return fmt.Sprintf("total: %d, sukses dibayarkan lewat online", total), nil
}

type Ficticious struct {
	Total []uint
}

func (f *Ficticious) GetTotalFictious() string {
	var total uint
	for _, v := range f.Total {
		total += v
	}
	return fmt.Sprintf("total: %d, total pembayaran dari pembayaran fiktif", total)

}

func (f *Ficticious) Pay(list []uint) (string, error) {
	var total uint
	for _, v := range list {
		if v <= 0 {
			return "", fmt.Errorf("input tidak boleh kurang dari sama dengan kosong")
		}
		total += v
	}
	f.Total = append(f.Total, total)
	return "", nil
}

/*
@Time   : 2021/12/8 16:45
@Author : ckx0709
@Remark : 年化计算
*/
package util

/*
total 总额
supplies 月供
rate 月息
*/
func GetAnnualized(total float64, supplies, rate []float64) (monthAd []float64, ad float64) {
	if len(supplies) != len(rate) {
		return
	}
	monthAd = make([]float64, len(supplies), cap(supplies))
	for i, supply := range supplies {
		monthAd[i] = rate[i] / total * 365 / 30 * 100
		ad += monthAd[i]
		total -= supply
	}

	ad = ad / float64(len(monthAd))
	return
}

package main


import (
	"fmt"
	"math"
)
// Savings
func main() {
	var grossPay int
	var taxRate float64
	var expenses int

	grossPay = 200
	taxRate = 0.2
	expenses = 45

	savings_amount := savings(grossPay, taxRate, expenses)
	fmt.Printf("Your Savings: %d\n", savings_amount)

}
func savings(grossPay int, taxRate float64, expenses int) int {

	gp_float := float64(grossPay)
	tax_amount := math.Floor(gp_float * taxRate)

	remain_bal := int(gp_float - tax_amount - float64(expenses))
	return remain_bal
}

// Material waste
func main() {
	var totalMaterial int
	var materialUnits string
	var numJobs int
	var jobConsumption int

	totalMaterial = 200
	materialUnits = "kg"
	numJobs = 20
	jobConsumption = 2
	waste := materialWaste(totalMaterial, materialUnits, numJobs, jobConsumption)
	fmt.Println(waste)
}
func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {

	total_job_consumption := numJobs * jobConsumption

	remaining_material := totalMaterial - total_job_consumption

	return fmt.Sprintf("%d%s", remaining_material, materialUnits)

}

// Interest
func main() {

	var principal int
	var rate float64
	var periods int

	principal = 2000
	rate = 0.2
	periods = 2

	investment := interest(principal, rate, periods)
	fmt.Println("Investment is currently at: ", investment)

}

func interest(principal int, rate float64, periods int) int {

	principal_float := float64(principal)
	periods_float := float64(periods)
	interest_int := principal_float * periods_float * rate
	interest_final := math.Floor(interest_int + principal_float)

	return int(interest_final)
}

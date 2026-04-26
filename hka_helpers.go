package hka

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/angelnereira/hka-sdk/types"
)

// FormatDateTime formats a time.Time to the HKA-required datetime string
// using Panama's fixed UTC-5:00 offset (no DST).
func FormatDateTime(t time.Time) string {
	loc := time.FixedZone("America/Panama", -5*60*60)
	// Go's reference offset is -07:00; with FixedZone at -5h this outputs -05:00.
	return t.In(loc).Format("2006-01-02T15:04:05-07:00")
}

// FormatDate formats a time.Time to a plain date string (YYYY-MM-DD).
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// PadDocNumber formats an integer as a 10-digit zero-padded fiscal document number.
func PadDocNumber(n int64) string {
	return fmt.Sprintf("%010d", n)
}

// PadSucursal formats a branch code string to exactly 4 characters, zero-padded on the left.
func PadSucursal(s string) string {
	return fmt.Sprintf("%04s", s)
}

// PadPunto formats a billing point number to exactly 3 digits, zero-padded.
func PadPunto(n int) string {
	return fmt.Sprintf("%03d", n)
}

// CalculateITBMS returns the ITBMS amount for a given item price and tax rate.
func CalculateITBMS(precioItem float64, tasa types.TasaITBMS) float64 {
	return precioItem * tasa.Rate()
}

// CalculatePrecioItem returns the net item price after discount.
func CalculatePrecioItem(cantidad, precioUnitario, descuento float64) float64 {
	return cantidad * (precioUnitario - descuento)
}

// CalculateValorTotal returns the total item value including all applicable taxes.
func CalculateValorTotal(precioItem, valorITBMS, valorISC, acarreo, seguro float64) float64 {
	return precioItem + valorITBMS + valorISC + acarreo + seguro
}

// FormatAmount formats a monetary amount with the given number of decimal places.
// HKA accepts 2 to 6 decimals depending on the field.
func FormatAmount(amount float64, decimals int) string {
	return strconv.FormatFloat(amount, 'f', decimals, 64)
}

// ValidateCUFE reports whether a CUFE string has the required 66-character length.
func ValidateCUFE(cufe string) bool {
	return len(cufe) == 66
}

var (
	reRUCNatural  = regexp.MustCompile(`^\d+-\d+-\d{4}$`)
	reRUCJuridico = regexp.MustCompile(`^\d{2}-\d+-\d+$`)
)

// ValidateRUCFormat performs a basic structural check on a Panamanian RUC.
// It does not validate the check digit; use QueryRUC for authoritative validation.
func ValidateRUCFormat(ruc string) bool {
	return reRUCNatural.MatchString(ruc) || reRUCJuridico.MatchString(ruc)
}

// IsWithin180Days reports whether refDate is within the 180-day window allowed
// for referenced credit/debit notes (tipos 04/05).
func IsWithin180Days(refDate time.Time) bool {
	return time.Since(refDate) <= 180*24*time.Hour
}

// IsWithin7Days reports whether emissionDate is within the 7-day cancellation window.
// The HKA service will reject cancellations outside this window with code 112.
func IsWithin7Days(emissionDate time.Time) bool {
	return time.Since(emissionDate) <= 7*24*time.Hour
}

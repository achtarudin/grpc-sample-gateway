package helper

import (
	"log"
	"os"
	"strconv"
)

// EnvType mendefinisikan tipe data yang didukung
// (string, int, bool) untuk fungsi generic kita.
type EnvType interface {
	string | int | bool
}

// GetEnvOrDefaultGeneric mengambil nilai lingkungan (string),
// mencoba mengonversinya ke tipe T, dan mengembalikan defaultValue jika gagal atau kosong.
func GetEnvOrDefault[T EnvType](key string, defaultValue T) T {
	value := os.Getenv(key)

	// Jika nilai lingkungan kosong, kembalikan defaultValue
	if value == "" {
		return defaultValue
	}

	// Gunakan type switch untuk menangani konversi berdasarkan tipe T
	switch any(defaultValue).(type) {
	case string:
		// Jika tipe T adalah string, kembalikan nilai lingkungan secara langsung
		return any(value).(T)

	case int:
		// Jika tipe T adalah int, coba konversi string ke int
		intValue, err := strconv.Atoi(value)
		if err != nil {
			// Jika konversi gagal (misal: "hello" ke int), kembalikan default
			log.Printf("Failed: Environment variable %s (%s) is not a valid integer. Using default value.\n", key, value)
			return defaultValue
		}
		return any(intValue).(T)

	case bool:
		// Jika tipe T adalah bool, coba konversi string ke bool
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			// Jika konversi gagal, kembalikan default
			log.Printf("Failed: Environment variable %s (%s) is not a valid boolean. Using default value.\n", key, value)
			return defaultValue
		}
		return any(boolValue).(T)

	default:
		// Seharusnya tidak tercapai karena EnvType sudah membatasi tipe
		return defaultValue
	}
}

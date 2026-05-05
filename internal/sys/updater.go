package sys

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

const (
	// CurrentVersion indica la versión actual en ejecución
	CurrentVersion = "7.4"
	// RemoteVersionURL es el archivo en GitHub que dice la última versión disponible
	RemoteVersionURL = "https://raw.githubusercontent.com/Depwisescript/BOT-TELEGRAM-VPN/main/version.txt"
)

// CheckForUpdate verifica si hay una actualización disponible comparando la versión local con la remota.
// Retorna (hayActualizacion, nuevaVersion, error)
func CheckForUpdate() (bool, string, error) {
	resp, err := http.Get(RemoteVersionURL)
	if err != nil {
		return false, "", fmt.Errorf("error conectando con GitHub: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, "", fmt.Errorf("código HTTP inesperado: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, "", fmt.Errorf("error leyendo versión remota: %v", err)
	}

	remoteVerStr := strings.TrimSpace(string(body))
	if remoteVerStr == "" {
		return false, "", fmt.Errorf("archivo de versión remoto vacío")
	}

	// Comparación muy simple asumiendo formato "7.4", "7.5", etc.
	localVer, errL := strconv.ParseFloat(CurrentVersion, 64)
	remoteVer, errR := strconv.ParseFloat(remoteVerStr, 64)

	if errL != nil || errR != nil {
		// Si no se puede parsear como flotante (ej: 7.4.1), comparamos como strings básicos.
		if remoteVerStr != CurrentVersion {
			return true, remoteVerStr, nil
		}
		return false, remoteVerStr, nil
	}

	if remoteVer > localVer {
		return true, remoteVerStr, nil
	}

	return false, remoteVerStr, nil
}

// RunUpdate lanza el proceso de actualización en segundo plano.
// Desvincula el proceso del bot para que sobreviva al reinicio del servicio.
func RunUpdate() error {
	// Usamos systemd-run para asegurar que el proceso sobrevive al systemctl restart depwise
	// El comando simulará un "1" por teclado para que install_go.sh haga la opción 1.
	cmdStr := `sleep 2 && echo "1" | bash <(curl -sL https://raw.githubusercontent.com/Depwisescript/BOT-TELEGRAM-VPN/main/install_go.sh)`
	
	cmd := exec.Command("systemd-run", "--unit=depwise-updater", "bash", "-c", cmdStr)
	err := cmd.Start()
	if err != nil {
		// Fallback por si systemd-run falla
		cmdFallback := exec.Command("sh", "-c", `nohup bash -c '`+cmdStr+`' > /dev/null 2>&1 &`)
		return cmdFallback.Start()
	}
	
	return nil
}

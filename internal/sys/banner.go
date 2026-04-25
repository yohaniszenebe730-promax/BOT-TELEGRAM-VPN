package sys

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Depwisescript/BOT-TELEGRAM-VPN/internal/db"
)

const (
	bannerDir       = "/etc/ssh_banners"
	sshdConfig      = "/etc/ssh/sshd_config"
	bannerMarkerStart = "# >>> DEPWISE_USER_BANNERS_START <<<"
	bannerMarkerEnd   = "# >>> DEPWISE_USER_BANNERS_END <<<"
)

// GenerateUserBanner genera el contenido HTML del banner para un usuario SSH
// Compatible con HTTP Injector, HTTP Custom, HA Tunnel y apps VPN
func GenerateUserBanner(username, title string, limit int, expireDate string) string {
	if title == "" {
		title = "INTERNET ILIMITADO"
	}

	// Calcular días restantes
	daysLeft := 0
	parsed, err := time.Parse("2006-01-02", expireDate)
	if err == nil {
		daysLeft = int(math.Ceil(time.Until(parsed).Hours() / 24))
		if daysLeft < 0 {
			daysLeft = 0
		}
	}

	limitStr := fmt.Sprintf("%d", limit)
	if limit <= 0 {
		limitStr = "∞"
	}

	var b strings.Builder

	b.WriteString("<html>\n")

	// Logo Depwise
	b.WriteString("<h5 style=\"text-align:center;\">\n")
	b.WriteString("<font face=\"monospace\" color=\"#00ff00\">\n")
	b.WriteString("⠀⠀⢀⣶⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣶⡀⠀⠀\n")
	b.WriteString("⠀⠀⢸⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀\n")
	b.WriteString("⠀⠀⢸⣿⡇⠀⠀⠀⣠⣶⣄⠀⠀⠀⢸⣿⡇⠀⠀\n")
	b.WriteString("⠀⠀⢸⣿⡇⠀⠀⢰⣿⣿⣿⡆⠀⠀⢸⣿⡇⠀⠀\n")
	b.WriteString("⠀⠀⠈⣿⣿⡄⢀⣿⣿⠻⣿⣿⡀⢠⣿⣿⠁⠀⠀\n")
	b.WriteString("⠀⠀⠀⠹⣿⣿⣾⣿⡏⠀⢹⣿⣷⣿⣿⠏⠀⠀⠀\n")
	b.WriteString("⠀⠀⠀⠀⠙⢿⣿⡿⠀⠀⠀⢿⣿⡿⠋⠀⠀⠀⠀\n")
	b.WriteString("</font>\n</h5>\n")

	// Título DEPWISE
	b.WriteString("<h1 style=\"text-align:center;\">\n")
	b.WriteString("<font face=\"monospace\" color=\"#00ff00\"><b>DEPWISE</b></font>\n")
	b.WriteString("</h1>\n")

	// Título personalizado
	b.WriteString("<h3 style=\"text-align:center;\">\n")
	b.WriteString(fmt.Sprintf("<font color='#FF00FF'><b>⚡ %s ⚡</b></font>\n", title))
	b.WriteString("</h3>\n")

	// Separador
	b.WriteString("<h5 style=\"text-align:center;\">\n")
	b.WriteString("<font color='#29b6f6'>==============================</font>\n")
	b.WriteString("</h5>\n")

	// Datos de la cuenta
	b.WriteString("<h5 style=\"text-align:center;\">\n")
	b.WriteString(fmt.Sprintf("<font color='#ffffff'>👤 Usuario: </font><font color='#f1c40f'><b>%s</b></font>\n", username))
	b.WriteString(fmt.Sprintf("<font color='#ffffff'>📅 Vence: </font><font color='#f1c40f'><b>%s</b></font>\n", expireDate))
	b.WriteString(fmt.Sprintf("<font color='#ffffff'>⏳ Días: </font><font color='#f1c40f'><b>%d</b></font>\n", daysLeft))
	b.WriteString(fmt.Sprintf("<font color='#ffffff'>💻 Límite: </font><font color='#f1c40f'><b>%s</b></font>\n", limitStr))
	b.WriteString("</h5>\n")

	// Separador
	b.WriteString("<h5 style=\"text-align:center;\">\n")
	b.WriteString("<font color='#29b6f6'>==============================</font>\n")
	b.WriteString("<font color='#29b6f6'><b>✈ TELEGRAM ✈</b></font>\n")
	b.WriteString("<font color='#29b6f6'>==============================</font>\n")
	b.WriteString("</h5>\n")

	// Contacto
	b.WriteString("<h5 style=\"text-align:center;\">\n")
	b.WriteString("<font color='#ffffff'>Dev: </font><a href=\"https://t.me/Dan3651\"><font color='#f1c40f'>@Dan3651</font></a>\n")
	b.WriteString("<font color='#ffffff'>Canal: </font><a href=\"https://t.me/Depwise2\"><font color='#f1c40f'>@Depwise2</font></a>\n")
	b.WriteString("</h5>\n")

	// Promoción
	b.WriteString("<h4 style=\"text-align:center;\">\n")
	b.WriteString("<font color='#FF00FF'><b>🔥 ¡SE VENDEN SERVIDORES PREMIUM! 🔥</b></font>\n")
	b.WriteString("</h4>\n")

	// Reglas
	b.WriteString("<h6 style=\"text-align:center;\">\n")
	b.WriteString("<font color='#ff9800'><b>⚠️ REGLAS DEL SERVIDOR ⚠️</b></font>\n")
	b.WriteString("<font color='#ffffff'>🚫 NO Torrent / P2P</font>\n")
	b.WriteString("<font color='#ffffff'>🚫 NO Spam / Fraude</font>\n")
	b.WriteString("<font color='#ffffff'>🚫 NO Ataques DDoS</font>\n")
	b.WriteString("<font color='#ff5252'><i>El incumplimiento genera ban automático</i></font>\n")
	b.WriteString("</h6>\n")

	// Crédito
	b.WriteString("<h5 style=\"text-align:center;\">\n")
	b.WriteString("<font color='#00e676'><b>CREADO EN : @Depwise_bot</b></font>\n")
	b.WriteString("</h5>\n")

	b.WriteString("</html>\n")

	return b.String()
}

// WriteUserBanner genera y escribe el banner de un usuario en /etc/ssh_banners/
func WriteUserBanner(username, title string, limit int, expireDate string) error {
	if err := os.MkdirAll(bannerDir, 0755); err != nil {
		return fmt.Errorf("error creando directorio de banners: %v", err)
	}

	content := GenerateUserBanner(username, title, limit, expireDate)
	path := filepath.Join(bannerDir, username+".banner")
	return os.WriteFile(path, []byte(content), 0644)
}

// RemoveUserBanner elimina el banner de un usuario
func RemoveUserBanner(username string) error {
	path := filepath.Join(bannerDir, username+".banner")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(path)
}

// EnsureBannerSystem configura sshd_config con Match User blocks para cada usuario SSH
func EnsureBannerSystem() error {
	if err := os.MkdirAll(bannerDir, 0755); err != nil {
		return err
	}
	return SyncSSHDBanners()
}

// SyncSSHDBanners actualiza los bloques Match User en sshd_config para apuntar
// al banner individual de cada usuario SSH
func SyncSSHDBanners() error {
	data, err := db.Load()
	if err != nil {
		return err
	}

	// Leer sshd_config actual
	raw, err := os.ReadFile(sshdConfig)
	if err != nil {
		return fmt.Errorf("no se pudo leer sshd_config: %v", err)
	}

	content := string(raw)

	// Eliminar bloque anterior de Depwise si existe
	if idx := strings.Index(content, bannerMarkerStart); idx >= 0 {
		endIdx := strings.Index(content, bannerMarkerEnd)
		if endIdx >= 0 {
			content = content[:idx] + content[endIdx+len(bannerMarkerEnd):]
		}
	}

	// Limpiar líneas vacías al final
	content = strings.TrimRight(content, "\n\t ") + "\n\n"

	// Construir nuevos bloques Match User
	var matchBlocks strings.Builder
	matchBlocks.WriteString(bannerMarkerStart + "\n")

	for user := range data.SSHTimeUsers {
		bannerFile := filepath.Join(bannerDir, user+".banner")
		if _, err := os.Stat(bannerFile); err == nil {
			matchBlocks.WriteString(fmt.Sprintf("Match User %s\n", user))
			matchBlocks.WriteString(fmt.Sprintf("    Banner %s\n\n", bannerFile))
		}
	}

	matchBlocks.WriteString(bannerMarkerEnd + "\n")

	// Escribir sshd_config actualizado
	newContent := content + matchBlocks.String()
	if err := os.WriteFile(sshdConfig, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("error escribiendo sshd_config: %v", err)
	}

	// Recargar SSH para aplicar
	exec.Command("systemctl", "reload", "ssh").Run()
	exec.Command("systemctl", "reload", "sshd").Run()

	return nil
}

// RefreshAllBanners regenera los banners de todos los usuarios SSH activos
// y sincroniza sshd_config
func RefreshAllBanners() {
	data, err := db.Load()
	if err != nil {
		return
	}

	// Solo regenerar si hay usuarios SSH
	if len(data.SSHTimeUsers) == 0 {
		return
	}

	// Asegurar que existe el directorio
	os.MkdirAll(bannerDir, 0755)

	for user, expire := range data.SSHTimeUsers {
		title := ""
		if data.SSHBannerTitles != nil {
			title = data.SSHBannerTitles[user]
		}
		limit := GetUserMaxLogins(user)
		WriteUserBanner(user, title, limit, expire)
	}

	// Sincronizar sshd_config con Match User blocks
	SyncSSHDBanners()
}

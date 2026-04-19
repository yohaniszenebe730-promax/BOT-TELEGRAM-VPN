package bot

import (
	"fmt"
	"os/exec"

	"github.com/Depwisescript/BOT-TELEGRAM-VPN/internal/db"
	"github.com/Depwisescript/BOT-TELEGRAM-VPN/internal/sys"
	tele "gopkg.in/telebot.v3"
)

func handleMenuAdmins(c tele.Context, b *tele.Bot) error {
	chatID := c.Chat().ID
	if !isAdmin(chatID) {
		return c.Send("⛔ Solo administradores pueden usar esta función.", tele.ModeHTML)
	}

	data, _ := db.Load()
	accStatus := "🔓 Público"
	if !data.PublicAccess {
		accStatus = "🔒 Privado"
	}

	markup := &tele.ReplyMarkup{}
	btnToggle := markup.Data("🔄 Acceso: "+accStatus, "toggle_public_access")
	btnList := markup.Data("📋 Listar Admins", "list_admins")
	btnAdd := markup.Data("➕ Agregar Admin", "add_admin")
	btnDel := markup.Data("➖ Quitar Admin", "del_admin_menu")
	btnRename := markup.Data("✏️ Renombrar Admin", "rename_admin_menu")
	btnInfo := markup.Data("📝 Editar Info Extra", "edit_extrainfo")
	btnCloudflare := markup.Data("☁️ Cloudflare Domain", "edit_cloudflare")
	btnCloudfront := markup.Data("🚀 Cloudfront Domain", "edit_cloudfront")
	btnBanner := markup.Data("📜 Banner SSH", "edit_banner")
	btnReset := markup.Data("🧹 Limpiar Historial", "reset_history")

	scanPubStatus := "🔓 ON"
	if !data.PublicScanner {
		scanPubStatus = "🔒 OFF"
	}
	btnScanToggle := markup.Data("🔍 Escaner Público: "+scanPubStatus, "toggle_public_scanner")

	btnReboot := markup.Data("🔄 Reiniciar VPS", "reboot_vps_confirm")
	btnAutoReboot := markup.Data("🕒 Auto Reboot", "menu_autoreboot")
	btnBackup := markup.Data("🔄 Backup a Drive", "drive_backup")
	btnRestore := markup.Data("📥 Restaurar Backup", "drive_restore")
	btnBack := markup.Data("🔙 Volver", "back_main")

	btnQuotas := markup.Data("📊 Cuotas Creación", "edit_quotas")

	markup.Inline(
		markup.Row(btnToggle),
		markup.Row(btnList, btnAdd),
		markup.Row(btnDel, btnRename),
		markup.Row(btnInfo),
		markup.Row(btnCloudflare, btnCloudfront),
		markup.Row(btnBanner, btnQuotas),
		markup.Row(btnBackup, btnRestore),
		markup.Row(btnReset, btnScanToggle),
		markup.Row(btnAutoReboot, btnReboot),
		markup.Row(btnBack),
	)

	texto := "⚙️ <b>CONFIGURACIÓN PRO (ADMIN)</b>\n"
	texto += "━━━━━━━━━━━━━━\n"
	texto += fmt.Sprintf("🛡️ <b>Acceso:</b> %s\n", accStatus)
	texto += fmt.Sprintf("🔍 <b>Escaner Público:</b> %s\n", scanPubStatus)
	texto += fmt.Sprintf("👤 <b>Admins:</b> %d\n", len(data.Admins)+1)
	texto += fmt.Sprintf("👥 <b>Historial:</b> %d IDs\n", len(data.UserHistory))
	texto += fmt.Sprintf("📊 <b>Cuotas Público:</b> %d días / %d disp.\n", data.GetMaxDaysPublic(), data.GetMaxLimitPublic())
	texto += fmt.Sprintf("📊 <b>Cuotas Admin:</b> %d días / %d disp.\n", data.GetMaxDaysAdmin(), data.GetMaxLimitAdmin())
	texto += fmt.Sprintf("💎 <b>VMess Público:</b> %d cuentas max\n", data.GetMaxXrayPublic())
	texto += fmt.Sprintf("💎 <b>VMess Admin:</b> %d cuentas max\n", data.GetMaxXrayAdmin())
	texto += "━━━━━━━━━━━━━━\n"
	texto += "<i>Selecciona una opción avanzada:</i>"

	return SafeEditCtx(c, b, texto, markup)
}

func handleTogglePublicAccess(c tele.Context, b *tele.Bot) error {
	db.Update(func(data *db.ConfigData) error {
		data.PublicAccess = !data.PublicAccess
		return nil
	})
	return handleMenuAdmins(c, b)
}

func handleEditQuotas(c tele.Context, b *tele.Bot) error {
	data, _ := db.Load()

	markup := &tele.ReplyMarkup{}
	btnDaysPub := markup.Data(fmt.Sprintf("📅 Días Público: %d", data.GetMaxDaysPublic()), "quota_days_public")
	btnLimitPub := markup.Data(fmt.Sprintf("📱 Disp. Público: %d", data.GetMaxLimitPublic()), "quota_limit_public")
	btnDaysAdm := markup.Data(fmt.Sprintf("📅 Días Admin: %d", data.GetMaxDaysAdmin()), "quota_days_admin")
	btnLimitAdm := markup.Data(fmt.Sprintf("📱 Disp. Admin: %d", data.GetMaxLimitAdmin()), "quota_limit_admin")
	btnXrayPub := markup.Data(fmt.Sprintf("💎 VMess Público: %d", data.GetMaxXrayPublic()), "quota_xray_public")
	btnXrayAdm := markup.Data(fmt.Sprintf("💎 VMess Admin: %d", data.GetMaxXrayAdmin()), "quota_xray_admin")
	btnBack := markup.Data("🔙 Volver", "menu_admins")

	markup.Inline(
		markup.Row(btnDaysPub, btnLimitPub),
		markup.Row(btnDaysAdm, btnLimitAdm),
		markup.Row(btnXrayPub, btnXrayAdm),
		markup.Row(btnBack),
	)

	texto := "📊 <b>Cuotas de Creación de Usuarios</b>\n"
	texto += "━━━━━━━━━━━━━━\n"
	texto += fmt.Sprintf("👥 <b>Público SSH:</b> %d días / %d dispositivos\n", data.GetMaxDaysPublic(), data.GetMaxLimitPublic())
	texto += fmt.Sprintf("👤 <b>Admin SSH:</b> %d días / %d dispositivos\n", data.GetMaxDaysAdmin(), data.GetMaxLimitAdmin())
	texto += "━━━━━━━━━━━━━━\n"
	texto += fmt.Sprintf("💎 <b>VMess Público:</b> máx %d cuentas\n", data.GetMaxXrayPublic())
	texto += fmt.Sprintf("💎 <b>VMess Admin:</b> máx %d cuentas\n", data.GetMaxXrayAdmin())
	texto += "━━━━━━━━━━━━━━\n"
	texto += "<i>Estos valores se aplican al crear usuarios SSH y VMess.\nEl SuperAdmin no tiene límites.</i>"

	return SafeEditCtx(c, b, texto, markup)
}

func handleQuotaPrompt(c tele.Context, b *tele.Bot, step string, label string) error {
	chatID := c.Chat().ID
	SetUserStep(chatID, step)
	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("❌ Cancelar", "edit_quotas")))
	return SafeEditCtx(c, b, fmt.Sprintf("✏️ <b>%s</b>\n\n<i>Escribe el nuevo valor (número):</i>", label), markup)
}

func handleListAdmins(c tele.Context, b *tele.Bot) error {
	data, _ := db.Load()
	res := "📋 <b>LISTADO DE ADMINISTRADORES</b>\n\n"
	res += fmt.Sprintf("⭐ <b>SuperAdmin (Root):</b> <code>%s</code>\n", superAdmin)

	if len(data.Admins) == 0 {
		res += "\n<i>No hay administradores adicionales.</i>"
	} else {
		i := 1
		for id, info := range data.Admins {
			res += fmt.Sprintf("\n%d. 👤 <b>%s</b>\n   └ ID: <code>%s</code>\n", i, info.Alias, id)
			i++
		}
	}

	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("🔙 Volver", "menu_admins")))
	return SafeEditCtx(c, b, res, markup)
}

func handleAddAdminPrompt(c tele.Context, b *tele.Bot) error {
	chatID := c.Chat().ID
	SetUserStep(chatID, "awaiting_vpn_admin_id")

	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("❌ Cancelar", "menu_admins")))

	return SafeEditCtx(c, b, "➕ <b>Agregar Nuevo Administrador</b>\n\n📝 <b>Paso 1/2:</b> Escribe el <b>ID numérico</b> del usuario de Telegram:\n\nEjemplo: <code>123456789</code>", markup)
}

func handleDelAdminMenu(c tele.Context, b *tele.Bot) error {
	data, _ := db.Load()
	if len(data.Admins) == 0 {
		return c.Respond(&tele.CallbackResponse{Text: "No hay administradores para quitar.", ShowAlert: true})
	}

	markup := &tele.ReplyMarkup{}
	var rows []tele.Row
	for id, info := range data.Admins {
		rows = append(rows, markup.Row(markup.Data("❌ "+info.Alias+" ("+id+")", "del_adm_exec", id)))
	}
	rows = append(rows, markup.Row(markup.Data("🔙 Volver", "menu_admins")))
	markup.Inline(rows...)

	return SafeEditCtx(c, b, "➖ <b>Quitar Administrador</b>\n\nSelecciona a quién deseas retirar los permisos:", markup)
}

func handleDelAdminExec(c tele.Context, b *tele.Bot) error {
	id := c.Data()

	// Buscar alias antes de borrar
	data, _ := db.Load()
	alias := "Admin"
	if info, ok := data.Admins[id]; ok {
		alias = info.Alias
	}

	db.Update(func(data *db.ConfigData) error {
		delete(data.Admins, id)
		return nil
	})

	// Responder al callback para desbloquear el botón
	c.Respond(&tele.CallbackResponse{Text: "✅ Admin eliminado"})

	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("🔙 Volver a Ajustes", "menu_admins")))
	return SafeEditCtx(c, b, fmt.Sprintf("✅ <b>Admin Eliminado</b>\n\n👤 <b>%s</b>\n🆔 ID: <code>%s</code>\n\n<i>Ya no tiene permisos de administrador.</i>", alias, id), markup)
}

func handleRenameAdminMenu(c tele.Context, b *tele.Bot) error {
	data, _ := db.Load()
	if len(data.Admins) == 0 {
		return c.Respond(&tele.CallbackResponse{Text: "No hay administradores para renombrar.", ShowAlert: true})
	}

	markup := &tele.ReplyMarkup{}
	var rows []tele.Row
	for id, info := range data.Admins {
		rows = append(rows, markup.Row(markup.Data("✏️ "+info.Alias+" ("+id+")", "rename_adm_sel", id)))
	}
	rows = append(rows, markup.Row(markup.Data("🔙 Volver", "menu_admins")))
	markup.Inline(rows...)

	return SafeEditCtx(c, b, "✏️ <b>Renombrar Administrador</b>\n\nSelecciona al admin que deseas renombrar:", markup)
}

func handleRenameAdminSelect(c tele.Context, b *tele.Bot) error {
	id := c.Data()
	chatID := c.Chat().ID

	data, _ := db.Load()
	info, exists := data.Admins[id]
	if !exists {
		return c.Respond(&tele.CallbackResponse{Text: "Admin no encontrado.", ShowAlert: true})
	}

	SetTempValue(chatID, "rename_admin_id", id)
	SetUserStep(chatID, "awaiting_rename_admin_alias")

	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("❌ Cancelar", "menu_admins")))

	return SafeEditCtx(c, b, fmt.Sprintf("✏️ <b>Renombrar Admin</b>\n\n👤 <b>Actual:</b> %s\n🆔 <b>ID:</b> <code>%s</code>\n\nEscribe el <b>nuevo alias</b>:", info.Alias, id), markup)
}

func handleEditExtraInfoPrompt(c tele.Context, b *tele.Bot) error {
	chatID := c.Chat().ID
	SetUserStep(chatID, "awaiting_vpn_extrainfo")

	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("❌ Cancelar", "menu_admins")))

	return SafeEditCtx(c, b, "📝 <b>Editar Información Extra</b>\n\nEsta información aparecerá en el menú /info.\n\n✏️ <i>Escribe el nuevo texto (soporta HTML):</i>", markup)
}

func handleEditCloudflarePrompt(c tele.Context, b *tele.Bot) error {
	chatID := c.Chat().ID
	SetUserStep(chatID, "awaiting_vpn_cloudflare")
	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("❌ Cancelar", "menu_admins")))
	return SafeEditCtx(c, b, "☁️ <b>Configurar Dominio Cloudflare</b>\n\n✏️ <i>Escribe el dominio :</i>\n\nEjemplo: <code>mi.host.com</code>", markup)
}

func handleEditCloudfrontPrompt(c tele.Context, b *tele.Bot) error {
	chatID := c.Chat().ID
	SetUserStep(chatID, "awaiting_vpn_cloudfront")
	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("❌ Cancelar", "menu_admins")))
	return SafeEditCtx(c, b, "🚀 <b>Configurar Dominio Cloudfront</b>\n\n✏️ <i>Escribe el dominio:</i>\n\nEjemplo: <code>xyz123.cloudfront.net</code>", markup)
}

// Banner predeterminado de Depwise
const defaultBanner = `<html>
<h5 style="text-align:center;">
<font face="monospace" color="#00ff00">
⠀⠀⢀⣶⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣶⡀⠀⠀
⠀⠀⢸⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀
⠀⠀⢸⣿⡇⠀⠀⠀⣠⣶⣄⠀⠀⠀⢸⣿⡇⠀⠀
⠀⠀⢸⣿⡇⠀⠀⢰⣿⣿⣿⡆⠀⠀⢸⣿⡇⠀⠀
⠀⠀⠈⣿⣿⡄⢀⣿⣿⠻⣿⣿⡀⢠⣿⣿⠁⠀⠀
⠀⠀⠀⠹⣿⣿⣾⣿⡏⠀⢹⣿⣷⣿⣿⠏⠀⠀⠀
⠀⠀⠀⠀⠙⢿⣿⡿⠀⠀⠀⢿⣿⡿⠋⠀⠀⠀⠀
</font>
</h5>
<h1 style="text-align:center;">
<font face="monospace" color="#00ff00"><b>DEPWISE</b></font>
</h1>
<h5 style="text-align:center;">
<font color='#29b6f6'>==============================</font>
<font color='#29b6f6'><b>✈ TELEGRAM ✈</b></font>
<font color='#29b6f6'>==============================</font>
</h5>
<h5 style="text-align:center;">
<font color='#ffffff'>Dev: </font><a href="https://t.me/Dan3651"><font color='#f1c40f'>@Dan3651</font></a>
<font color='#ffffff'>Canal: </font><a href="https://t.me/Depwise2"><font color='#f1c40f'>@Depwise2</font></a>
</h5>
<h4 style="text-align:center;">
<font color='#FF00FF'><b>🔥 ¡SE VENDEN SERVIDORES PREMIUM 35 DÍAS A 8.5 SOLES! 🔥</b></font>
</h4>
<h5 style="text-align:center;">
<font color='#ff0000'>==============================</font>
<font color='#ff0000'><b>⚡ SERVIDORES FREE ⚡</b></font>
<font color='#ff0000'>==============================</font>
</h5>
<h6 style="text-align:center;">
<font color='#ff9800'><b>⚠️ REGLAS DEL SERVIDOR ⚠️</b></font>
<font color='#ffffff'>🚫 NO Torrent / P2P</font>
<font color='#ffffff'>🚫 NO Spam / Fraude</font>
<font color='#ffffff'>🚫 NO Ataques DDoS</font>
<font color='#ff5252'><i>El incumplimiento genera ban automático</i></font>
</h6>
<h5 style="text-align:center;">
<font color='#00e676'><b>CREADO EN : @Depwise_bot</b></font>
</h5>
</html>`

func handleEditBannerPrompt(c tele.Context, b *tele.Bot) error {
	data, _ := db.Load()

	status := "❌ Sin banner"
	bannerType := ""
	if data.SSHBanner != "" {
		status = "✅ Activo"
		if data.SSHBanner == defaultBanner {
			bannerType = "\n📋 <b>Tipo:</b> Predeterminado (Depwise)"
		} else {
			bannerType = "\n📋 <b>Tipo:</b> Personalizado"
		}
	}

	markup := &tele.ReplyMarkup{}
	btnDefault := markup.Data("🎨 Activar Predeterminado", "banner_set_default")
	btnCustom := markup.Data("✏️ Personalizar", "banner_set_custom")
	btnDeactivate := markup.Data("🚫 Desactivar Banner", "banner_deactivate")
	btnBack := markup.Data("🔙 Volver", "menu_admins")

	markup.Inline(
		markup.Row(btnDefault),
		markup.Row(btnCustom),
		markup.Row(btnDeactivate),
		markup.Row(btnBack),
	)

	texto := fmt.Sprintf("📜 <b>Gestión de Banner SSH</b>\n\n📊 <b>Estado:</b> %s%s\n\n<i>El banner se muestra al usuario al conectar por SSH.</i>\n\n¿Qué deseas hacer?", status, bannerType)
	return SafeEditCtx(c, b, texto, markup)
}

func handleBannerSetDefault(c tele.Context, b *tele.Bot) error {
	SafeEditCtx(c, b, "⏳ <i>Activando banner predeterminado...</i>", nil)

	db.Update(func(data *db.ConfigData) error {
		data.SSHBanner = defaultBanner
		return nil
	})

	err := sys.SetSSHBanner(defaultBanner)
	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("🔙 Volver", "edit_banner")))

	if err != nil {
		return SafeEditCtx(c, b, fmt.Sprintf("⚠️ <b>Banner guardado pero error al aplicar:</b>\n%v", err), markup)
	}
	return SafeEditCtx(c, b, "✅ <b>Banner Depwise predeterminado activado y aplicado.</b>", markup)
}

func handleBannerSetCustom(c tele.Context, b *tele.Bot) error {
	chatID := c.Chat().ID
	SetUserStep(chatID, "awaiting_vpn_ssh_banner")
	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("❌ Cancelar", "edit_banner")))
	return SafeEditCtx(c, b, "📜 <b>Banner SSH Personalizado</b>\n\n✏️ <i>Escribe el texto del banner (admite HTML básico):</i>\n\nEsto se mostrará al conectar por SSH.", markup)
}

func handleBannerDeactivate(c tele.Context, b *tele.Bot) error {
	db.Update(func(data *db.ConfigData) error {
		data.SSHBanner = ""
		return nil
	})

	// Quitar banner del sistema
	exec.Command("sh", "-c", "rm -f /etc/sshd_banner").Run()
	exec.Command("sed", "-i", "/^Banner/d", "/etc/ssh/sshd_config").Run()
	exec.Command("systemctl", "reload", "ssh").Run()

	markup := &tele.ReplyMarkup{}
	markup.Inline(markup.Row(markup.Data("🔙 Volver", "edit_banner")))
	return SafeEditCtx(c, b, "✅ <b>Banner SSH desactivado.</b>\n\n<i>Ya no se mostrará ningún banner al conectar.</i>", markup)
}

func handleResetHistoryConfirm(c tele.Context, b *tele.Bot) error {
	markup := &tele.ReplyMarkup{}
	btnYes := markup.Data("✅ Sí, Limpiar", "reset_history_exec")
	btnNo := markup.Data("❌ No, Cancelar", "menu_admins")
	markup.Inline(markup.Row(btnYes, btnNo))

	return SafeEditCtx(c, b, "⚠️ <b>¿Estás seguro de limpiar el historial?</b>\n\nSe borrarán todos los IDs de usuarios registrados (el broadcast ya no les llegará hasta que vuelvan a iniciar el bot).", markup)
}

func handleResetHistoryExec(c tele.Context, b *tele.Bot) error {
	db.Update(func(data *db.ConfigData) error {
		data.UserHistory = []int64{}
		return nil
	})
	return c.Respond(&tele.CallbackResponse{Text: "Historial de IDs reseteado.", ShowAlert: true})
}

func handleServerRebootConfirm(c tele.Context, b *tele.Bot) error {
	markup := &tele.ReplyMarkup{}
	btnYes := markup.Data("🔄 Reiniciar AHORA", "reboot_vps_exec")
	btnNo := markup.Data("🔙 Cancelar", "menu_admins")
	markup.Inline(markup.Row(btnYes, btnNo))

	return SafeEditCtx(c, b, "🚨 <b>ADVERTENCIA: REINICIO DEL SERVIDOR</b>\n\n¿Estás seguro de que quieres reiniciar la VPS? Todas las conexiones actuales se cortarán.", markup)
}

func handleServerRebootExec(c tele.Context, b *tele.Bot) error {
	c.Edit("⏳ <b>Reiniciando VPS...</b> el bot estará offline unos minutos.", tele.ModeHTML)
	exec.Command("reboot").Run()
	return nil
}

func handleTogglePublicScanner(c tele.Context, b *tele.Bot) error {
	db.Update(func(data *db.ConfigData) error {
		data.PublicScanner = !data.PublicScanner
		return nil
	})
	return handleMenuAdmins(c, b)
}

func handleAutoRebootMenu(c tele.Context, b *tele.Bot) error {
	data, _ := db.Load()
	status := "❌ Desactivado"
	if data.AutoReboot {
		status = "✅ Activado"
	}

	markup := &tele.ReplyMarkup{}
	btnToggle := markup.Data("🔄 Switch: "+status, "toggle_autoreboot")
	btnBack := markup.Data("🔙 Volver", "menu_admins")

	markup.Inline(
		markup.Row(btnToggle),
		markup.Row(btnBack),
	)

	texto := "🕒 <b>CONFIGURACIÓN DE AUTO-REINICIO</b>\n"
	texto += "━━━━━━━━━━━━━━\n"
	texto += "<i>El servidor se reiniciará automáticamente cuando alcance 24 Horas de Uptime continuo.</i>\n\n"
	texto += fmt.Sprintf("📊 <b>Estado:</b> %s\n", status)
	texto += "━━━━━━━━━━━━━━\n"
	texto += "<i>Selecciona una opción:</i>"

	return SafeEditCtx(c, b, texto, markup)
}

func handleToggleAutoReboot(c tele.Context, b *tele.Bot) error {
	db.Update(func(data *db.ConfigData) error {
		data.AutoReboot = !data.AutoReboot
		return nil
	})
	return handleAutoRebootMenu(c, b)
}


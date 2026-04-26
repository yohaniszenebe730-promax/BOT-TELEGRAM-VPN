# 💎 DEPWISE BOT GO EDITION

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/OS-Ubuntu%2024.04-E95420?style=for-the-badge&logo=ubuntu" alt="Ubuntu 24.04">
  <img src="https://img.shields.io/badge/Platform-Linux-FCC624?style=for-the-badge&logo=linux" alt="Linux">
  <img src="https://img.shields.io/badge/Status-Stable-success?style=for-the-badge" alt="Status">
  <img src="https://img.shields.io/badge/Version-7.2-blue?style=for-the-badge" alt="Version">
</p>

---

## 🚀 ¿Qué es Depwise Bot?

**Depwise Bot Go Edition** es una solución integral y de alto rendimiento para la gestión de servidores VPN y cuentas SSH a través de Telegram. Reescrito completamente en **Go** para garantizar velocidad, estabilidad y bajo consumo de recursos, este bot transforma tu VPS en un panel de control profesional y automatizado.

---

## 🆕 Novedades v7.2 — Banner SSH Dinámico por Usuario

Cada cuenta SSH ahora genera automáticamente un **banner HTML personalizado** que se muestra al conectarse, compatible con **HTTP Injector**, **HTTP Custom**, **HA Tunnel** y todas las apps VPN.

### ¿Qué incluye el banner?

| Elemento | Descripción |
|----------|-------------|
| 🎨 **Logo Depwise** | Logo animado en arte braille (idéntico al banner global) |
| 🏷️ **Título Personalizado** | El admin elige el título al crear la cuenta (ej: `INTERNET ILIMITADO`, `SPEED PREMIUM VIP`) |
| 👤 **Datos de la Cuenta** | Usuario, fecha de vencimiento, días restantes, límite de dispositivos |
| 📢 **Promoción** | Canal @Depwise2 y soporte @Dan3651 |
| ⚠️ **Reglas** | Normas del servidor con advertencia de ban automático |

### Ejemplo del Banner (vista en HTTP Injector/Custom)

<p align="center">
  <img width="350" alt="Banner SSH Depwise" src="https://img.shields.io/badge/Formato-HTML%20VPN%20Apps-29b6f6?style=for-the-badge">
</p>

```text
══════════════════════
⠀⠀⢀⣶⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣶⡀⠀⠀
⠀⠀⢸⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀
⠀⠀⢸⣿⡇⠀⠀⠀⣠⣶⣄⠀⠀⠀⢸⣿⡇⠀⠀
⠀⠀⢸⣿⡇⠀⠀⢰⣿⣿⣿⡆⠀⠀⢸⣿⡇⠀⠀
⠀⠀⠈⣿⣿⡄⢀⣿⣿⠻⣿⣿⡀⢠⣿⣿⠁⠀⠀
⠀⠀⠀⠹⣿⣿⣾⣿⡏⠀⢹⣿⣷⣿⣿⠏⠀⠀⠀
⠀⠀⠀⠀⠙⢿⣿⡿⠀⠀⠀⢿⣿⡿⠋⠀⠀⠀⠀
        DEPWISE       
══════════════════════
 ⚡ INTERNET ILIMITADO ⚡
══════════════════════
 👤 Usuario: pepito
 📅 Vence: 2026-05-25
 ⏳ Días Restant.: 30
 💻 Límite: 3
══════════════════════
 🔥 ¡SERVIDORES PREMIUM A 8.5 SOLES! 🔥
 📢 Canal: @Depwise2
 👤 Soporte: @Dan3651
══════════════════════
 ✅ CREADO EN : @Depwise_bot
══════════════════════
```

> [!NOTE]
> El banner real usa **formato HTML con colores** (verde, cyan, amarillo, magenta). La vista anterior es una representación simplificada. En las apps VPN se ve con los colores del banner predeterminado de Depwise.

### ¿Cómo funciona?

### ¿Cómo funciona?

- **Banners Individuales (Por Defecto):**
  - Al crear una cuenta SSH, el bot pide un **título personalizado** para ese cliente.
  - Se genera un archivo HTML en `/etc/ssh_banners/{usuario}.banner` con sus días restantes.
  - Los **textos promocionales** (Canal, Soporte y Mensaje de Venta) se pueden editar directamente desde el menú del bot en *Ajustes Pro -> Banner -> Editar Textos Promo*.
  - Los **días restantes se actualizan automáticamente** cada 60 segundos.
  - Al renovar o eliminar un usuario, el banner se regenera o limpia automáticamente.
- **Banner Global:**
  - Si prefieres no pedir títulos personalizados para cada usuario, puedes activar el **Banner Global** desde *Ajustes Pro -> Banner*.
  - Al hacer esto, **el bot omitirá la pregunta del título al crear cuentas SSH** y todos los usuarios verán el mismo banner general.

---

## ✨ Características Principales

### 🛠️ Gestión de Protocolos (All-in-One)
- **SSH/Dropbear/SSL Tunnel:** Gestión completa de cuentas con límites de conexión y banners individuales.
- **SlowDNS:** Instalación y configuración de túneles DNS.
- **Xray (VMess):** Protocolo de última generación sobre WebSocket compatible con Cloudflare y HAProxy.
- **ZiVPN & UDP Custom:** Soporte para protocolos de gaming y bypass robusto.
- **ProxyDT:** Integración con ProxyDT Cracked para túneles HTTP estables.
- **Falcon Proxy:** Proxy de alto rendimiento integrado.
- **SSH WebSocket:** Proxy WS/WSS en puertos 10015 y 2082.

### 🔍 Escáner de Red (Módulo Independiente)
- Submenú dedicado en **Protocolos → 🔍 Escaner** con estado de instalación y barra de progreso.
- Instalación/desinstalación on-demand de `assetfinder` y `httpx`.
- No bloquea la instalación del bot — se instala cuando el admin lo necesite.

### 🛡️ Administración Pro y Utilidades
- **Ajustes Pro y Auto-Reboot:** Panel interno con control de estado, configuración de accesos públicos/privados y **reinicio automático diario**.
- **Mensaje Global (Broadcast):** Envía anuncios masivos a todos los usuarios.
- **Monitoreo en Tiempo Real:** Visualización en vivo de métricas de VPS (Uptime, Consumo) y conexiones activas (SSH, ZiVPN, Xray).
- **Gestión Avanzada:** Soporte nativo para dominios Cloudflare (CF) y Cloudfront con autogeneración de payloads.
- **Cuotas de Creación:** Límites configurables de días y dispositivos para usuarios públicos y admins.

### 🧹 Mantenimiento Inteligente
- **Persistencia de Datos Inquebrantable:** Tu tráfico de red y configuraciones de usuario están a salvo ante cualquier reinicio.
- **Resiliencia de Servicios:** Recuperación automática de protocolos mediante systemd (`Restart=always`).
- **Deep System Cleanup:** Liberación automática de memoria y procesos huérfanos.

---

## 📥 Instalación Rápida (Universal)

> [!NOTE]
> **Compatibilidad OS:** Este bot fue desarrollado y probado rigurosamente en **Ubuntu 24.04**. Se recomienda encarecidamente utilizar esta versión (o distribuciones basadas en ella) para garantizar el correcto funcionamiento de todas las dependencias (Go, Systemd, SSH, Xray, SlowDNS, etc).

Ejecuta el siguiente comando en tu terminal como usuario **root**:

```bash
bash <(curl -sL https://raw.githubusercontent.com/Depwisescript/BOT-TELEGRAM-VPN/main/install_go.sh)
```

> [!IMPORTANT]
> Selecciona la opción **1** del menú. El instalador configurará automáticamente Go, compilará el bot y desplegará el servicio Systemd 24/7.

---

## 🔄 Cómo Actualizar

Si ya tienes el bot funcionando y quieres recibir parches y nuevas funciones **sin perder usuarios ni configuraciones**:

```bash
bash <(curl -sL https://raw.githubusercontent.com/Depwisescript/BOT-TELEGRAM-VPN/main/install_go.sh)
```

> [!NOTE]
> Selecciona la opción **1 (Instalar / Actualizar Bot)**. El sistema detectará tus credenciales existentes y solo actualizará el código.

### Método alternativo (manual):

```bash
systemctl stop depwise
cd /tmp && rm -rf BOT-TELEGRAM-VPN
git clone https://github.com/Depwisescript/BOT-TELEGRAM-VPN.git
cd BOT-TELEGRAM-VPN
export PATH=$PATH:/usr/local/go/bin
go build -o /usr/local/bin/depwise-bot cmd/depwise/main.go
systemctl restart depwise
```

---

## ☁️ Copias de Seguridad en Google Drive

El bot incluye un sistema de respaldos integrado con tu Google Drive personal. Permite copias **manuales** desde el panel y **automáticas cada 24 horas**. Solo necesitas configurarlo una vez.

### Requisitos Previos

- Una cuenta de Google (Gmail)
- Acceso a [Google Cloud Console](https://console.cloud.google.com/)

### Paso 1: Crear Proyecto en Google Cloud

1. Ve a [Google Cloud Console](https://console.cloud.google.com/) e inicia sesión con tu Gmail.
2. Crea un **nuevo proyecto** (o usa uno existente). Ponle un nombre como `DepwiseBackup`.
3. En el menú lateral, ve a **APIs y Servicios** → **Biblioteca**.
4. Busca **"Google Drive API"** y haz clic en **Habilitar**.

### Paso 2: Configurar Pantalla de Consentimiento OAuth

1. Ve a **APIs y Servicios** → **Pantalla de consentimiento de OAuth**.
2. Selecciona **Externo** como tipo de usuario y haz clic en **Crear**.
3. Rellena solo los campos obligatorios:
   - **Nombre de la aplicación:** `DepwiseBot`
   - **Correo de asistencia:** tu Gmail
   - **Correo del desarrollador:** tu Gmail
4. Haz clic en **Guardar y Continuar** hasta llegar a **Usuarios de Prueba**.
5. **⚠️ IMPORTANTE:** Haz clic en **+ Agregar Usuarios** y escribe **tu propio correo de Gmail**. Sin esto, no podrás autorizar el bot.
6. Haz clic en **Guardar y Continuar** → **Volver al Panel**.

### Paso 3: Crear Credenciales OAuth

1. Ve a **APIs y Servicios** → **Credenciales**.
2. Haz clic en **+ Crear Credenciales** → **ID de cliente de OAuth**.
3. En **Tipo de aplicación**, selecciona **Aplicación de escritorio**.
4. Ponle un nombre (ej: `DepwiseBot`) y haz clic en **Crear**.
5. Se mostrará un diálogo con tu **Client ID** y **Client Secret**.
6. Haz clic en **⬇ Descargar JSON** para descargar el archivo de credenciales.

### Paso 4: Subir Credenciales al VPS

1. Renombra el archivo descargado a exactamente: `credentials.json`
2. Súbelo a tu VPS en la ruta: `/opt/depwise_bot/credentials.json`

Puedes subirlo con `scp` desde tu PC:
```bash
scp credentials.json root@TU_IP_VPS:/opt/depwise_bot/credentials.json
```

O crearlo directamente en el VPS:
```bash
nano /opt/depwise_bot/credentials.json
# Pega el contenido del JSON y guarda (Ctrl+O, Enter, Ctrl+X)
```

### Paso 5: Vincular con Telegram

1. Abre Telegram y envía al bot: `/authdrive`
2. El bot te enviará un **enlace de Google**. Ábrelo en tu navegador.
3. Inicia sesión con el **mismo Gmail** que agregaste como usuario de prueba.
4. Google te pedirá permisos — haz clic en **Continuar** y **Permitir**.
5. Al final verás una página con un **código de autorización** (o si la página muestra error de `localhost`, copia el código largo que aparece en la barra de direcciones después de `code=` y antes de `&scope`).
6. Envía al bot: `/authdrive PEGA_TU_CODIGO_AQUI`

> [!TIP]
> El código es largo (puede tener más de 50 caracteres). Cópialo completo.

### ¡Listo!

- El bot vinculará tu Drive **de por vida** (el token se renueva automáticamente).
- Se creará la carpeta **`BotVPN_Backups`** en tu Drive.
- El bot subirá tu base de datos (`bot_data.json`) automáticamente cada 24 horas.
- Mantiene solo las **2 copias más recientes** para no ocupar espacio.
- Usa el botón **📥 Restaurar Backup** desde Ajustes Pro cuando instales el bot en un nuevo VPS para recuperar todos tus usuarios automáticamente.

> [!WARNING]
> Si el token se revoca o caduca (poco frecuente), el bot te avisará. Solo necesitas repetir el paso 5 con `/authdrive`.

---

## 🛠️ Solución de Problemas (Troubleshooting)

| Síntoma | Causa Probable | Solución |
| :--- | :--- | :--- |
| **El bot no responde** | OOM mató el proceso o token inválido | `systemctl status depwise` → `systemctl restart depwise` |
| **Xray/VMess no conecta** | HAProxy o Xray no iniciaron | `systemctl status haproxy` → Reinstalar desde Protocolos |
| **VPS muy lenta** | RAM saturada | Activar **Auto Reboot** en Ajustes Pro |
| **Banner no aparece** | sshd no recargó | `systemctl reload ssh` o recrear la cuenta |
| **Error en Google Drive** | Token expirado/revocado | Enviar `/authdrive` de nuevo al bot |
| **Escáner no funciona** | Herramientas no instaladas | Ir a **Protocolos → 🔍 Escaner → 📥 Instalar Todo** |

### Comandos Útiles de Diagnóstico

```bash
# Estado del bot
systemctl status depwise

# Logs en tiempo real
journalctl -u depwise -f --no-pager -n 50

# Reiniciar bot
systemctl restart depwise

# Verificar banners de usuarios
ls -la /etc/ssh_banners/

# Ver configuración SSH (Match User blocks)
grep -A1 "Match User" /etc/ssh/sshd_config
```

---

## 💎 Créditos y Soporte

Este proyecto es desarrollado y mantenido con pasión por:

- **👨‍💻 Desarrollador:** [@Dan3651](https://t.me/Dan3651)
- **📢 Canal Oficial:** [Depwise Channel](https://t.me/Depwise2)

---

<p align="center">
  <i>"Potenciando tu VPS con la velocidad de Go."</i><br>
  <b>© 2026 Depwise Project</b>
</p>

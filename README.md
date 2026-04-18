# 💎 DEPWISE BOT GO EDITION

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Platform-Linux-FCC624?style=for-the-badge&logo=linux" alt="Linux">
  <img src="https://img.shields.io/badge/Status-Stable-success?style=for-the-badge" alt="Status">
  <img src="https://img.shields.io/badge/Version-7.1-blue?style=for-the-badge" alt="Version">
</p>

---

## 🚀 ¿Qué es Depwise Bot?

**Depwise Bot Go Edition** es una solución integral y de alto rendimiento para la gestión de servidores VPN y cuentas SSH a través de Telegram. Reesclito completamente en **Go** para garantizar velocidad, estabilidad y bajo consumo de recursos, este bot transforma tu VPS en un panel de control profesional y automatizado.

---

## ✨ Características Principales

### 🛠️ Gestión de Protocolos (All-in-One)
- **SSH/Dropbear/SSL Tunnel:** Gestión completa de cuentas con límites de conexión.
- **SlowDNS:** Instalación y configuración de túneles DNS.
- **Xray (VMess):** Protocolo de última generación sobre WebSocket compatible con Cloudflare y HAProxy.
- **ZiVPN & UDP Custom:** Soporte para protocolos de gaming y bypass robusto.
- **ProxyDT:** Integración con ProxyDT Cracked para túneles HTTP estables.
- **Falcon Proxy:** Proxy de alto rendimiento integrado.

### 🛡️ Administración Pro y Utilidades
- **Ajustes Pro y Auto-Reboot:** Panel interno con control de estado, configuración de accesos públicos/privados y **reinicio automático diario** para limpiar procesos huérfanos.
- **Mensaje Global (Broadcast):** Envía anuncios masivos a todos los usuarios.
- **Monitoreo en Tiempo Real:** Visualización en vivo de métricas de VPS (Uptime, Consumo) y escáner de red para buscar conexiones activas (SSH, ZiVPN, Xray).
- **Gestión Avanzada:** Soporte nativo para dominios Cloudflare (CF) y Cloudfront con autogeneración de payloads.

### 🧹 Mantenimiento Inteligente
- **Persistencia de Datos Inquebrantable:** Tu tráfico de red y configuraciones de usuario están a salvo; no se pierden ni siquiera ante un reinicio forzado del servidor (OOM o `reboot`).
- **Resiliencia de Servicios (Xray & HAProxy):** Recuperación automática de los protocolos mediante políticas avanzadas de systemd (Restart=always).
- **Deep System Cleanup:** Botón de un solo clic para liberar cuellos de botella de memoria (cachés, logs pesados, paquetes huérfanos).

### ☁️ Copias de Seguridad "Plug & Play" (Google Drive)
Tu bot incluye un sistema de respaldos directamente integrado con Telegram, ¡sin necesidad de comandos complicados vía SSH! Para activar las copias de seguridad (manuales y automáticas cada 24H):
1. Entra a **Google Cloud Console**, busca `Cuentas de Servicio` (Service Accounts) y crea una cuenta.
2. Entra a las opciones de esa cuenta, añade una **Clave JSON** nueva y descárgala a tu celular o PC.
3. Comparte la carpeta que prefieras en tu Google Drive con el correo robótico de la cuenta de servicio que acabas de crear.
4. **⚠️ Arrastra y envía ese archivo `.json` como un documento a tu propio Bot de Telegram.**

¡Listo! El bot encriptará permanentemente las credenciales y comenzará a enviarte tu base de datos y configuraciones directo a la nube. Si tu VPS muere, instala un bot nuevo, sube tu JSON de nuevo y ¡presiona **Restaurar Backup**!

---

## 📥 Instalación Rápida (Universal)

Ejecuta el siguiente comando en tu terminal como usuario **root**:

```bash
apt update && apt install -y git && git clone https://github.com/Depwisescript/BOT-TELEGRAM-VPN.git && cd BOT-TELEGRAM-VPN && chmod +x install_go.sh && ./install_go.sh
```

> [!IMPORTANT]
> El instalador configurará automáticamente dependencias clave, el entorno de `Go`, compilará el código y desplegará un servicio estructurado (Systemd) asegurando encendido automático 24/7.

---

## 🔄 Cómo Actualizar

Si ya tienes el bot funcionando y quieres recibir parches de seguridad y últimas funciones **sin perder usuarios ni configuraciones**, ejecuta:

```bash
wget -O install_go.sh https://raw.githubusercontent.com/Depwisescript/BOT-TELEGRAM-VPN/main/install_go.sh && chmod +x install_go.sh && ./install_go.sh
```
> [!NOTE]
> Al mostrarse el menú en terminal, elige la opción **"1. Instalar / Actualizar Bot"**. El módulo detectará tus datos y simplemente refrescará el código base.

---

## 🛠️ Solución de Problemas (Troubleshooting)

Si te encuentras con algún problema o la VPS no responde a un protocolo, revisa esta tabla:

| Síntoma / Problema | Causa Probable | Solución Recomendada (Terminal) |
| :--- | :--- | :--- |
| **El bot no responde en Telegram** | Fuga de memoria (OOM) mató el proceso o Token inválido | Verifica si está corriendo: `systemctl status depwise`. Reinícialo: `systemctl restart depwise`. |
| **Xray/VMess no conecta** | HAProxy o Xray no iniciaron o el dominio Cloudflare no es válido (TLS error). | Revisa HAProxy: `systemctl status haproxy`. Reinstala Xray desde el menú principal de utilidades del bot. |
| **La VPS se siente muy lenta** | Demasiada carga o memoria RAM saturada (sin swap). | Activa la opción **"Auto Reboot"** en el panel PRO del bot (ej: 03:00) o presiona *"Deep System Cleanup"*. |
| **Pérdida de historial de tráfico** | Apagado violento del servidor antes de grabar los datos de los primeros 60s. | No hacer hard-reset continuo. El bot automáticamente guarda el tráfico global cada ~60 segundos en caché. |
| **Problemas de puertos al instalar** | Un proceso antiguo (ej. Python, Dropbear previo) ocupa el puerto deseado. | Revisa con `netstat -tulnp`. Puedes matar el proceso usando `kill -9 PID`. |

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

#!/usr/bin/env python3
import os
import sys
import subprocess
import venv
from pathlib import Path

def setup_and_activate():
    """Создание venv и подготовка к активации"""
    
    venv_dir = Path("venv")
    
    # Создание venv
    if not venv_dir.exists():
        print("🔧 Создание виртуального окружения...")
        venv.create(venv_dir, with_pip=True)
        print("✅ Виртуальное окружение создано!")
        
        # Обновление pip
        if sys.platform == "win32":
            pip_path = venv_dir / "Scripts" / "pip.exe"
        else:
            pip_path = venv_dir / "bin" / "pip"
        
        subprocess.run([str(pip_path), "install", "--upgrade", "pip"])
        
        # Установка зависимостей
        if Path("requirements.txt").exists():
            print("📦 Установка зависимостей...")
            subprocess.run([str(pip_path), "install", "-r", "requirements.txt"])
    else:
        print("✅ Виртуальное окружение уже существует")
    
    # Создание скрипта активации
    if sys.platform == "win32":
        # Windows batch файл
        with open("activate_venv.bat", "w") as f:
            f.write("@echo off\n")
            f.write("call .\\venv\\Scripts\\activate\n")
            f.write("echo ✅ Виртуальное окружение активировано!\n")
            f.write("cmd /k\n")
        print("\n👉 Запустите для активации: activate_venv.bat")
        
        # PowerShell скрипт
        with open("activate_venv.ps1", "w") as f:
            f.write(".\\venv\\Scripts\\Activate.ps1\n")
            f.write("Write-Host \"✅ Виртуальное окружение активировано!\" -ForegroundColor Green\n")
        print("👉 Или: .\\activate_venv.ps1")
        
    else:
        # Unix shell скрипт
        with open("activate_venv.sh", "w") as f:
            f.write("#!/bin/bash\n")
            f.write("source venv/bin/activate\n")
            f.write('echo "✅ Виртуальное окружение активировано!"\n')
            f.write('exec "$SHELL"\n')
        
        # Делаем скрипт исполняемым
        os.chmod("activate_venv.sh", 0o755)
        print("\n👉 Запустите для активации: source activate_venv.sh")
        print("   или: . ./activate_venv.sh")

if __name__ == "__main__":
    setup_and_activate()
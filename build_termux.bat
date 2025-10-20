@echo off
REM Windows Batch Script for Termux Build Instructions
REM This script provides instructions for building in Termux

echo ========================================
echo Advanced HTTP Tool - Termux Build Guide
echo ========================================
echo.
echo This is a Windows system. To build for Termux:
echo.
echo 1. Install Termux on your Android device
echo 2. Open Termux
echo 3. Run these commands in Termux:
echo.
echo    pkg update ^&^& pkg upgrade
echo    pkg install golang git
echo    git clone [your-repo-url]
echo    cd [repo-directory]
echo    chmod +x build_termux.sh
echo    ./build_termux.sh
echo.
echo 4. Run the tool:
echo    ./http_tool -target=https://example.com
echo.
echo ========================================
echo Files created for Termux:
echo - advanced_http_tool.go (main tool)
echo - build_termux.sh (build script)
echo - test_termux.sh (test script)
echo - termux_config.json (configuration)
echo - TERMUX_GUIDE.md (detailed guide)
echo ========================================
echo.
pause

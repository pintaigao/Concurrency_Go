const electron = require("electron");

const { app, BrowserWindow, ipcMain } = electron;

app.on('ready', () => {
  let mainWindow = new BrowserWindow({});
  mainWindow.loadURL(`file://${__dirname}/index.html`)
});
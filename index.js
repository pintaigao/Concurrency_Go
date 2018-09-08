const electron = require("electron");

const { app, BrowserWindow, ipcMain } = electron;
const ffmpeg = require("fluent-ffmpeg")

let mainWindow;

app.on('ready', () => {
  mainWindow = new BrowserWindow({});
  mainWindow.loadURL(`file://${__dirname}/index.html`)
});

ipcMain.on("video:submit", (event, path) => {
  ffmpeg.ffprobe(path, (error, metadata) => {
    mainWindow.webContents.send("video:metadata".metadata.format.duration);
  })
})
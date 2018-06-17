// Modules to control application life and create native browser window
import { app, BrowserWindow, Menu, MenuItem } from 'electron';
import { autoUpdater } from 'electron-updater';

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let mainWindow: Electron.BrowserWindow | null;

function createWindow() {
  // Create the browser window.
  mainWindow = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      webSecurity: false
    }
  });

  const fileName = (() => {
    console.log(process.env.NODE_ENV);
    if (process.env.NODE_ENV === 'production') {
      return `file://${__dirname}/index.html`;
    } else {
      return `http://localhost:${process.env.ELECTRON_WEBPACK_WDS_PORT}`;
    }
  })();
  // and load the index.html of the app.
  mainWindow.loadURL(fileName);

  // Open the DevTools.
  // mainWindow.webContents.openDevTools()

  // Emitted when the window is closed.
  mainWindow.on('closed', function() {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    mainWindow = null;
  });
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', () => {
  const menu = Menu.buildFromTemplate(menuTemplate);
  Menu.setApplicationMenu(menu);
  createWindow();
  setTimeout(() => {
    if (mainWindow && mainWindow.isFocused()) {
      autoUpdater.checkForUpdatesAndNotify();
    }
  }, 30000);
});

// Quit when all windows are closed.
app.on('window-all-closed', function() {
  // On OS X it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('activate', function() {
  // On OS X it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (mainWindow === null) {
    createWindow();
  }
});

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.

const openAbout = () => {
  const modalPath = 'http://localhost:4200#/menu/about';
  let win: BrowserWindow | null = new BrowserWindow({
    width: 400,
    height: 320
  });

  win.on('close', () => {
    win = null;
  });
  win.loadURL(modalPath);
  win.show();
};

const menuTemplate: any[] = [
  new MenuItem({
    label: 'About',
    submenu: [
      {
        label: 'About',
        click: () => openAbout()
      },
      {
        label: 'Toggle Developer Tools',
        accelerator: (() => {
          if (process.platform === 'darwin') {
            return 'Alt+Command+I';
          } else {
            return 'Ctrl+Shift+I';
          }
        })(),
        click: (_, focusedWindow) => {
          if (focusedWindow) {
            focusedWindow.webContents.openDevTools();
          }
        }
      },
      {
        label: 'Reload window',
        accelerator: (() => {
          if (process.platform === 'darwin') {
            return 'Command+R';
          } else {
            return 'Ctrl+R';
          }
        })(),
        click: (_, focusedWindow) => {
          if (focusedWindow) {
            focusedWindow.webContents.reload();
          }
        }
      }
    ]
  })
];

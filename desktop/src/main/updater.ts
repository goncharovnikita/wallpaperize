import { BrowserWindow } from 'electron';
import { autoUpdater } from 'electron-updater';

export class Updater {
  constructor(private win: BrowserWindow) {}
  init(): void {
    autoUpdater.on('checking-for-update', () => {
      this.sendStatus('Checking for update...');
    });
    autoUpdater.on('update-available', info => {
      this.sendStatus('Update available.');
    });
    autoUpdater.on('update-not-available', info => {
      this.sendStatus('Update not available.');
    });
    autoUpdater.on('error', err => {
      this.sendStatus('Error in auto-updater. ' + err);
    });
    autoUpdater.on('download-progress', progressObj => {
      let log_message = 'Download speed: ' + progressObj.bytesPerSecond;
      log_message = log_message + ' - Downloaded ' + progressObj.percent + '%';
      log_message =
        log_message +
        ' (' +
        progressObj.transferred +
        '/' +
        progressObj.total +
        ')';
      this.sendStatus(log_message);
    });
    autoUpdater.on('update-downloaded', info => {
      this.sendStatus('Update downloaded');
    });

    autoUpdater.checkForUpdatesAndNotify();
    setTimeout(() => {
      this.sendStatus('tests');
    }, 2000);
  }

  sendStatus(msg: string): void {
    console.log(msg);
    this.win.webContents.send('updater-message', msg);
  }
}

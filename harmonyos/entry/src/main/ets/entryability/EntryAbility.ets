import { AbilityConstant, UIAbility, Want } from '@kit.AbilityKit';
import { hilog } from '@kit.PerformanceAnalysisKit';
import { window } from '@kit.ArkUI';
import { Log } from '@abner/log';
import { ZRouter } from '@hzw/zrouter';
import BuildProfile from 'BuildProfile';
import { GlobalContext } from 'common';
import { emitter } from '@kit.BasicServicesKit';
import { PreferencesUtil } from '@pura/harmony-utils';

export default class EntryAbility extends UIAbility {
  onCreate(want: Want, launchParam: AbilityConstant.LaunchParam): void {
    GlobalContext.getContext().setObject('abilityWant', want)
    GlobalContext.getContext().setObject('context', this.context)
    if (want.parameters) {
      if (want.parameters.currentTime) {
        GlobalContext.getContext().setObject('currentTime', want.parameters.currentTime);
        console.info( 'time: ' + want.parameters.currentTime);
      }
    }
    hilog.info(0x0000, 'testTag', '%{public}s', 'Ability onCreate');
    Log.init({
      tag: "HarmonyOSLog",
      domain: 0x0000,
      close: false,
      isHilog: true,
      showLogLocation: true,
      logSize: 800
    })
    ZRouter.initialize((config) => {
      config.isLoggingEnabled = BuildProfile.DEBUG
      config.isHSPModuleDependent = false
    })
  }

  onDestroy(): void {
    hilog.info(0x0000, 'testTag', '%{public}s', 'Ability onDestroy');
  }

  onWindowStageCreate(windowStage: window.WindowStage): void {
    // Main window is created, set main page for this ability
    hilog.info(0x0000, 'testTag', '%{public}s', 'Ability onWindowStageCreate');
    windowStage.getMainWindow().then((win: window.Window) => {

      win.on('windowSizeChange', (newSize: window.Size) => {
        let eventWHData: emitter.EventData = {
          data: {
            'width': newSize.width,
            'height': newSize.height
          }
        };
        let innerEventWH: emitter.InnerEvent = {
          eventId: 3,
          priority: emitter.EventPriority.HIGH
        };
        emitter.emit(innerEventWH, eventWHData);
      });
    });
    // PersistentStorage.persistProp("user",undefined)
    windowStage.loadContent('pages/Index', (err) => {
      PreferencesUtil.deleteSync("user")
      if (err.code) {
        hilog.error(0x0000, 'testTag', 'Failed to load the content. Cause: %{public}s', JSON.stringify(err) ?? '');
        return;
      }
      hilog.info(0x0000, 'testTag', 'Succeeded in loading the content.');
    });
  }

  onWindowStageDestroy(): void {
    // Main window is destroyed, release UI related resources
    hilog.info(0x0000, 'testTag', '%{public}s', 'Ability onWindowStageDestroy');
  }

  onForeground(): void {
    // Ability has brought to foreground
    hilog.info(0x0000, 'testTag', '%{public}s', 'Ability onForeground');
  }

  onBackground(): void {
    // Ability has back to background
    hilog.info(0x0000, 'testTag', '%{public}s', 'Ability onBackground');
  }
}

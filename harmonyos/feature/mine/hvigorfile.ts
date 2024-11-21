import { harTasks } from '@ohos/hvigor-ohos-plugin';
import { routerRegisterPlugin, PluginConfig } from 'router-register-plugin'
// 2、初始化配置
const config: PluginConfig = {
    scanDirs: ['src/main/ets/components'], // 扫描的目录，如果不设置，默认是扫描src/main/ets目录
    logEnabled: true, // 查看日志
    viewNodeInfo: false, // 查看节点信息
    isAutoDeleteHistoryFiles: true // 删除无用编译产物

}
export default {
    system: harTasks,  /* Built-in plugin of Hvigor. It cannot be modified. */
    plugins:[routerRegisterPlugin(config)]
}

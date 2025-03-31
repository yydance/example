import type { ProLayoutProps } from '@ant-design/pro-components';

/**
 * @name
 */
const Settings: ProLayoutProps & {
  pwa?: boolean;
  logo?: string;
} = {
  navTheme: 'light',
  // 拂晓蓝
  colorPrimary: '#1890ff',
  layout: 'side',
  contentWidth: 'Fluid',
  fixedHeader: false,
  fixSiderbar: true,
  colorWeak: false,
  title: 'demo ui',
  pwa: true,
  logo: './power.svg',
  iconfontUrl: '',
  token: {
    // 参见ts声明，demo 见文档，通过token 修改样式
    //https://procomponents.ant.design/components/layout#%E9%80%9A%E8%BF%87-token-%E4%BF%AE%E6%94%B9%E6%A0%B7%E5%BC%8F
    sider: {
      colorBgCollapsedButton: '#1890ff',
      colorTextMenuTitle: '#1890ff',
      colorMenuItemDivider: '#fff',
      colorTextCollapsedButtonHover: '#fff',
      colorTextMenuSecondary: '#fff',
      colorTextCollapsedButton: '#fff',
      colorMenuBackground: '#003066',
      // 修改为已知属性 colorBgMenuItemCollapsedElevated
      colorBgMenuItemCollapsedElevated: '#001529',
      colorTextMenu: '#fff',
      colorTextMenuSelected: '#fff',
      colorTextMenuActive: '#fff',
      colorBgMenuItemHover: '#9933cc',
      colorBgMenuItemSelected: '#9933cc',
      colorTextMenuItemHover: '#fff',
      colorTextSubMenuSelected: '#fff',
      //menuHeight: 30,
    },
    /*
    header: {
      colorBgHeader: '#fff',
      colorHeaderTitle: '#fff',
    },
    */
  },
};

export default Settings;

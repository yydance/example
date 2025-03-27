import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-components';
import React from 'react';

const Footer: React.FC = () => {
  return (
    <DefaultFooter
      copyright='翼鸥科技'
      style={{
        background: 'none',
      }}
      links={[
        {
          key: 'DevOps',
          title: 'DevOps',
          href: 'https://www.eeo.cn',
          blankTarget: true,

        },
        {
          key: 'github',
          title: <GithubOutlined />,
          href: 'https://github.com/ant-design/ant-design-pro',
          blankTarget: true,
        },
        {
          key: 'EEOLab',
          title: 'EEOLab',
          href: 'https://ant.design',
          blankTarget: true,
        },
      ]}
    />
  );
};

export default Footer;

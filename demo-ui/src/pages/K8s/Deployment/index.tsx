import React, { useState, useEffect } from 'react';
import { Button, Modal, Form, Input, message, Select, Space } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import ProTable  from '@ant-design/pro-table';
import { SearchOutlined, PlusOutlined } from '@ant-design/icons';
import { listDeployment } from '@/services/k8s/deployment';


const fetchServiceData = async (params: {
    current?: number;
    pageSize?: number;
    name?: string;
    namespace?: string;
  }) => {
    try {
      const response = await listDeployment(params);
      return {
        data: response.data,
        total: response.total,
        success: response.success,
      };
    } catch (error) {
      console.error('获取deployment数据失败:', error);
      return {
        data: [],
        total: 0,
        success: false,
      };
    }
  };

  